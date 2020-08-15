package master

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/rpc"
	"os"
	"os/exec"
	"strconv"
	"sync"
	"time"

	"context"

	"github.com/gobench-io/gobench/agent"
	"github.com/gobench-io/gobench/ent"
	"github.com/gobench-io/gobench/ent/application"
	"github.com/gobench-io/gobench/logger"

	_ "github.com/mattn/go-sqlite3"
)

// job status. The job is in either pending, provisioning, running, finished
// cancel, error states
type jobState string

// App states
const (
	jobPending      jobState = "pending"
	jobProvisioning jobState = "provisioning"
	jobRunning      jobState = "running"
	jobFinished     jobState = "finished"
	jobCancel       jobState = "cancel"
	jobError        jobState = "error"
)

type Master struct {
	mu          sync.Mutex
	addr        string // host name
	port        int    // api port
	clusterPort int    // cluster port

	status  status
	logger  logger.Logger
	program string

	// database
	isScheduled bool
	dbFilename  string
	db          *ent.Client

	la  *agent.Agent // local agent
	job *job
}

type job struct {
	app    *ent.Application
	plugin string // plugin path
	cancel context.CancelFunc
}

type Options struct {
	Port    int
	Addr    string
	DbPath  string
	Program string
}

func NewMaster(opts *Options, logger logger.Logger) (m *Master, err error) {
	logger.Infow("new master program",
		"port", opts.Port,
		"db file path", opts.DbPath,
	)

	m = &Master{
		addr:       opts.Addr,
		port:       opts.Port,
		dbFilename: opts.DbPath,
		logger:     logger,
		program:    opts.Program,
	}
	la, err := agent.NewAgent(m)
	if err != nil {
		return
	}
	m.la = la

	return
}

func (m *Master) Start() (err error) {
	if err = m.setupDb(); err != nil {
		return
	}

	m.handleSignals()

	go m.schedule()

	// start the local agent socket server that communicate with local executor
	agentSocket := fmt.Sprintf("/tmp/gobench-agentsocket-%d", os.Getpid())
	err = m.la.StartSocketServer(agentSocket)

	return
}

// DB returns the db client
func (m *Master) DB() *ent.Client {
	return m.db
}

func (m *Master) finish(status status) error {
	m.logger.Infow("server is shutting down")

	m.mu.Lock()
	m.status = status
	m.mu.Unlock()

	// todo: if there is a running scenario, shutdown
	// todo: send email if needed
	return m.db.Close()
}

// WebPort returns the master HTTP web port
func (m *Master) WebPort() int {
	return m.port
}

// NewApplication create a new application with a name and a scenario
// return the application id and error
func (m *Master) NewApplication(ctx context.Context, name, scenario string) (
	*ent.Application, error,
) {
	app, err := m.db.Application.
		Create().
		SetName(name).
		SetScenario(scenario).
		SetStatus(string(jobPending)).
		Save(ctx)
	if err != nil {
		return nil, err
	}

	m.LogEvent(ctx, app.ID, fmt.Sprintf("application is %s", app.Status), "master:newApplication", "", "")

	return app, err
}

// DeleteApplication a pending/finished/canceled/error application
func (m *Master) DeleteApplication(ctx context.Context, appID int) error {
	app, err := m.db.Application.
		Query().
		Where(application.ID(appID)).
		Only(ctx)

	if err != nil {
		return err
	}

	if app.Status != string(jobPending) && app.Status != string(jobCancel) &&
		app.Status != string(jobFinished) && app.Status != string(jobError) {
		return fmt.Errorf(ErrCantDeleteApp.Error(), string(app.Status))
	}

	return m.db.Application.
		DeleteOneID(appID).
		Exec(ctx)
}

// CancelApplication terminates an application
// if the app is running, send cancel signal
// if the app is finished/error, return ErrAppIsFinished error
// if the app is canceled, return with current app status
// else update app status with cancel
func (m *Master) CancelApplication(ctx context.Context, appID int) (*ent.Application, error) {
	err := m.cancel(ctx, appID)

	if err == nil {
		return m.db.Application.
			Query().
			Where(application.ID(appID)).
			Only(ctx)
	}

	// if err and err is not the app is not running
	if err != nil && !errors.Is(err, ErrAppNotRunning) {
		return nil, err
	}

	// if the app is not running, update directly on the db
	// query the app
	// if the app status is finished or error, return error
	// if the app status is cancel (already), just return
	// else, update the app table
	app, err := m.db.Application.
		Query().
		Where(application.ID(appID)).
		Only(ctx)

	if err != nil {
		return app, err
	}

	if app.Status == string(jobCancel) {
		return app, nil
	}
	if app.Status == string(jobFinished) || app.Status == string(jobError) {
		return app, ErrAppIsFinished
	}

	currentStatus := app.Status
	// else, update the status on db
	app, err = m.db.Application.
		UpdateOneID(appID).
		SetStatus(string(jobCancel)).
		Save(ctx)
	if err != nil {
		return app, err
	}

	m.LogEvent(ctx, app.ID, fmt.Sprintf("application status change from %s to %s", currentStatus, app.Status), "master:cancelApplication", "", "")

	return app, err
}

