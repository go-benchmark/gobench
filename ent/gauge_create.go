// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/gobench-io/gobench/ent/gauge"
	"github.com/gobench-io/gobench/ent/metric"
)

// GaugeCreate is the builder for creating a Gauge entity.
type GaugeCreate struct {
	config
	mutation *GaugeMutation
	hooks    []Hook
}

// SetTime sets the time field.
func (gc *GaugeCreate) SetTime(i int64) *GaugeCreate {
	gc.mutation.SetTime(i)
	return gc
}

// SetValue sets the value field.
func (gc *GaugeCreate) SetValue(i int64) *GaugeCreate {
	gc.mutation.SetValue(i)
	return gc
}

// SetWID sets the wID field.
func (gc *GaugeCreate) SetWID(s string) *GaugeCreate {
	gc.mutation.SetWID(s)
	return gc
}

// SetMetricID sets the metric edge to Metric by id.
func (gc *GaugeCreate) SetMetricID(id int) *GaugeCreate {
	gc.mutation.SetMetricID(id)
	return gc
}

// SetNillableMetricID sets the metric edge to Metric by id if the given value is not nil.
func (gc *GaugeCreate) SetNillableMetricID(id *int) *GaugeCreate {
	if id != nil {
		gc = gc.SetMetricID(*id)
	}
	return gc
}

// SetMetric sets the metric edge to Metric.
func (gc *GaugeCreate) SetMetric(m *Metric) *GaugeCreate {
	return gc.SetMetricID(m.ID)
}

// Save creates the Gauge in the database.
func (gc *GaugeCreate) Save(ctx context.Context) (*Gauge, error) {
	if _, ok := gc.mutation.Time(); !ok {
		return nil, errors.New("ent: missing required field \"time\"")
	}
	if _, ok := gc.mutation.Value(); !ok {
		return nil, errors.New("ent: missing required field \"value\"")
	}
	if _, ok := gc.mutation.WID(); !ok {
		return nil, errors.New("ent: missing required field \"wID\"")
	}
	var (
		err  error
		node *Gauge
	)
	if len(gc.hooks) == 0 {
		node, err = gc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*GaugeMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			gc.mutation = mutation
			node, err = gc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(gc.hooks) - 1; i >= 0; i-- {
			mut = gc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, gc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (gc *GaugeCreate) SaveX(ctx context.Context) *Gauge {
	v, err := gc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (gc *GaugeCreate) sqlSave(ctx context.Context) (*Gauge, error) {
	var (
		ga    = &Gauge{config: gc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: gauge.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: gauge.FieldID,
			},
		}
	)
	if value, ok := gc.mutation.Time(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: gauge.FieldTime,
		})
		ga.Time = value
	}
	if value, ok := gc.mutation.Value(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: gauge.FieldValue,
		})
		ga.Value = value
	}
	if value, ok := gc.mutation.WID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: gauge.FieldWID,
		})
		ga.WID = value
	}
	if nodes := gc.mutation.MetricIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   gauge.MetricTable,
			Columns: []string{gauge.MetricColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: metric.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if err := sqlgraph.CreateNode(ctx, gc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	ga.ID = int(id)
	return ga, nil
}
