// Code generated by entc, DO NOT EDIT.

package eventlog

import (
	"time"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/gobench-io/gobench/ent/predicate"
)

// ID filters vertices based on their identifier.
func ID(id int) predicate.EventLog {
	return predicate.EventLog(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.EventLog {
	return predicate.EventLog(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.EventLog {
	return predicate.EventLog(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.EventLog {
	return predicate.EventLog(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.EventLog {
	return predicate.EventLog(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.EventLog {
	return predicate.EventLog(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.EventLog {
	return predicate.EventLog(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.EventLog {
	return predicate.EventLog(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.EventLog {
	return predicate.EventLog(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.EventLog {
	return predicate.EventLog(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldName), v))
	})
}

// Message applies equality check predicate on the "message" field. It's identical to MessageEQ.
func Message(v string) predicate.EventLog {
	return predicate.EventLog(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldMessage), v))
	})
}

// Level applies equality check predicate on the "level" field. It's identical to LevelEQ.
func Level(v string) predicate.EventLog {
	return predicate.EventLog(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldLevel), v))
	})
}

// Source applies equality check predicate on the "source" field. It's identical to SourceEQ.
func Source(v string) predicate.EventLog {
	return predicate.EventLog(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSource), v))
	})
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.EventLog {
	return predicate.EventLog(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.EventLog {
	return predicate.EventLog(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldName), v))
	})
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.EventLog {
	return predicate.EventLog(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldName), v))
	})
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.EventLog {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.EventLog(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(vs) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldName), v...))
	})
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.EventLog {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.EventLog(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(vs) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldName), v...))
	})
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.EventLog {
	return predicate.EventLog(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldName), v))
	})
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.EventLog {
	return predicate.EventLog(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldName), v))
	})
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.EventLog {
	return predicate.EventLog(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldName), v))
	})
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.EventLog {
	return predicate.EventLog(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldName), v))
	})
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.EventLog {
	return predicate.EventLog(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldName), v))
	})
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.EventLog {
	return predicate.EventLog(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldName), v))
	})
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.EventLog {
	return predicate.EventLog(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldName), v))
	})
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.EventLog {
	return predicate.EventLog(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldName), v))
	})
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.EventLog {
	return predicate.EventLog(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldName), v))
	})
}

// MessageEQ applies the EQ predicate on the "message" field.
func MessageEQ(v string) predicate.EventLog {
	return predicate.EventLog(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldMessage), v))
	})
}

// MessageNEQ applies the NEQ predicate on the "message" field.
func MessageNEQ(v string) predicate.EventLog {
	return predicate.EventLog(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldMessage), v))
	})
}

// MessageIn applies the In predicate on the "message" field.
func MessageIn(vs ...string) predicate.EventLog {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.EventLog(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(vs) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldMessage), v...))
	})
}

// MessageNotIn applies the NotIn predicate on the "message" field.
func MessageNotIn(vs ...string) predicate.EventLog {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.EventLog(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(vs) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldMessage), v...))
	})
}

// MessageGT applies the GT predicate on the "message" field.
func MessageGT(v string) predicate.EventLog {
	return predicate.EventLog(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldMessage), v))
	})
}

// MessageGTE applies the GTE predicate on the "message" field.
func MessageGTE(v string) predicate.EventLog {
	return predicate.EventLog(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldMessage), v))
	})
}

// MessageLT applies the LT predicate on the "message" field.
func MessageLT(v string) predicate.EventLog {
	return predicate.EventLog(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldMessage), v))
	})
}

// MessageLTE applies the LTE predicate on the "message" field.
func MessageLTE(v string) predicate.EventLog {
	return predicate.EventLog(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldMessage), v))
	})
}

// MessageContains applies the Contains predicate on the "message" field.
func MessageContains(v string) predicate.EventLog {
	return predicate.EventLog(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldMessage), v))
	})
}

// MessageHasPrefix applies the HasPrefix predicate on the "message" field.
func MessageHasPrefix(v string) predicate.EventLog {
	return predicate.EventLog(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldMessage), v))
	})
}

// MessageHasSuffix applies the HasSuffix predicate on the "message" field.
func MessageHasSuffix(v string) predicate.EventLog {
	return predicate.EventLog(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldMessage), v))
	})
}

// MessageEqualFold applies the EqualFold predicate on the "message" field.
func MessageEqualFold(v string) predicate.EventLog {
	return predicate.EventLog(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldMessage), v))
	})
}

// MessageContainsFold applies the ContainsFold predicate on the "message" field.
func MessageContainsFold(v string) predicate.EventLog {
	return predicate.EventLog(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldMessage), v))
	})
}

// LevelEQ applies the EQ predicate on the "level" field.
func LevelEQ(v string) predicate.EventLog {
	return predicate.EventLog(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldLevel), v))
	})
}

// LevelNEQ applies the NEQ predicate on the "level" field.
func LevelNEQ(v string) predicate.EventLog {
	return predicate.EventLog(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldLevel), v))
	})
}

// LevelIn applies the In predicate on the "level" field.
func LevelIn(vs ...string) predicate.EventLog {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.EventLog(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(vs) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldLevel), v...))
	})
}

// LevelNotIn applies the NotIn predicate on the "level" field.
func LevelNotIn(vs ...string) predicate.EventLog {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.EventLog(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(vs) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldLevel), v...))
	})
}