// SetApplicationTags set application tags
func (m *Master) SetApplicationTags(ctx context.Context, appID int, tags string) (*ent.Application, error) {
	return m.db.Application.
		UpdateOneID(appID).
		SetTags(tags).
		Save(ctx)
}

// cleanupDB is the helper function to cleanup the DB for testing
func (m *Master) cleanupDB() error {
	ctx := context.TODO()
	_, err := m.db.Application.Delete().Exec(ctx)
	return err
}

// to is the function to set new state for an application
// save new state to the db
func (m *Master) jobTo(ctx context.Context, state jobState) (err error) {
	currentStatus := m.job.app.Status

	m.job.app, err = m.job.app.Update().
		SetStatus(string(state)).
		Save(ctx)

	m.LogEvent(ctx, m.job.app.ID, fmt.Sprintf("application status change from %s to %s", currentStatus, m.job.app.Status), "master:jobTo", "", "")

	return
}

// setupDb setup the db in the master
func (m *Master) setupDb() error {
	filename := m.dbFilename
	client, err := ent.Open(
		"sqlite3",
		filename+"?mode=rwc&cache=shared&&_busy_timeout=9999999&_fk=1")

	if err != nil {
		return fmt.Errorf("failed opening sqlite3 connection: %v", err)
	}

	if err = client.Schema.Create(context.Background()); err != nil {
		return fmt.Errorf("failed creating schema resources: %v", err)
	}

	m.db = client

	return nil
}

// schedule get a pending application from the db if there is no active job
func (m *Master) schedule() {
	for {
		ctx, cancel := context.WithCancel(context.Background())
		time.Sleep(1 * time.Second)

		// finding pending application
		app, err := m.nextApplication(ctx)
		if err != nil {
			continue
		}
		job := &job{
			app:    app,
			cancel: cancel,
		}
		m.run(ctx, job)
	}
}

func (m *Master) run(ctx context.Context, j *job) (err error) {
	// create new job from the application
	m.job = j

	defer func() {
		je := jobFinished

		// normalize je
		if err != nil {
			m.logger.Infow("failed run job",
				"application id", m.job.app.ID,
				"err", err,
			)

			m.LogEvent(ctx, m.job.app.ID, err.Error(), "master:run", "error", "")
			je = jobError

			if ctx.Err() != nil {
				je = jobCancel
				err = ErrAppIsCanceled
				m.LogEvent(ctx, m.job.app.ID, err.Error(), "master:run", "error", "")
			}
		}

		// create new context
		ctx := context.TODO()
		_ = m.jobTo(ctx, je)

		m.logger.Infow("job new status",
			"application id", m.job.app.ID,
			"status", m.job.app.Status,
		)
	}()

	m.logger.Infow("job new status",
		"application id", m.job.app.ID,
		"status", m.job.app.Status,
	)

	// change job to provisioning
	if err = m.jobTo(ctx, jobProvisioning); err != nil {
		return
	}

	m.logger.Infow("job new status",
		"application id", m.job.app.ID,
		"status", m.job.app.Status,
	)

	if err = m.jobCompile(ctx); err != nil {
		m.LogEvent(ctx, m.job.app.ID, err.Error(), "master:run:jobCompile", "error", "")
		return
	}
	// todo: ditribute the plugin to other worker when run in cloud mode
	// in this phase, the server run in local mode

	// change job to running state
	if err = m.jobTo(ctx, jobRunning); err != nil {
		return
	}

	m.logger.Infow("job new status",
		"application id", m.job.app.ID,
		"status", m.job.app.Status,
	)

	if err = m.runJob(ctx); err != nil {
		m.LogEvent(ctx, m.job.app.ID, err.Error(), "master:run", "error", "")
		return
	}

	return
}

// cancel terminates a running job with the same app ID
func (m *Master) cancel(ctx context.Context, appID int) error {
	if m.job == nil {
		return ErrAppNotRunning
	}
	if m.job.app.ID != appID {
		return ErrAppNotRunning
	}

	m.job.cancel()

	return nil
}

// provision compiles a scenario to golang plugin, distribute the application to
// worker. Return success when the workers confirm that the plugin is ready
func (m *Master) provision() (*ent.Application, error) {
	// compile
	return nil, nil
}

