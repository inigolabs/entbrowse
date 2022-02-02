// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/inigolabs/entbrowse/starwars/ent/gen/ent/person"
	"github.com/inigolabs/entbrowse/starwars/ent/gen/ent/predicate"
)

// PersonDelete is the builder for deleting a Person entity.
type PersonDelete struct {
	config
	hooks    []Hook
	mutation *PersonMutation
}

// Where appends a list predicates to the PersonDelete builder.
func (pd *PersonDelete) Where(ps ...predicate.Person) *PersonDelete {
	pd.mutation.Where(ps...)
	return pd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (pd *PersonDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(pd.hooks) == 0 {
		affected, err = pd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*PersonMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			pd.mutation = mutation
			affected, err = pd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(pd.hooks) - 1; i >= 0; i-- {
			if pd.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = pd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, pd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (pd *PersonDelete) ExecX(ctx context.Context) int {
	n, err := pd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (pd *PersonDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: person.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: person.FieldID,
			},
		},
	}
	if ps := pd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, pd.driver, _spec)
}

// PersonDeleteOne is the builder for deleting a single Person entity.
type PersonDeleteOne struct {
	pd *PersonDelete
}

// Exec executes the deletion query.
func (pdo *PersonDeleteOne) Exec(ctx context.Context) error {
	n, err := pdo.pd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{person.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (pdo *PersonDeleteOne) ExecX(ctx context.Context) {
	pdo.pd.ExecX(ctx)
}
