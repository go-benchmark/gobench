package web

import (
	"net/http"

	"github.com/go-chi/render"
	"github.com/gobench-io/gobench/ent"
)

// Err is the error struct that compatible to Google API recommendation
// https://cloud.google.com/apis/design/errors#error_model
type Err struct {
	Code    int    `json:"code,omitempty"`    // application-specific error code
	Message string `json:"message,omitempty"` // application-level error message, for debugging
	Status  string `json:"status"`            // user-level status message
}

// ErrResponse is the error struct that compatible to Google API recommendation
// https://cloud.google.com/apis/design/errors#error_model
type ErrResponse struct {
	Error Err `json:"error"`
}

func (e *ErrResponse) Render(w http.ResponseWriter, r *http.Request) error {
	render.Status(r, e.Error.Code)
	return nil
}

func ErrInternalServer(err error) render.Renderer {
	return &ErrResponse{
		Error: Err{
			Code:    500,
			Message: err.Error(),
			Status:  "Internal Server Error",
		},
	}
}

func ErrInvalidRequest(err error) render.Renderer {
	return &ErrResponse{
		Error: Err{
			Code:    400,
			Message: err.Error(),
			Status:  "Invalid Request",
		},
	}
}

func ErrNotFoundRequest(err error) render.Renderer {
	return &ErrResponse{
		Error: Err{
			Code:    404,
			Message: "Request data not found",
			Status:  "Model Not Found",
		},
	}
}

func ErrRender(err error) render.Renderer {
	return &ErrResponse{
		Error: Err{
			Code:    422,
			Message: err.Error(),
			Status:  "Error Rendering Response.",
		},
	}
}

// application response
type applicationRequest struct {
	*ent.Application
	ProtectedID int `json:"id"`
}

func (a *applicationRequest) Bind(r *http.Request) (err error) {
	return nil
}

type applicationResponse struct {
	*ent.Application
	Edges *struct{} `json:"edges,omitempty"`
}

func (ar *applicationResponse) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}
func newApplicationResponse(a *ent.Application) *applicationResponse {
	return &applicationResponse{
		a,
		nil,
	}
}
func newApplicationListResponse(aps []*ent.Application) []render.Renderer {
	list := []render.Renderer{}
	for _, ap := range aps {
		list = append(list, newApplicationResponse(ap))
	}
	return list
}

// group response
type groupResponse struct {
	*ent.Group
	Edges *struct{} `json:"edges,omitempty"`
}

func (gr *groupResponse) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}
func newGroupResponse(g *ent.Group) *groupResponse {
	return &groupResponse{
		g,
		nil,
	}
}
func newGroupListResponse(gs []*ent.Group) []render.Renderer {
	list := []render.Renderer{}
	for _, g := range gs {
		list = append(list, newGroupResponse(g))
	}
	return list
}

// graph response
type graphResponse struct {
	*ent.Graph
	Edges *struct{} `json:"edges,omitempty"`
}

func (gr *graphResponse) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}
func newGraphResponse(g *ent.Graph) *graphResponse {
	return &graphResponse{
		g,
		nil,
	}
}
func newGraphListResponse(gs []*ent.Graph) []render.Renderer {
	list := []render.Renderer{}
	for _, g := range gs {
		list = append(list, newGraphResponse(g))
	}
	return list
}

// metric response
type metricResponse struct {
	*ent.Metric
	Edges *struct{} `json:"edges,omitempty"`
}

func (gr *metricResponse) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}
func newMetricResponse(m *ent.Metric) *metricResponse {
	return &metricResponse{
		m,
		nil,
	}
}
func newMetricListResponse(ms []*ent.Metric) []render.Renderer {
	list := []render.Renderer{}
	for _, m := range ms {
		list = append(list, newMetricResponse(m))
	}
	return list
}

// counter response
type counterResponse struct {
	*ent.Counter
	Count int64     `json:"count"`
	Edges *struct{} `json:"edges,omitempty"`
}

func (gr *counterResponse) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}
func newCounterResponse(e *ent.Counter) *counterResponse {
	return &counterResponse{
		e,
		e.Count,
		nil,
	}
}
func newCounterListResponse(es []*ent.Counter) []render.Renderer {
	list := []render.Renderer{}
	for _, e := range es {
		list = append(list, newCounterResponse(e))
	}
	return list
}

// histogram response
type histogramResponse struct {
	*ent.Histogram
	Edges *struct{} `json:"edges,omitempty"`
}

func (gr *histogramResponse) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}
func newHistogramResponse(e *ent.Histogram) *histogramResponse {
	return &histogramResponse{
		e,
		nil,
	}
}
func newHistogramListResponse(es []*ent.Histogram) []render.Renderer {
	list := []render.Renderer{}
	for _, e := range es {
		list = append(list, newHistogramResponse(e))
	}
	return list
}

// gauge response
type gaugeResponse struct {
	*ent.Gauge
	Edges *struct{} `json:"edges,omitempty"`
}

func (gr *gaugeResponse) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}
func newGaugeResponse(e *ent.Gauge) *gaugeResponse {
	return &gaugeResponse{
		e,
		nil,
	}
}
func newGaugeListResponse(es []*ent.Gauge) []render.Renderer {
	list := []render.Renderer{}
	for _, e := range es {
		list = append(list, newGaugeResponse(e))
	}
	return list
}
