// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/inigolabs/entbrowse/starwars/ent/gen/ent/film"
	"github.com/inigolabs/entbrowse/starwars/ent/gen/ent/person"
	"github.com/inigolabs/entbrowse/starwars/ent/gen/ent/planet"
	"github.com/inigolabs/entbrowse/starwars/ent/gen/ent/species"
)

// SpeciesCreate is the builder for creating a Species entity.
type SpeciesCreate struct {
	config
	mutation *SpeciesMutation
	hooks    []Hook
}

// SetAverageHeight sets the "average_height" field.
func (sc *SpeciesCreate) SetAverageHeight(i int) *SpeciesCreate {
	sc.mutation.SetAverageHeight(i)
	return sc
}

// SetAverageLifespan sets the "average_lifespan" field.
func (sc *SpeciesCreate) SetAverageLifespan(s string) *SpeciesCreate {
	sc.mutation.SetAverageLifespan(s)
	return sc
}

// SetClassification sets the "classification" field.
func (sc *SpeciesCreate) SetClassification(s string) *SpeciesCreate {
	sc.mutation.SetClassification(s)
	return sc
}

// SetDesignation sets the "designation" field.
func (sc *SpeciesCreate) SetDesignation(s string) *SpeciesCreate {
	sc.mutation.SetDesignation(s)
	return sc
}

// SetName sets the "name" field.
func (sc *SpeciesCreate) SetName(s string) *SpeciesCreate {
	sc.mutation.SetName(s)
	return sc
}

// SetSkinColor sets the "skin_color" field.
func (sc *SpeciesCreate) SetSkinColor(s string) *SpeciesCreate {
	sc.mutation.SetSkinColor(s)
	return sc
}

// SetEyeColor sets the "eye_color" field.
func (sc *SpeciesCreate) SetEyeColor(s string) *SpeciesCreate {
	sc.mutation.SetEyeColor(s)
	return sc
}

// SetHairColor sets the "hair_color" field.
func (sc *SpeciesCreate) SetHairColor(s string) *SpeciesCreate {
	sc.mutation.SetHairColor(s)
	return sc
}

// SetLanguage sets the "language" field.
func (sc *SpeciesCreate) SetLanguage(s string) *SpeciesCreate {
	sc.mutation.SetLanguage(s)
	return sc
}

// AddOriginatesFromIDs adds the "originates_from" edge to the Planet entity by IDs.
func (sc *SpeciesCreate) AddOriginatesFromIDs(ids ...int) *SpeciesCreate {
	sc.mutation.AddOriginatesFromIDs(ids...)
	return sc
}

// AddOriginatesFrom adds the "originates_from" edges to the Planet entity.
func (sc *SpeciesCreate) AddOriginatesFrom(p ...*Planet) *SpeciesCreate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return sc.AddOriginatesFromIDs(ids...)
}

// AddAppearedInIDs adds the "appeared_in" edge to the Film entity by IDs.
func (sc *SpeciesCreate) AddAppearedInIDs(ids ...int) *SpeciesCreate {
	sc.mutation.AddAppearedInIDs(ids...)
	return sc
}

// AddAppearedIn adds the "appeared_in" edges to the Film entity.
func (sc *SpeciesCreate) AddAppearedIn(f ...*Film) *SpeciesCreate {
	ids := make([]int, len(f))
	for i := range f {
		ids[i] = f[i].ID
	}
	return sc.AddAppearedInIDs(ids...)
}

// AddIncludesPersonIDs adds the "includes_person" edge to the Person entity by IDs.
func (sc *SpeciesCreate) AddIncludesPersonIDs(ids ...int) *SpeciesCreate {
	sc.mutation.AddIncludesPersonIDs(ids...)
	return sc
}

// AddIncludesPerson adds the "includes_person" edges to the Person entity.
func (sc *SpeciesCreate) AddIncludesPerson(p ...*Person) *SpeciesCreate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return sc.AddIncludesPersonIDs(ids...)
}

// Mutation returns the SpeciesMutation object of the builder.
func (sc *SpeciesCreate) Mutation() *SpeciesMutation {
	return sc.mutation
}

