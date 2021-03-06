// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"errors"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/inigolabs/entbrowse/starwars/ent/gen/ent/film"
	"github.com/inigolabs/entbrowse/starwars/ent/gen/ent/person"
	"github.com/inigolabs/entbrowse/starwars/ent/gen/ent/planet"
	"github.com/inigolabs/entbrowse/starwars/ent/gen/ent/predicate"
	"github.com/inigolabs/entbrowse/starwars/ent/gen/ent/species"
	"github.com/inigolabs/entbrowse/starwars/ent/gen/ent/starship"
	"github.com/inigolabs/entbrowse/starwars/ent/gen/ent/vehicle"
)

// FilmQuery is the builder for querying Film entities.
type FilmQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.Film
	// eager-loading edges.
	withHasPerson   *PersonQuery
	withHasPlanet   *PlanetQuery
	withHasStarship *StarshipQuery
	withHasVehicle  *VehicleQuery
	withHasSpecies  *SpeciesQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the FilmQuery builder.
func (fq *FilmQuery) Where(ps ...predicate.Film) *FilmQuery {
	fq.predicates = append(fq.predicates, ps...)
	return fq
}

// Limit adds a limit step to the query.
func (fq *FilmQuery) Limit(limit int) *FilmQuery {
	fq.limit = &limit
	return fq
}