// LevelGT applies the GT predicate on the "level" field.
func LevelGT(v string) predicate.EventLog {
	return predicate.EventLog(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldLevel), v))
	})
}

// LevelGTE applies the GTE predicate on the "level" field.
func LevelGTE(v string) predicate.EventLog {
	return predicate.EventLog(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldLevel), v))
	})
}

// LevelLT applies the LT predicate on the "level" field.
func LevelLT(v string) predicate.EventLog {
	return predicate.EventLog(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldLevel), v))
	})
}

// LevelLTE applies the LTE predicate on the "level" field.
func LevelLTE(v string) predicate.EventLog {
	return predicate.EventLog(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldLevel), v))
	})
}

// LevelContains applies the Contains predicate on the "level" field.
func LevelContains(v string) predicate.EventLog {
	return predicate.EventLog(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldLevel), v))
	})
}

// LevelHasPrefix applies the HasPrefix predicate on the "level" field.
func LevelHasPrefix(v string) predicate.EventLog {
	return predicate.EventLog(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldLevel), v))
	})
}

// LevelHasSuffix applies the HasSuffix predicate on the "level" field.
func LevelHasSuffix(v string) predicate.EventLog {
	return predicate.EventLog(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldLevel), v))
	})
}

// LevelEqualFold applies the EqualFold predicate on the "level" field.
func LevelEqualFold(v string) predicate.EventLog {
	return predicate.EventLog(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldLevel), v))
	})
}

// LevelContainsFold applies the ContainsFold predicate on the "level" field.
func LevelContainsFold(v string) predicate.EventLog {
	return predicate.EventLog(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldLevel), v))
	})
}

// SourceEQ applies the EQ predicate on the "source" field.
func SourceEQ(v string) predicate.EventLog {
	return predicate.EventLog(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSource), v))
	})
}

// SourceNEQ applies the NEQ predicate on the "source" field.
func SourceNEQ(v string) predicate.EventLog {
	return predicate.EventLog(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldSource), v))
	})
}

// SourceIn applies the In predicate on the "source" field.
func SourceIn(vs ...string) predicate.EventLog {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.EventLog(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(vs) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldSource), v...))
	})
}

// SourceNotIn applies the NotIn predicate on the "source" field.
func SourceNotIn(vs ...string) predicate.EventLog {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.EventLog(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(vs) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldSource), v...))
	})
}

// SourceGT applies the GT predicate on the "source" field.
func SourceGT(v string) predicate.EventLog {
	return predicate.EventLog(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldSource), v))
	})
}

// SourceGTE applies the GTE predicate on the "source" field.
func SourceGTE(v string) predicate.EventLog {
	return predicate.EventLog(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldSource), v))
	})
}

// SourceLT applies the LT predicate on the "source" field.
func SourceLT(v string) predicate.EventLog {
	return predicate.EventLog(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldSource), v))
	})
}

// SourceLTE applies the LTE predicate on the "source" field.
func SourceLTE(v string) predicate.EventLog {
	return predicate.EventLog(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldSource), v))
	})
}

// SourceContains applies the Contains predicate on the "source" field.
func SourceContains(v string) predicate.EventLog {
	return predicate.EventLog(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldSource), v))
	})
}

// SourceHasPrefix applies the HasPrefix predicate on the "source" field.
func SourceHasPrefix(v string) predicate.EventLog {
	return predicate.EventLog(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldSource), v))
	})
}

// SourceHasSuffix applies the HasSuffix predicate on the "source" field.
func SourceHasSuffix(v string) predicate.EventLog {
	return predicate.EventLog(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldSource), v))
	})
}

// SourceEqualFold applies the EqualFold predicate on the "source" field.
func SourceEqualFold(v string) predicate.EventLog {
	return predicate.EventLog(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldSource), v))
	})
}

// SourceContainsFold applies the ContainsFold predicate on the "source" field.
func SourceContainsFold(v string) predicate.EventLog {
	return predicate.EventLog(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldSource), v))
	})
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.EventLog {
	return predicate.EventLog(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.EventLog {
	return predicate.EventLog(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.EventLog {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.EventLog(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(vs) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.EventLog {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.EventLog(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(vs) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.EventLog {
	return predicate.EventLog(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.EventLog {
	return predicate.EventLog(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.EventLog {
	return predicate.EventLog(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.EventLog {
	return predicate.EventLog(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreatedAt), v))
	})
}

// HasApplications applies the HasEdge predicate on the "applications" edge.
func HasApplications() predicate.EventLog {
	return predicate.EventLog(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ApplicationsTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, ApplicationsTable, ApplicationsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasApplicationsWith applies the HasEdge predicate on the "applications" edge with a given conditions (other predicates).
func HasApplicationsWith(preds ...predicate.Application) predicate.EventLog {
	return predicate.EventLog(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ApplicationsInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, ApplicationsTable, ApplicationsColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups list of predicates with the AND operator between them.
func And(predicates ...predicate.EventLog) predicate.EventLog {
	return predicate.EventLog(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups list of predicates with the OR operator between them.
func Or(predicates ...predicate.EventLog) predicate.EventLog {
	return predicate.EventLog(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.EventLog) predicate.EventLog {
	return predicate.EventLog(func(s *sql.Selector) {
		p(s.Not())
	})
}