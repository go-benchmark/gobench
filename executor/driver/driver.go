package driver

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"

	"github.com/gobench-io/gobench/dis"
	"github.com/gobench-io/gobench/logger"
	"github.com/gobench-io/gobench/metrics"
	"github.com/gobench-io/gobench/pb"
	"github.com/gobench-io/gobench/scenario"
	gometrics "github.com/rcrowley/go-metrics"
)

// Error
var (
	ErrIDNotFound    = errors.New("id not found")
	ErrNodeIsRunning = errors.New("driver is running")

	ErrAppCancel = errors.New("application is cancel")
	ErrAppPanic  = errors.New("application is panic")
)

// driver status. the driver is in either idle, or running state
type status string

const (
	Idle    status = "idle"
	Running status = "running"
)

type unit struct {
	Title    string             // metric title
	Type     metrics.MetricType // to know the current unit type
	metricID int                // metric table foreign key
	c        gometrics.Counter
	h        gometrics.Histogram
	g        gometrics.Gauge
}

// Driver is the main structure for a running driver
// contains host information, the scenario (plugin)
// and gometrics unit
type Driver struct {
	mu     sync.Mutex
	appID  int
	eID    string
	status status
	vus    scenario.Vus
	units  map[string]unit // title - gometrics

	logger logger.Logger
	ml     pb.AgentClient
}

// the singleton driver variable
var driver Driver

func init() {
	driver = Driver{
		status: Idle,
	}
}

// NewDriver returns the singleton driver
func NewDriver(ml pb.AgentClient, logger logger.Logger, vus scenario.Vus, appID int, eID string) (*Driver, error) {
	driver.mu.Lock()
	defer driver.mu.Unlock()

	driver.ml = ml
	driver.logger = logger
	driver.units = make(map[string]unit)
	driver.appID = appID
	driver.eID = eID
	// reset metrics
	driver.unregisterGometrics()

	driver.vus = vus

	return &driver, nil
}

func (d *Driver) unregisterGometrics() {
	gometrics.Each(func(name string, i interface{}) {
		gometrics.Unregister(name)
	})
}

func (d *Driver) reset() {
	d.mu.Lock()
	d.status = Idle
	d.units = make(map[string]unit)
	d.mu.Unlock()
}

// SetNopMetricLog update the driver metric logger to the nop one. Mostly use
// for testing
func (d *Driver) SetNopMetricLog() error {
	nop := newNopMetricLog()

	d.mu.Lock()
	d.ml = nop
	d.mu.Unlock()

	return nil
}

// Run starts the preloaded plugin
// return error if the driver is running already
func (d *Driver) Run(ctx context.Context) (err error) {
	// first, setup driver system load
	if err = d.systemloadSetup(); err != nil {
		return
	}

	d.mu.Lock()

	if d.status == Running {
		d.mu.Unlock()
		return ErrNodeIsRunning
	}

	d.status = Running
	d.mu.Unlock()

	err = d.run(ctx)

	return err
}

func (d *Driver) run(ctx context.Context) (err error) {
	finished := make(chan error)

	// when the runScen finished, we should stop the logScaled and systemloadRun
	// also; however, not necessary since the executor will be shutdown anyway
	go d.logScaled(ctx, 10*time.Second)
	go d.runScen(ctx, finished)
	go d.systemloadRun(ctx)

	select {
	case err = <-finished:
	case <-ctx.Done():
		err = ErrAppCancel
	}

	// when finish, reset the driver
	d.reset()

	return
}

// Running returns a bool value indicating that the working is running
func (d *Driver) Running() bool {
	d.mu.Lock()
	defer d.mu.Unlock()

	return d.status == Running
}

func (d *Driver) runScen(ctx context.Context, done chan<- error) {
	var totalVu int

	vus := d.vus
	for i := range vus {
		totalVu += vus[i].Nu
	}

	var wg sync.WaitGroup
	wg.Add(totalVu)

	for i := range vus {
		go func(i int) {
			for j := 0; j < vus[i].Nu; j++ {
				go func(i, j int) {
					vus[i].Fu(ctx, j)
					wg.Done()
				}(i, j)
				dis.SleepRatePoisson(vus[i].Rate)
			}
		}(i)
	}

	wg.Wait()
	done <- nil
}

// logScaled extract the metric log from a driver
// should run this function in a routine
func (d *Driver) logScaled(ctx context.Context, freq time.Duration) {
	ch := make(chan interface{})

	go func(channel chan interface{}) {
		for range time.Tick(freq) {
			channel <- struct{}{}
		}
	}(ch)

	if err := d.logScaledOnCue(ctx, ch); err != nil {
		d.logger.Fatalw("failed logScaledOnCue", "err", err)
	}
}

