// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"github.com/facebook/ent/dialect/sql"
	"github.com/gobench-io/gobench/ent/graph"
	"github.com/gobench-io/gobench/ent/group"
)

// Graph is the model entity for the Graph schema.
type Graph struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Title holds the value of the "title" field.
	Title string `json:"title"`
	// Unit holds the value of the "unit" field.
	Unit string `json:"unit"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the GraphQuery when eager-loading is set.
	Edges        GraphEdges `json:"edges"`
	group_graphs *int
}

// GraphEdges holds the relations/edges for other nodes in the graph.
type GraphEdges struct {
	// Group holds the value of the group edge.
	Group *Group
	// Metrics holds the value of the metrics edge.
	Metrics []*Metric
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// GroupOrErr returns the Group value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e GraphEdges) GroupOrErr() (*Group, error) {
	if e.loadedTypes[0] {
		if e.Group == nil {
			// The edge group was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: group.Label}
		}
		return e.Group, nil
	}
	return nil, &NotLoadedError{edge: "group"}
}

// MetricsOrErr returns the Metrics value or an error if the edge
// was not loaded in eager-loading.
func (e GraphEdges) MetricsOrErr() ([]*Metric, error) {
	if e.loadedTypes[1] {
		return e.Metrics, nil
	}
	return nil, &NotLoadedError{edge: "metrics"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Graph) scanValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{},  // id
		&sql.NullString{}, // title
		&sql.NullString{}, // unit
	}
}

// fkValues returns the types for scanning foreign-keys values from sql.Rows.
func (*Graph) fkValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{}, // group_graphs
	}
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Graph fields.
func (gr *Graph) assignValues(values ...interface{}) error {
	if m, n := len(values), len(graph.Columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	value, ok := values[0].(*sql.NullInt64)
	if !ok {
		return fmt.Errorf("unexpected type %T for field id", value)
	}
	gr.ID = int(value.Int64)
	values = values[1:]
	if value, ok := values[0].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field title", values[0])
	} else if value.Valid {
		gr.Title = value.String
	}
	if value, ok := values[1].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field unit", values[1])
	} else if value.Valid {
		gr.Unit = value.String
	}
	values = values[2:]
	if len(values) == len(graph.ForeignKeys) {
		if value, ok := values[0].(*sql.NullInt64); !ok {
			return fmt.Errorf("unexpected type %T for edge-field group_graphs", value)
		} else if value.Valid {
			gr.group_graphs = new(int)
			*gr.group_graphs = int(value.Int64)
		}
	}
	return nil
}

// QueryGroup queries the group edge of the Graph.
func (gr *Graph) QueryGroup() *GroupQuery {
	return (&GraphClient{config: gr.config}).QueryGroup(gr)
}

// QueryMetrics queries the metrics edge of the Graph.
func (gr *Graph) QueryMetrics() *MetricQuery {
	return (&GraphClient{config: gr.config}).QueryMetrics(gr)
}

// Update returns a builder for updating this Graph.
// Note that, you need to call Graph.Unwrap() before calling this method, if this Graph
// was returned from a transaction, and the transaction was committed or rolled back.
func (gr *Graph) Update() *GraphUpdateOne {
	return (&GraphClient{config: gr.config}).UpdateOne(gr)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (gr *Graph) Unwrap() *Graph {
	tx, ok := gr.config.driver.(*txDriver)
	if !ok {
		panic("ent: Graph is not a transactional entity")
	}
	gr.config.driver = tx.drv
	return gr
}

// String implements the fmt.Stringer.
func (gr *Graph) String() string {
	var builder strings.Builder
	builder.WriteString("Graph(")
	builder.WriteString(fmt.Sprintf("id=%v", gr.ID))
	builder.WriteString(", title=")
	builder.WriteString(gr.Title)
	builder.WriteString(", unit=")
	builder.WriteString(gr.Unit)
	builder.WriteByte(')')
	return builder.String()
}

// Graphs is a parsable slice of Graph.
type Graphs []*Graph

func (gr Graphs) config(cfg config) {
	for _i := range gr {
		gr[_i].config = cfg
	}
}