// Offset adds an offset step to the query.
func (fq *FilmQuery) Offset(offset int) *FilmQuery {
	fq.offset = &offset
	return fq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (fq *FilmQuery) Unique(unique bool) *FilmQuery {
	fq.unique = &unique
	return fq
}

// Order adds an order step to the query.
func (fq *FilmQuery) Order(o ...OrderFunc) *FilmQuery {
	fq.order = append(fq.order, o...)
	return fq
}

// QueryHasPerson chains the current query on the "has_person" edge.
func (fq *FilmQuery) QueryHasPerson() *PersonQuery {
	query := &PersonQuery{config: fq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := fq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := fq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(film.Table, film.FieldID, selector),
			sqlgraph.To(person.Table, person.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, film.HasPersonTable, film.HasPersonPrimaryKey...),
		)
		fromU = sqlgraph.SetNeighbors(fq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryHasPlanet chains the current query on the "has_planet" edge.
func (fq *FilmQuery) QueryHasPlanet() *PlanetQuery {
	query := &PlanetQuery{config: fq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := fq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := fq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(film.Table, film.FieldID, selector),
			sqlgraph.To(planet.Table, planet.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, film.HasPlanetTable, film.HasPlanetPrimaryKey...),
		)
		fromU = sqlgraph.SetNeighbors(fq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryHasStarship chains the current query on the "has_starship" edge.
func (fq *FilmQuery) QueryHasStarship() *StarshipQuery {
	query := &StarshipQuery{config: fq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := fq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := fq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(film.Table, film.FieldID, selector),
			sqlgraph.To(starship.Table, starship.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, film.HasStarshipTable, film.HasStarshipPrimaryKey...),
		)
		fromU = sqlgraph.SetNeighbors(fq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryHasVehicle chains the current query on the "has_vehicle" edge.
func (fq *FilmQuery) QueryHasVehicle() *VehicleQuery {
	query := &VehicleQuery{config: fq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := fq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := fq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(film.Table, film.FieldID, selector),
			sqlgraph.To(vehicle.Table, vehicle.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, film.HasVehicleTable, film.HasVehiclePrimaryKey...),
		)
		fromU = sqlgraph.SetNeighbors(fq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryHasSpecies chains the current query on the "has_species" edge.
func (fq *FilmQuery) QueryHasSpecies() *SpeciesQuery {
	query := &SpeciesQuery{config: fq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := fq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := fq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(film.Table, film.FieldID, selector),
			sqlgraph.To(species.Table, species.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, film.HasSpeciesTable, film.HasSpeciesPrimaryKey...),
		)
		fromU = sqlgraph.SetNeighbors(fq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Film entity from the query.
// Returns a *NotFoundError when no Film was found.
func (fq *FilmQuery) First(ctx context.Context) (*Film, error) {
	nodes, err := fq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{film.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (fq *FilmQuery) FirstX(ctx context.Context) *Film {
	node, err := fq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Film ID from the query.
// Returns a *NotFoundError when no Film ID was found.
func (fq *FilmQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = fq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{film.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (fq *FilmQuery) FirstIDX(ctx context.Context) int {
	id, err := fq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Film entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when exactly one Film entity is not found.
// Returns a *NotFoundError when no Film entities are found.
func (fq *FilmQuery) Only(ctx context.Context) (*Film, error) {
	nodes, err := fq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{film.Label}
	default:
		return nil, &NotSingularError{film.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (fq *FilmQuery) OnlyX(ctx context.Context) *Film {
	node, err := fq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Film ID in the query.
// Returns a *NotSingularError when exactly one Film ID is not found.
// Returns a *NotFoundError when no entities are found.
func (fq *FilmQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = fq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{film.Label}
	default:
		err = &NotSingularError{film.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (fq *FilmQuery) OnlyIDX(ctx context.Context) int {
	id, err := fq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Films.
func (fq *FilmQuery) All(ctx context.Context) ([]*Film, error) {
	if err := fq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return fq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (fq *FilmQuery) AllX(ctx context.Context) []*Film {
	nodes, err := fq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Film IDs.
func (fq *FilmQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := fq.Select(film.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (fq *FilmQuery) IDsX(ctx context.Context) []int {
	ids, err := fq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (fq *FilmQuery) Count(ctx context.Context) (int, error) {
	if err := fq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return fq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (fq *FilmQuery) CountX(ctx context.Context) int {
	count, err := fq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (fq *FilmQuery) Exist(ctx context.Context) (bool, error) {
	if err := fq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return fq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (fq *FilmQuery) ExistX(ctx context.Context) bool {
	exist, err := fq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the FilmQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (fq *FilmQuery) Clone() *FilmQuery {
	if fq == nil {
		return nil
	}
	return &FilmQuery{
		config:          fq.config,
		limit:           fq.limit,
		offset:          fq.offset,
		order:           append([]OrderFunc{}, fq.order...),
		predicates:      append([]predicate.Film{}, fq.predicates...),
		withHasPerson:   fq.withHasPerson.Clone(),
		withHasPlanet:   fq.withHasPlanet.Clone(),
		withHasStarship: fq.withHasStarship.Clone(),
		withHasVehicle:  fq.withHasVehicle.Clone(),
		withHasSpecies:  fq.withHasSpecies.Clone(),
		// clone intermediate query.
		sql:  fq.sql.Clone(),
		path: fq.path,
	}
}

// WithHasPerson tells the query-builder to eager-load the nodes that are connected to
// the "has_person" edge. The optional arguments are used to configure the query builder of the edge.
func (fq *FilmQuery) WithHasPerson(opts ...func(*PersonQuery)) *FilmQuery {
	query := &PersonQuery{config: fq.config}
	for _, opt := range opts {
		opt(query)
	}
	fq.withHasPerson = query
	return fq
}

// WithHasPlanet tells the query-builder to eager-load the nodes that are connected to
// the "has_planet" edge. The optional arguments are used to configure the query builder of the edge.
func (fq *FilmQuery) WithHasPlanet(opts ...func(*PlanetQuery)) *FilmQuery {
	query := &PlanetQuery{config: fq.config}
	for _, opt := range opts {
		opt(query)
	}
	fq.withHasPlanet = query
	return fq
}

// WithHasStarship tells the query-builder to eager-load the nodes that are connected to
// the "has_starship" edge. The optional arguments are used to configure the query builder of the edge.
func (fq *FilmQuery) WithHasStarship(opts ...func(*StarshipQuery)) *FilmQuery {
	query := &StarshipQuery{config: fq.config}
	for _, opt := range opts {
		opt(query)
	}
	fq.withHasStarship = query
	return fq
}

// WithHasVehicle tells the query-builder to eager-load the nodes that are connected to
// the "has_vehicle" edge. The optional arguments are used to configure the query builder of the edge.
func (fq *FilmQuery) WithHasVehicle(opts ...func(*VehicleQuery)) *FilmQuery {
	query := &VehicleQuery{config: fq.config}
	for _, opt := range opts {
		opt(query)
	}
	fq.withHasVehicle = query
	return fq
}

// WithHasSpecies tells the query-builder to eager-load the nodes that are connected to
// the "has_species" edge. The optional arguments are used to configure the query builder of the edge.
func (fq *FilmQuery) WithHasSpecies(opts ...func(*SpeciesQuery)) *FilmQuery {
	query := &SpeciesQuery{config: fq.config}
	for _, opt := range opts {
		opt(query)
	}
	fq.withHasSpecies = query
	return fq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Director string `json:"director,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Film.Query().
//		GroupBy(film.FieldDirector).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (fq *FilmQuery) GroupBy(field string, fields ...string) *FilmGroupBy {
	group := &FilmGroupBy{config: fq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := fq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return fq.sqlQuery(ctx), nil
	}
	return group
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Director string `json:"director,omitempty"`
//	}
//
//	client.Film.Query().
//		Select(film.FieldDirector).
//		Scan(ctx, &v)
//
func (fq *FilmQuery) Select(fields ...string) *FilmSelect {
	fq.fields = append(fq.fields, fields...)
	return &FilmSelect{FilmQuery: fq}
}

func (fq *FilmQuery) prepareQuery(ctx context.Context) error {
	for _, f := range fq.fields {
		if !film.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if fq.path != nil {
		prev, err := fq.path(ctx)
		if err != nil {
			return err
		}
		fq.sql = prev
	}
	return nil
}

func (fq *FilmQuery) sqlAll(ctx context.Context) ([]*Film, error) {
	var (
		nodes       = []*Film{}
		_spec       = fq.querySpec()
		loadedTypes = [5]bool{
			fq.withHasPerson != nil,
			fq.withHasPlanet != nil,
			fq.withHasStarship != nil,
			fq.withHasVehicle != nil,
			fq.withHasSpecies != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		node := &Film{config: fq.config}
		nodes = append(nodes, node)
		return node.scanValues(columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		if len(nodes) == 0 {
			return fmt.Errorf("ent: Assign called without calling ScanValues")
		}
		node := nodes[len(nodes)-1]
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if err := sqlgraph.QueryNodes(ctx, fq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := fq.withHasPerson; query != nil {
		fks := make([]driver.Value, 0, len(nodes))
		ids := make(map[int]*Film, len(nodes))
		for _, node := range nodes {
			ids[node.ID] = node
			fks = append(fks, node.ID)
			node.Edges.HasPerson = []*Person{}
		}
		var (
			edgeids []int
			edges   = make(map[int][]*Film)
		)
		_spec := &sqlgraph.EdgeQuerySpec{
			Edge: &sqlgraph.EdgeSpec{
				Inverse: false,
				Table:   film.HasPersonTable,
				Columns: film.HasPersonPrimaryKey,
			},
			Predicate: func(s *sql.Selector) {
				s.Where(sql.InValues(film.HasPersonPrimaryKey[0], fks...))
			},
			ScanValues: func() [2]interface{} {
				return [2]interface{}{new(sql.NullInt64), new(sql.NullInt64)}
			},
			Assign: func(out, in interface{}) error {
				eout, ok := out.(*sql.NullInt64)
				if !ok || eout == nil {
					return fmt.Errorf("unexpected id value for edge-out")
				}
				ein, ok := in.(*sql.NullInt64)
				if !ok || ein == nil {
					return fmt.Errorf("unexpected id value for edge-in")
				}
				outValue := int(eout.Int64)
				inValue := int(ein.Int64)
				node, ok := ids[outValue]
				if !ok {
					return fmt.Errorf("unexpected node id in edges: %v", outValue)
				}
				if _, ok := edges[inValue]; !ok {
					edgeids = append(edgeids, inValue)
				}
				edges[inValue] = append(edges[inValue], node)
				return nil
			},
		}
		if err := sqlgraph.QueryEdges(ctx, fq.driver, _spec); err != nil {
			return nil, fmt.Errorf(`query edges "has_person": %w`, err)
		}
		query.Where(person.IDIn(edgeids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := edges[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected "has_person" node returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.HasPerson = append(nodes[i].Edges.HasPerson, n)
			}
		}
	}

	if query := fq.withHasPlanet; query != nil {
		fks := make([]driver.Value, 0, len(nodes))
		ids := make(map[int]*Film, len(nodes))
		for _, node := range nodes {
			ids[node.ID] = node
			fks = append(fks, node.ID)
			node.Edges.HasPlanet = []*Planet{}
		}
		var (
			edgeids []int
			edges   = make(map[int][]*Film)
		)
		_spec := &sqlgraph.EdgeQuerySpec{
			Edge: &sqlgraph.EdgeSpec{
				Inverse: false,
				Table:   film.HasPlanetTable,
				Columns: film.HasPlanetPrimaryKey,
			},
			Predicate: func(s *sql.Selector) {
				s.Where(sql.InValues(film.HasPlanetPrimaryKey[0], fks...))
			},
			ScanValues: func() [2]interface{} {
				return [2]interface{}{new(sql.NullInt64), new(sql.NullInt64)}
			},
			Assign: func(out, in interface{}) error {
				eout, ok := out.(*sql.NullInt64)
				if !ok || eout == nil {
					return fmt.Errorf("unexpected id value for edge-out")
				}
				ein, ok := in.(*sql.NullInt64)
				if !ok || ein == nil {
					return fmt.Errorf("unexpected id value for edge-in")
				}
				outValue := int(eout.Int64)
				inValue := int(ein.Int64)
				node, ok := ids[outValue]
				if !ok {
					return fmt.Errorf("unexpected node id in edges: %v", outValue)
				}
				if _, ok := edges[inValue]; !ok {
					edgeids = append(edgeids, inValue)
				}
				edges[inValue] = append(edges[inValue], node)
				return nil
			},
		}
		if err := sqlgraph.QueryEdges(ctx, fq.driver, _spec); err != nil {
			return nil, fmt.Errorf(`query edges "has_planet": %w`, err)
		}
		query.Where(planet.IDIn(edgeids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := edges[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected "has_planet" node returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.HasPlanet = append(nodes[i].Edges.HasPlanet, n)
			}
		}
	}

	if query := fq.withHasStarship; query != nil {
		fks := make([]driver.Value, 0, len(nodes))
		ids := make(map[int]*Film, len(nodes))
		for _, node := range nodes {
			ids[node.ID] = node
			fks = append(fks, node.ID)
			node.Edges.HasStarship = []*Starship{}
		}
		var (
			edgeids []int
			edges   = make(map[int][]*Film)
		)
		_spec := &sqlgraph.EdgeQuerySpec{
			Edge: &sqlgraph.EdgeSpec{
				Inverse: false,
				Table:   film.HasStarshipTable,
				Columns: film.HasStarshipPrimaryKey,
			},
			Predicate: func(s *sql.Selector) {
				s.Where(sql.InValues(film.HasStarshipPrimaryKey[0], fks...))
			},
			ScanValues: func() [2]interface{} {
				return [2]interface{}{new(sql.NullInt64), new(sql.NullInt64)}
			},
			Assign: func(out, in interface{}) error {
				eout, ok := out.(*sql.NullInt64)
				if !ok || eout == nil {
					return fmt.Errorf("unexpected id value for edge-out")
				}
				ein, ok := in.(*sql.NullInt64)
				if !ok || ein == nil {
					return fmt.Errorf("unexpected id value for edge-in")
				}
				outValue := int(eout.Int64)
				inValue := int(ein.Int64)
				node, ok := ids[outValue]
				if !ok {
					return fmt.Errorf("unexpected node id in edges: %v", outValue)
				}
				if _, ok := edges[inValue]; !ok {
					edgeids = append(edgeids, inValue)
				}
				edges[inValue] = append(edges[inValue], node)
				return nil
			},
		}
		if err := sqlgraph.QueryEdges(ctx, fq.driver, _spec); err != nil {
			return nil, fmt.Errorf(`query edges "has_starship": %w`, err)
		}
		query.Where(starship.IDIn(edgeids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := edges[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected "has_starship" node returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.HasStarship = append(nodes[i].Edges.HasStarship, n)
			}
		}
	}

	if query := fq.withHasVehicle; query != nil {
		fks := make([]driver.Value, 0, len(nodes))
		ids := make(map[int]*Film, len(nodes))
		for _, node := range nodes {
			ids[node.ID] = node
			fks = append(fks, node.ID)
			node.Edges.HasVehicle = []*Vehicle{}
		}
		var (
			edgeids []int
			edges   = make(map[int][]*Film)
		)
		_spec := &sqlgraph.EdgeQuerySpec{
			Edge: &sqlgraph.EdgeSpec{
				Inverse: false,
				Table:   film.HasVehicleTable,
				Columns: film.HasVehiclePrimaryKey,
			},
			Predicate: func(s *sql.Selector) {
				s.Where(sql.InValues(film.HasVehiclePrimaryKey[0], fks...))
			},
			ScanValues: func() [2]interface{} {
				return [2]interface{}{new(sql.NullInt64), new(sql.NullInt64)}
			},
			Assign: func(out, in interface{}) error {
				eout, ok := out.(*sql.NullInt64)
				if !ok || eout == nil {
					return fmt.Errorf("unexpected id value for edge-out")
				}
				ein, ok := in.(*sql.NullInt64)
				if !ok || ein == nil {
					return fmt.Errorf("unexpected id value for edge-in")
				}
				outValue := int(eout.Int64)
				inValue := int(ein.Int64)
				node, ok := ids[outValue]
				if !ok {
					return fmt.Errorf("unexpected node id in edges: %v", outValue)
				}
				if _, ok := edges[inValue]; !ok {
					edgeids = append(edgeids, inValue)
				}
				edges[inValue] = append(edges[inValue], node)
				return nil
			},
		}
		if err := sqlgraph.QueryEdges(ctx, fq.driver, _spec); err != nil {
			return nil, fmt.Errorf(`query edges "has_vehicle": %w`, err)
		}
		query.Where(vehicle.IDIn(edgeids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := edges[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected "has_vehicle" node returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.HasVehicle = append(nodes[i].Edges.HasVehicle, n)
			}
		}
	}

	if query := fq.withHasSpecies; query != nil {
		fks := make([]driver.Value, 0, len(nodes))
		ids := make(map[int]*Film, len(nodes))
		for _, node := range nodes {
			ids[node.ID] = node
			fks = append(fks, node.ID)
			node.Edges.HasSpecies = []*Species{}
		}
		var (
			edgeids []int
			edges   = make(map[int][]*Film)
		)
		_spec := &sqlgraph.EdgeQuerySpec{
			Edge: &sqlgraph.EdgeSpec{
				Inverse: false,
				Table:   film.HasSpeciesTable,
				Columns: film.HasSpeciesPrimaryKey,
			},
			Predicate: func(s *sql.Selector) {
				s.Where(sql.InValues(film.HasSpeciesPrimaryKey[0], fks...))
			},
			ScanValues: func() [2]interface{} {
				return [2]interface{}{new(sql.NullInt64), new(sql.NullInt64)}
			},
			Assign: func(out, in interface{}) error {
				eout, ok := out.(*sql.NullInt64)
				if !ok || eout == nil {
					return fmt.Errorf("unexpected id value for edge-out")
				}
				ein, ok := in.(*sql.NullInt64)
				if !ok || ein == nil {
					return fmt.Errorf("unexpected id value for edge-in")
				}
				outValue := int(eout.Int64)
				inValue := int(ein.Int64)
				node, ok := ids[outValue]
				if !ok {
					return fmt.Errorf("unexpected node id in edges: %v", outValue)
				}
				if _, ok := edges[inValue]; !ok {
					edgeids = append(edgeids, inValue)
				}
				edges[inValue] = append(edges[inValue], node)
				return nil
			},
		}
		if err := sqlgraph.QueryEdges(ctx, fq.driver, _spec); err != nil {
			return nil, fmt.Errorf(`query edges "has_species": %w`, err)
		}
		query.Where(species.IDIn(edgeids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := edges[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected "has_species" node returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.HasSpecies = append(nodes[i].Edges.HasSpecies, n)
			}
		}
	}

	return nodes, nil
}

func (fq *FilmQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := fq.querySpec()
	_spec.Node.Columns = fq.fields
	if len(fq.fields) > 0 {
		_spec.Unique = fq.unique != nil && *fq.unique
	}
	return sqlgraph.CountNodes(ctx, fq.driver, _spec)
}

func (fq *FilmQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := fq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (fq *FilmQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   film.Table,
			Columns: film.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: film.FieldID,
			},
		},
		From:   fq.sql,
		Unique: true,
	}
	if unique := fq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := fq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, film.FieldID)
		for i := range fields {
			if fields[i] != film.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := fq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := fq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := fq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := fq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (fq *FilmQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(fq.driver.Dialect())
	t1 := builder.Table(film.Table)
	columns := fq.fields
	if len(columns) == 0 {
		columns = film.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if fq.sql != nil {
		selector = fq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if fq.unique != nil && *fq.unique {
		selector.Distinct()
	}
	for _, p := range fq.predicates {
		p(selector)
	}
	for _, p := range fq.order {
		p(selector)
	}
	if offset := fq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := fq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// FilmGroupBy is the group-by builder for Film entities.
type FilmGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (fgb *FilmGroupBy) Aggregate(fns ...AggregateFunc) *FilmGroupBy {
	fgb.fns = append(fgb.fns, fns...)
	return fgb
}

// Scan applies the group-by query and scans the result into the given value.
func (fgb *FilmGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := fgb.path(ctx)
	if err != nil {
		return err
	}
	fgb.sql = query
	return fgb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (fgb *FilmGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := fgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by.
// It is only allowed when executing a group-by query with one field.
func (fgb *FilmGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(fgb.fields) > 1 {
		return nil, errors.New("ent: FilmGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := fgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (fgb *FilmGroupBy) StringsX(ctx context.Context) []string {
	v, err := fgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (fgb *FilmGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = fgb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{film.Label}
	default:
		err = fmt.Errorf("ent: FilmGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (fgb *FilmGroupBy) StringX(ctx context.Context) string {
	v, err := fgb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by.
// It is only allowed when executing a group-by query with one field.
func (fgb *FilmGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(fgb.fields) > 1 {
		return nil, errors.New("ent: FilmGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := fgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (fgb *FilmGroupBy) IntsX(ctx context.Context) []int {
	v, err := fgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (fgb *FilmGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = fgb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{film.Label}
	default:
		err = fmt.Errorf("ent: FilmGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (fgb *FilmGroupBy) IntX(ctx context.Context) int {
	v, err := fgb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by.
// It is only allowed when executing a group-by query with one field.
func (fgb *FilmGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(fgb.fields) > 1 {
		return nil, errors.New("ent: FilmGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := fgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (fgb *FilmGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := fgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (fgb *FilmGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = fgb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{film.Label}
	default:
		err = fmt.Errorf("ent: FilmGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (fgb *FilmGroupBy) Float64X(ctx context.Context) float64 {
	v, err := fgb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by.
// It is only allowed when executing a group-by query with one field.
func (fgb *FilmGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(fgb.fields) > 1 {
		return nil, errors.New("ent: FilmGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := fgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (fgb *FilmGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := fgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (fgb *FilmGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = fgb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{film.Label}
	default:
		err = fmt.Errorf("ent: FilmGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (fgb *FilmGroupBy) BoolX(ctx context.Context) bool {
	v, err := fgb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (fgb *FilmGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range fgb.fields {
		if !film.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := fgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := fgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (fgb *FilmGroupBy) sqlQuery() *sql.Selector {
	selector := fgb.sql.Select()
	aggregation := make([]string, 0, len(fgb.fns))
	for _, fn := range fgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(fgb.fields)+len(fgb.fns))
		for _, f := range fgb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(fgb.fields...)...)
}

// FilmSelect is the builder for selecting fields of Film entities.
type FilmSelect struct {
	*FilmQuery
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (fs *FilmSelect) Scan(ctx context.Context, v interface{}) error {
	if err := fs.prepareQuery(ctx); err != nil {
		return err
	}
	fs.sql = fs.FilmQuery.sqlQuery(ctx)
	return fs.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (fs *FilmSelect) ScanX(ctx context.Context, v interface{}) {
	if err := fs.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from a selector. It is only allowed when selecting one field.
func (fs *FilmSelect) Strings(ctx context.Context) ([]string, error) {
	if len(fs.fields) > 1 {
		return nil, errors.New("ent: FilmSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := fs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (fs *FilmSelect) StringsX(ctx context.Context) []string {
	v, err := fs.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a selector. It is only allowed when selecting one field.
func (fs *FilmSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = fs.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{film.Label}
	default:
		err = fmt.Errorf("ent: FilmSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (fs *FilmSelect) StringX(ctx context.Context) string {
	v, err := fs.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from a selector. It is only allowed when selecting one field.
func (fs *FilmSelect) Ints(ctx context.Context) ([]int, error) {
	if len(fs.fields) > 1 {
		return nil, errors.New("ent: FilmSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := fs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (fs *FilmSelect) IntsX(ctx context.Context) []int {
	v, err := fs.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a selector. It is only allowed when selecting one field.
func (fs *FilmSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = fs.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{film.Label}
	default:
		err = fmt.Errorf("ent: FilmSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (fs *FilmSelect) IntX(ctx context.Context) int {
	v, err := fs.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from a selector. It is only allowed when selecting one field.
func (fs *FilmSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(fs.fields) > 1 {
		return nil, errors.New("ent: FilmSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := fs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (fs *FilmSelect) Float64sX(ctx context.Context) []float64 {
	v, err := fs.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a selector. It is only allowed when selecting one field.
func (fs *FilmSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = fs.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{film.Label}
	default:
		err = fmt.Errorf("ent: FilmSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (fs *FilmSelect) Float64X(ctx context.Context) float64 {
	v, err := fs.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from a selector. It is only allowed when selecting one field.
func (fs *FilmSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(fs.fields) > 1 {
		return nil, errors.New("ent: FilmSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := fs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (fs *FilmSelect) BoolsX(ctx context.Context) []bool {
	v, err := fs.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a selector. It is only allowed when selecting one field.
func (fs *FilmSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = fs.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{film.Label}
	default:
		err = fmt.Errorf("ent: FilmSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (fs *FilmSelect) BoolX(ctx context.Context) bool {
	v, err := fs.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (fs *FilmSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := fs.sql.Query()
	if err := fs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