// Save creates the Species in the database.
func (sc *SpeciesCreate) Save(ctx context.Context) (*Species, error) {
	var (
		err  error
		node *Species
	)
	if len(sc.hooks) == 0 {
		if err = sc.check(); err != nil {
			return nil, err
		}
		node, err = sc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*SpeciesMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = sc.check(); err != nil {
				return nil, err
			}
			sc.mutation = mutation
			if node, err = sc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(sc.hooks) - 1; i >= 0; i-- {
			if sc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = sc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, sc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (sc *SpeciesCreate) SaveX(ctx context.Context) *Species {
	v, err := sc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (sc *SpeciesCreate) Exec(ctx context.Context) error {
	_, err := sc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (sc *SpeciesCreate) ExecX(ctx context.Context) {
	if err := sc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (sc *SpeciesCreate) check() error {
	if _, ok := sc.mutation.AverageHeight(); !ok {
		return &ValidationError{Name: "average_height", err: errors.New(`ent: missing required field "Species.average_height"`)}
	}
	if _, ok := sc.mutation.AverageLifespan(); !ok {
		return &ValidationError{Name: "average_lifespan", err: errors.New(`ent: missing required field "Species.average_lifespan"`)}
	}
	if _, ok := sc.mutation.Classification(); !ok {
		return &ValidationError{Name: "classification", err: errors.New(`ent: missing required field "Species.classification"`)}
	}
	if _, ok := sc.mutation.Designation(); !ok {
		return &ValidationError{Name: "designation", err: errors.New(`ent: missing required field "Species.designation"`)}
	}
	if _, ok := sc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "Species.name"`)}
	}
	if _, ok := sc.mutation.SkinColor(); !ok {
		return &ValidationError{Name: "skin_color", err: errors.New(`ent: missing required field "Species.skin_color"`)}
	}
	if _, ok := sc.mutation.EyeColor(); !ok {
		return &ValidationError{Name: "eye_color", err: errors.New(`ent: missing required field "Species.eye_color"`)}
	}
	if _, ok := sc.mutation.HairColor(); !ok {
		return &ValidationError{Name: "hair_color", err: errors.New(`ent: missing required field "Species.hair_color"`)}
	}
	if _, ok := sc.mutation.Language(); !ok {
		return &ValidationError{Name: "language", err: errors.New(`ent: missing required field "Species.language"`)}
	}
	return nil
}

func (sc *SpeciesCreate) sqlSave(ctx context.Context) (*Species, error) {
	_node, _spec := sc.createSpec()
	if err := sqlgraph.CreateNode(ctx, sc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (sc *SpeciesCreate) createSpec() (*Species, *sqlgraph.CreateSpec) {
	var (
		_node = &Species{config: sc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: species.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: species.FieldID,
			},
		}
	)
	if value, ok := sc.mutation.AverageHeight(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: species.FieldAverageHeight,
		})
		_node.AverageHeight = value
	}
	if value, ok := sc.mutation.AverageLifespan(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: species.FieldAverageLifespan,
		})
		_node.AverageLifespan = value
	}
	if value, ok := sc.mutation.Classification(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: species.FieldClassification,
		})
		_node.Classification = value
	}
	if value, ok := sc.mutation.Designation(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: species.FieldDesignation,
		})
		_node.Designation = value
	}
	if value, ok := sc.mutation.Name(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: species.FieldName,
		})
		_node.Name = value
	}
	if value, ok := sc.mutation.SkinColor(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: species.FieldSkinColor,
		})
		_node.SkinColor = value
	}
	if value, ok := sc.mutation.EyeColor(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: species.FieldEyeColor,
		})
		_node.EyeColor = value
	}
	if value, ok := sc.mutation.HairColor(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: species.FieldHairColor,
		})
		_node.HairColor = value
	}
	if value, ok := sc.mutation.Language(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: species.FieldLanguage,
		})
		_node.Language = value
	}
	if nodes := sc.mutation.OriginatesFromIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   species.OriginatesFromTable,
			Columns: species.OriginatesFromPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: planet.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := sc.mutation.AppearedInIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   species.AppearedInTable,
			Columns: species.AppearedInPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: film.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := sc.mutation.IncludesPersonIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   species.IncludesPersonTable,
			Columns: species.IncludesPersonPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: person.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// SpeciesCreateBulk is the builder for creating many Species entities in bulk.
type SpeciesCreateBulk struct {
	config
	builders []*SpeciesCreate
}

// Save creates the Species entities in the database.
func (scb *SpeciesCreateBulk) Save(ctx context.Context) ([]*Species, error) {
	specs := make([]*sqlgraph.CreateSpec, len(scb.builders))
	nodes := make([]*Species, len(scb.builders))
	mutators := make([]Mutator, len(scb.builders))
	for i := range scb.builders {
		func(i int, root context.Context) {
			builder := scb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*SpeciesMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, scb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, scb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{err.Error(), err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, scb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (scb *SpeciesCreateBulk) SaveX(ctx context.Context) []*Species {
	v, err := scb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (scb *SpeciesCreateBulk) Exec(ctx context.Context) error {
	_, err := scb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (scb *SpeciesCreateBulk) ExecX(ctx context.Context) {
	if err := scb.Exec(ctx); err != nil {
		panic(err)
	}
}