func (m *Master) nextApplication(ctx context.Context) (*ent.Application, error) {
	app, err := m.db.
		Application.
		Query().
		Where(
			application.Status(string(jobPending)),
		).
		Order(
			ent.Asc(application.FieldCreatedAt),
		).
		First(ctx)

	return app, err
}

// jobCompile using go to compile a scenario in plugin build mode
// the result is path to so file
func (m *Master) jobCompile(ctx context.Context) error {
	var path string

	scen := m.job.app.Scenario

	// save the scenario to a tmp file
	tmpScenF, err := ioutil.TempFile("", "gobench-scenario-*.go")
	if err != nil {
		return fmt.Errorf("failed creating temp scenario file: %v", err)
	}
	tmpScenName := tmpScenF.Name()

	defer os.Remove(tmpScenName) // cleanup

	_, err = tmpScenF.Write([]byte(scen))
	if err != nil {
		return fmt.Errorf("failed write to scenario file: %v", err)
	}

	if err = tmpScenF.Close(); err != nil {
		return fmt.Errorf("failed close the scenario file: %v", err)
	}

	path = fmt.Sprintf("%s.out", tmpScenName)

	// compile the scenario to a tmp file
	cmd := exec.Command("go", "build", "-buildmode=plugin", "-o",
		path, tmpScenName)

	// if out, err := cmd.CombinedOutput(); err != nil {
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed compiling the scenario: %v", err)
	}

	m.job.plugin = path

	return nil
}

// runJob run a application in a job
// by create a local worker
func (m *Master) runJob(ctx context.Context) (err error) {
	// todo: move this to agent
	driverPath := m.job.plugin
	appID := strconv.Itoa(m.job.app.ID)
	agentSock := m.la.GetSocketName()
	executorSock := fmt.Sprintf("/tmp/executorsock-%s", appID)

	cmd := exec.CommandContext(ctx, m.program,
		"--mode", "executor",
		"--agent-sock", agentSock,
		"--executor-sock", executorSock,
		"--driver-path", driverPath,
		"--app-id", appID)

	// get the stderr log
	stderr, err := cmd.StderrPipe()
	if err != nil {
		err = fmt.Errorf("cmd pipe stderr: %v", err)
		return
	}
	go func() {
		slurp, _ := ioutil.ReadAll(stderr)
		fmt.Printf("%s\n", string(slurp))
	}()

	// get the stdout log
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		err = fmt.Errorf("cmd pipe stdout: %v", err)
		return
	}
	go func() {
		slurp, _ := ioutil.ReadAll(stdout)
		fmt.Printf("%s\n", string(slurp))
	}()

	// start the cmd, does not wait for it to complete
	if err = cmd.Start(); err != nil {
		err = fmt.Errorf("executor start: %v", err)
		return
	}

	// waiting for the executor rpc to be ready
	b := time.Now()
	client, err := waitForReady(ctx, executorSock, 5*time.Second)
	if err != nil {
		err = fmt.Errorf("rpc dial: %v", err)
		return
	}
	m.logger.Infow("local executor is ready", "startup", time.Now().Sub(b))

	m.logger.Infow("local executor to run driver")

	req := true
	res := new(bool)
	if err = client.Call("Executor.Start", &req, &res); err != nil {
		err = fmt.Errorf("rpc start: %v", err)
		return
	}

	m.logger.Infow("local executor is shutting down")
	terReq := 0
	terRes := new(bool)
	// ignore error, since when the executor is terminated, this rpc will fail
	_ = client.Call("Executor.Terminate", &terReq, &terRes)

	if err = cmd.Wait(); err != nil {
		m.logger.Errorw("executor wait", "err", err)
		return
	}

	return nil
}

func waitForReady(ctx context.Context, executorSock string, expiredIn time.Duration) (
	*rpc.Client, error,
) {
	timeout := time.After(expiredIn)
	sleep := 10 * time.Millisecond
	for {
		time.Sleep(sleep)

		select {
		case <-ctx.Done():
			return nil, errors.New("cancel")
		case <-timeout:
			return nil, errors.New("timeout")
		default:
			client, err := rpc.DialHTTP("unix", executorSock)
			if err != nil {
				continue
			}
			return client, nil
		}
	}
}

// LogEvent log event for application
func (m *Master) LogEvent(ctx context.Context, appID int, message, source, level, name string) {

	q := m.db.EventLog.
		Create().
		SetMessage(message).
		SetSource(source)
	if appID != 0 {
		q.SetApplicationsID(appID)
	} else {
		q.SetApplicationsID(m.job.app.ID)
	}
	if name != "" {
		q.SetName(name)
	}
	if level != "" {
		q.SetLevel(level)
	}
	q.Save(ctx)
}