func (d *Driver) logScaledOnCue(ctx context.Context, ch chan interface{}) error {
	var err error
	for {
		select {
		case <-ch:
			now := timestampMs()
			d.mu.Lock()
			units := d.units
			d.mu.Unlock()

			for _, u := range units {
				base := &pb.BasedReqMetric{
					AppID: int64(d.appID),
					EID:   d.eID,
					MID:   int64(u.metricID),
					Time:  now,
				}

				switch u.Type {
				case metrics.Counter:
					_, err = d.ml.Counter(ctx, &pb.CounterReq{
						Base:  base,
						Count: u.c.Count(),
					})
				case metrics.Histogram:
					h := u.h.Snapshot()
					ps := h.Percentiles([]float64{0.5, 0.75, 0.95, 0.99, 0.999})
					hv := &pb.HistogramValues{
						Count:  h.Count(),
						Min:    h.Min(),
						Max:    h.Max(),
						Mean:   h.Mean(),
						Stddev: h.StdDev(),
						Median: ps[0],
						P75:    ps[1],
						P95:    ps[2],
						P99:    ps[3],
						P999:   ps[4],
					}
					_, err = d.ml.Histogram(ctx, &pb.HistogramReq{
						Base:      base,
						Histogram: hv,
					})
				case metrics.Gauge:
					_, err = d.ml.Gauge(ctx, &pb.GaugeReq{
						Base:  base,
						Gauge: u.g.Value(),
					})
				}

				if err != nil {
					d.logger.Errorw("metric log failed", "err", err)
				}
			}
		case <-ctx.Done():
			d.logger.Infow("logScaledOnCue canceled")
			return nil
		}
	}
}

func timestampMs() int64 {
	return time.Now().UnixNano() / 1e6 // ms
}

// Setup is used for the driver to report the metrics that it will generate
func Setup(groups []metrics.Group) error {
	ctx := context.TODO()

	units := make(map[string]unit)

	driver.mu.Lock()
	defer driver.mu.Unlock()

	for _, group := range groups {
		// create a new group if not existed
		egroup, err := driver.ml.FindCreateGroup(ctx, &pb.FCGroupReq{
			AppID: int64(driver.appID),
			Name:  group.Name,
		})
		if err != nil {
			return fmt.Errorf("failed create group: %v", err)
		}

		for _, graph := range group.Graphs {
			// create new graph if not existed
			egraph, err := driver.ml.FindCreateGraph(ctx, &pb.FCGraphReq{
				AppID:   int64(driver.appID),
				Title:   graph.Title,
				Unit:    graph.Unit,
				GroupID: egroup.Id,
			})
			if err != nil {
				return fmt.Errorf("failed create graph: %v", err)
			}

			for _, m := range graph.Metrics {
				// create new metric if not existed
				emetric, err := driver.ml.FindCreateMetric(ctx, &pb.FCMetricReq{
					AppID:   int64(driver.appID),
					Title:   m.Title,
					Type:    string(m.Type),
					GraphID: egraph.Id,
				})
				if err != nil {
					return fmt.Errorf("failed create metric: %v", err)
				}

				// counter type
				if m.Type == metrics.Counter {
					c := gometrics.NewCounter()
					if err := gometrics.Register(m.Title, c); err != nil {
						if _, ok := err.(gometrics.DuplicateMetric); ok {
							continue
						}
						return err
					}

					units[m.Title] = unit{
						Title:    m.Title,
						Type:     m.Type,
						metricID: int(emetric.Id),
						c:        c,
					}
				}

				if m.Type == metrics.Histogram {
					s := gometrics.NewExpDecaySample(1028, 0.015)
					h := gometrics.NewHistogram(s)
					if err := gometrics.Register(m.Title, h); err != nil {
						if _, ok := err.(gometrics.DuplicateMetric); ok {
							continue
						}
						return err
					}
					units[m.Title] = unit{
						Title:    m.Title,
						Type:     m.Type,
						metricID: int(emetric.Id),
						h:        h,
					}
				}

				if m.Type == metrics.Gauge {
					g := gometrics.NewGauge()
					if err := gometrics.Register(m.Title, g); err != nil {
						if _, ok := err.(gometrics.DuplicateMetric); ok {
							continue
						}
						return err
					}
					units[m.Title] = unit{
						Title:    m.Title,
						Type:     m.Type,
						metricID: int(emetric.Id),
						g:        g,
					}
				}
			}
		}
	}

	// aggregate units
	for k, v := range units {
		driver.units[k] = v
	}

	return nil
}

// Notify saves the id with value into metrics which later save to database
// Return error when the title is not found from the metric list.
// The not found error may occur because
// a. The title has never ever register before
// b. The session is cancel but the scenario does not handle the ctx.Done signal
func Notify(title string, value int64) error {
	driver.mu.Lock()
	defer driver.mu.Unlock()

	u, ok := driver.units[title]
	if !ok {
		driver.logger.Infow("metric not found", "title", title)
		return ErrIDNotFound
	}

	if u.Type == metrics.Counter {
		u.c.Inc(value)
	}

	if u.Type == metrics.Histogram {
		u.h.Update(value)
	}

	if u.Type == metrics.Gauge {
		u.g.Update(value)
	}

	return nil
}
