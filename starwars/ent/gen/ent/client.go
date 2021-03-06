// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"log"

	"github.com/inigolabs/entbrowse/starwars/ent/gen/ent/migrate"

	"github.com/inigolabs/entbrowse/starwars/ent/gen/ent/film"
	"github.com/inigolabs/entbrowse/starwars/ent/gen/ent/person"
	"github.com/inigolabs/entbrowse/starwars/ent/gen/ent/planet"
	"github.com/inigolabs/entbrowse/starwars/ent/gen/ent/species"
	"github.com/inigolabs/entbrowse/starwars/ent/gen/ent/starship"
	"github.com/inigolabs/entbrowse/starwars/ent/gen/ent/vehicle"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// Film is the client for interacting with the Film builders.
	Film *FilmClient
	// Person is the client for interacting with the Person builders.
	Person *PersonClient
	// Planet is the client for interacting with the Planet builders.
	Planet *PlanetClient
	// Species is the client for interacting with the Species builders.
	Species *SpeciesClient
	// Starship is the client for interacting with the Starship builders.
	Starship *StarshipClient
	// Vehicle is the client for interacting with the Vehicle builders.
	Vehicle *VehicleClient
}

// NewClient creates a new client configured with the given options.
func NewClient(opts ...Option) *Client {
	cfg := config{log: log.Println, hooks: &hooks{}}
	cfg.options(opts...)
	client := &Client{config: cfg}
	client.init()
	return client
}

func (c *Client) init() {
	c.Schema = migrate.NewSchema(c.driver)
	c.Film = NewFilmClient(c.config)
	c.Person = NewPersonClient(c.config)
	c.Planet = NewPlanetClient(c.config)
	c.Species = NewSpeciesClient(c.config)
	c.Starship = NewStarshipClient(c.config)
	c.Vehicle = NewVehicleClient(c.config)
}

// Open opens a database/sql.DB specified by the driver name and
// the data source name, and returns a new client attached to it.
// Optional parameters can be added for configuring the client.
func Open(driverName, dataSourceName string, options ...Option) (*Client, error) {
	switch driverName {
	case dialect.MySQL, dialect.Postgres, dialect.SQLite:
		drv, err := sql.Open(driverName, dataSourceName)
		if err != nil {
			return nil, err
		}
		return NewClient(append(options, Driver(drv))...), nil
	default:
		return nil, fmt.Errorf("unsupported driver: %q", driverName)
	}
}

// Tx returns a new transactional client. The provided context
// is used until the transaction is committed or rolled back.
func (c *Client) Tx(ctx context.Context) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, fmt.Errorf("ent: cannot start a transaction within a transaction")
	}
	tx, err := newTx(ctx, c.driver)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = tx
	return &Tx{
		ctx:      ctx,
		config:   cfg,
		Film:     NewFilmClient(cfg),
		Person:   NewPersonClient(cfg),
		Planet:   NewPlanetClient(cfg),
		Species:  NewSpeciesClient(cfg),
		Starship: NewStarshipClient(cfg),
		Vehicle:  NewVehicleClient(cfg),
	}, nil
}

// BeginTx returns a transactional client with specified options.
func (c *Client) BeginTx(ctx context.Context, opts *sql.TxOptions) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, fmt.Errorf("ent: cannot start a transaction within a transaction")
	}
	tx, err := c.driver.(interface {
		BeginTx(context.Context, *sql.TxOptions) (dialect.Tx, error)
	}).BeginTx(ctx, opts)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = &txDriver{tx: tx, drv: c.driver}
	return &Tx{
		ctx:      ctx,
		config:   cfg,
		Film:     NewFilmClient(cfg),
		Person:   NewPersonClient(cfg),
		Planet:   NewPlanetClient(cfg),
		Species:  NewSpeciesClient(cfg),
		Starship: NewStarshipClient(cfg),
		Vehicle:  NewVehicleClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		Film.
//		Query().
//		Count(ctx)
//
func (c *Client) Debug() *Client {
	if c.debug {
		return c
	}
	cfg := c.config
	cfg.driver = dialect.Debug(c.driver, c.log)
	client := &Client{config: cfg}
	client.init()
	return client
}

// Close closes the database connection and prevents new queries from starting.
func (c *Client) Close() error {
	return c.driver.Close()
}

// Use adds the mutation hooks to all the entity clients.
// In order to add hooks to a specific client, call: `client.Node.Use(...)`.
func (c *Client) Use(hooks ...Hook) {
	c.Film.Use(hooks...)
	c.Person.Use(hooks...)
	c.Planet.Use(hooks...)
	c.Species.Use(hooks...)
	c.Starship.Use(hooks...)
	c.Vehicle.Use(hooks...)
}

// FilmClient is a client for the Film schema.
type FilmClient struct {
	config
}

// NewFilmClient returns a client for the Film from the given config.
func NewFilmClient(c config) *FilmClient {
	return &FilmClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `film.Hooks(f(g(h())))`.
func (c *FilmClient) Use(hooks ...Hook) {
	c.hooks.Film = append(c.hooks.Film, hooks...)
}

// Create returns a create builder for Film.
func (c *FilmClient) Create() *FilmCreate {
	mutation := newFilmMutation(c.config, OpCreate)
	return &FilmCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Film entities.
func (c *FilmClient) CreateBulk(builders ...*FilmCreate) *FilmCreateBulk {
	return &FilmCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Film.
func (c *FilmClient) Update() *FilmUpdate {
	mutation := newFilmMutation(c.config, OpUpdate)
	return &FilmUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *FilmClient) UpdateOne(f *Film) *FilmUpdateOne {
	mutation := newFilmMutation(c.config, OpUpdateOne, withFilm(f))
	return &FilmUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *FilmClient) UpdateOneID(id int) *FilmUpdateOne {
	mutation := newFilmMutation(c.config, OpUpdateOne, withFilmID(id))
	return &FilmUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Film.
func (c *FilmClient) Delete() *FilmDelete {
	mutation := newFilmMutation(c.config, OpDelete)
	return &FilmDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *FilmClient) DeleteOne(f *Film) *FilmDeleteOne {
	return c.DeleteOneID(f.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *FilmClient) DeleteOneID(id int) *FilmDeleteOne {
	builder := c.Delete().Where(film.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &FilmDeleteOne{builder}
}

// Query returns a query builder for Film.
func (c *FilmClient) Query() *FilmQuery {
	return &FilmQuery{
		config: c.config,
	}
}

// Get returns a Film entity by its id.
func (c *FilmClient) Get(ctx context.Context, id int) (*Film, error) {
	return c.Query().Where(film.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *FilmClient) GetX(ctx context.Context, id int) *Film {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryHasPerson queries the has_person edge of a Film.
func (c *FilmClient) QueryHasPerson(f *Film) *PersonQuery {
	query := &PersonQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := f.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(film.Table, film.FieldID, id),
			sqlgraph.To(person.Table, person.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, film.HasPersonTable, film.HasPersonPrimaryKey...),
		)
		fromV = sqlgraph.Neighbors(f.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryHasPlanet queries the has_planet edge of a Film.
func (c *FilmClient) QueryHasPlanet(f *Film) *PlanetQuery {
	query := &PlanetQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := f.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(film.Table, film.FieldID, id),
			sqlgraph.To(planet.Table, planet.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, film.HasPlanetTable, film.HasPlanetPrimaryKey...),
		)
		fromV = sqlgraph.Neighbors(f.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryHasStarship queries the has_starship edge of a Film.
func (c *FilmClient) QueryHasStarship(f *Film) *StarshipQuery {
	query := &StarshipQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := f.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(film.Table, film.FieldID, id),
			sqlgraph.To(starship.Table, starship.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, film.HasStarshipTable, film.HasStarshipPrimaryKey...),
		)
		fromV = sqlgraph.Neighbors(f.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryHasVehicle queries the has_vehicle edge of a Film.
func (c *FilmClient) QueryHasVehicle(f *Film) *VehicleQuery {
	query := &VehicleQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := f.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(film.Table, film.FieldID, id),
			sqlgraph.To(vehicle.Table, vehicle.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, film.HasVehicleTable, film.HasVehiclePrimaryKey...),
		)
		fromV = sqlgraph.Neighbors(f.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryHasSpecies queries the has_species edge of a Film.
func (c *FilmClient) QueryHasSpecies(f *Film) *SpeciesQuery {
	query := &SpeciesQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := f.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(film.Table, film.FieldID, id),
			sqlgraph.To(species.Table, species.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, film.HasSpeciesTable, film.HasSpeciesPrimaryKey...),
		)
		fromV = sqlgraph.Neighbors(f.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *FilmClient) Hooks() []Hook {
	return c.hooks.Film
}

// PersonClient is a client for the Person schema.
type PersonClient struct {
	config
}

// NewPersonClient returns a client for the Person from the given config.
func NewPersonClient(c config) *PersonClient {
	return &PersonClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `person.Hooks(f(g(h())))`.
func (c *PersonClient) Use(hooks ...Hook) {
	c.hooks.Person = append(c.hooks.Person, hooks...)
}

// Create returns a create builder for Person.
func (c *PersonClient) Create() *PersonCreate {
	mutation := newPersonMutation(c.config, OpCreate)
	return &PersonCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Person entities.
func (c *PersonClient) CreateBulk(builders ...*PersonCreate) *PersonCreateBulk {
	return &PersonCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Person.
func (c *PersonClient) Update() *PersonUpdate {
	mutation := newPersonMutation(c.config, OpUpdate)
	return &PersonUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *PersonClient) UpdateOne(pe *Person) *PersonUpdateOne {
	mutation := newPersonMutation(c.config, OpUpdateOne, withPerson(pe))
	return &PersonUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *PersonClient) UpdateOneID(id int) *PersonUpdateOne {
	mutation := newPersonMutation(c.config, OpUpdateOne, withPersonID(id))
	return &PersonUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Person.
func (c *PersonClient) Delete() *PersonDelete {
	mutation := newPersonMutation(c.config, OpDelete)
	return &PersonDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *PersonClient) DeleteOne(pe *Person) *PersonDeleteOne {
	return c.DeleteOneID(pe.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *PersonClient) DeleteOneID(id int) *PersonDeleteOne {
	builder := c.Delete().Where(person.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &PersonDeleteOne{builder}
}

// Query returns a query builder for Person.
func (c *PersonClient) Query() *PersonQuery {
	return &PersonQuery{
		config: c.config,
	}
}

// Get returns a Person entity by its id.
func (c *PersonClient) Get(ctx context.Context, id int) (*Person, error) {
	return c.Query().Where(person.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *PersonClient) GetX(ctx context.Context, id int) *Person {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryPilotedStarship queries the piloted_starship edge of a Person.
func (c *PersonClient) QueryPilotedStarship(pe *Person) *StarshipQuery {
	query := &StarshipQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := pe.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(person.Table, person.FieldID, id),
			sqlgraph.To(starship.Table, starship.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, person.PilotedStarshipTable, person.PilotedStarshipPrimaryKey...),
		)
		fromV = sqlgraph.Neighbors(pe.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryPilotedVehicle queries the piloted_vehicle edge of a Person.
func (c *PersonClient) QueryPilotedVehicle(pe *Person) *VehicleQuery {
	query := &VehicleQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := pe.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(person.Table, person.FieldID, id),
			sqlgraph.To(vehicle.Table, vehicle.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, person.PilotedVehicleTable, person.PilotedVehiclePrimaryKey...),
		)
		fromV = sqlgraph.Neighbors(pe.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryIsOfType queries the is_of_type edge of a Person.
func (c *PersonClient) QueryIsOfType(pe *Person) *SpeciesQuery {
	query := &SpeciesQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := pe.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(person.Table, person.FieldID, id),
			sqlgraph.To(species.Table, species.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, person.IsOfTypeTable, person.IsOfTypePrimaryKey...),
		)
		fromV = sqlgraph.Neighbors(pe.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryAppearedIn queries the appeared_in edge of a Person.
func (c *PersonClient) QueryAppearedIn(pe *Person) *FilmQuery {
	query := &FilmQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := pe.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(person.Table, person.FieldID, id),
			sqlgraph.To(film.Table, film.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, person.AppearedInTable, person.AppearedInPrimaryKey...),
		)
		fromV = sqlgraph.Neighbors(pe.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryFromPlanet queries the from_planet edge of a Person.
func (c *PersonClient) QueryFromPlanet(pe *Person) *PlanetQuery {
	query := &PlanetQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := pe.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(person.Table, person.FieldID, id),
			sqlgraph.To(planet.Table, planet.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, person.FromPlanetTable, person.FromPlanetPrimaryKey...),
		)
		fromV = sqlgraph.Neighbors(pe.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *PersonClient) Hooks() []Hook {
	return c.hooks.Person
}

// PlanetClient is a client for the Planet schema.
type PlanetClient struct {
	config
}

// NewPlanetClient returns a client for the Planet from the given config.
func NewPlanetClient(c config) *PlanetClient {
	return &PlanetClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `planet.Hooks(f(g(h())))`.
func (c *PlanetClient) Use(hooks ...Hook) {
	c.hooks.Planet = append(c.hooks.Planet, hooks...)
}

// Create returns a create builder for Planet.
func (c *PlanetClient) Create() *PlanetCreate {
	mutation := newPlanetMutation(c.config, OpCreate)
	return &PlanetCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Planet entities.
func (c *PlanetClient) CreateBulk(builders ...*PlanetCreate) *PlanetCreateBulk {
	return &PlanetCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Planet.
func (c *PlanetClient) Update() *PlanetUpdate {
	mutation := newPlanetMutation(c.config, OpUpdate)
	return &PlanetUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *PlanetClient) UpdateOne(pl *Planet) *PlanetUpdateOne {
	mutation := newPlanetMutation(c.config, OpUpdateOne, withPlanet(pl))
	return &PlanetUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *PlanetClient) UpdateOneID(id int) *PlanetUpdateOne {
	mutation := newPlanetMutation(c.config, OpUpdateOne, withPlanetID(id))
	return &PlanetUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Planet.
func (c *PlanetClient) Delete() *PlanetDelete {
	mutation := newPlanetMutation(c.config, OpDelete)
	return &PlanetDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *PlanetClient) DeleteOne(pl *Planet) *PlanetDeleteOne {
	return c.DeleteOneID(pl.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *PlanetClient) DeleteOneID(id int) *PlanetDeleteOne {
	builder := c.Delete().Where(planet.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &PlanetDeleteOne{builder}
}

// Query returns a query builder for Planet.
func (c *PlanetClient) Query() *PlanetQuery {
	return &PlanetQuery{
		config: c.config,
	}
}

// Get returns a Planet entity by its id.
func (c *PlanetClient) Get(ctx context.Context, id int) (*Planet, error) {
	return c.Query().Where(planet.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *PlanetClient) GetX(ctx context.Context, id int) *Planet {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryHomeTo queries the home_to edge of a Planet.
func (c *PlanetClient) QueryHomeTo(pl *Planet) *PersonQuery {
	query := &PersonQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := pl.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(planet.Table, planet.FieldID, id),
			sqlgraph.To(person.Table, person.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, planet.HomeToTable, planet.HomeToPrimaryKey...),
		)
		fromV = sqlgraph.Neighbors(pl.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryAppearedIn queries the appeared_in edge of a Planet.
func (c *PlanetClient) QueryAppearedIn(pl *Planet) *FilmQuery {
	query := &FilmQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := pl.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(planet.Table, planet.FieldID, id),
			sqlgraph.To(film.Table, film.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, planet.AppearedInTable, planet.AppearedInPrimaryKey...),
		)
		fromV = sqlgraph.Neighbors(pl.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryOriginOf queries the origin_of edge of a Planet.
func (c *PlanetClient) QueryOriginOf(pl *Planet) *SpeciesQuery {
	query := &SpeciesQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := pl.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(planet.Table, planet.FieldID, id),
			sqlgraph.To(species.Table, species.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, planet.OriginOfTable, planet.OriginOfPrimaryKey...),
		)
		fromV = sqlgraph.Neighbors(pl.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *PlanetClient) Hooks() []Hook {
	return c.hooks.Planet
}

// SpeciesClient is a client for the Species schema.
type SpeciesClient struct {
	config
}

// NewSpeciesClient returns a client for the Species from the given config.
func NewSpeciesClient(c config) *SpeciesClient {
	return &SpeciesClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `species.Hooks(f(g(h())))`.
func (c *SpeciesClient) Use(hooks ...Hook) {
	c.hooks.Species = append(c.hooks.Species, hooks...)
}

// Create returns a create builder for Species.
func (c *SpeciesClient) Create() *SpeciesCreate {
	mutation := newSpeciesMutation(c.config, OpCreate)
	return &SpeciesCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Species entities.
func (c *SpeciesClient) CreateBulk(builders ...*SpeciesCreate) *SpeciesCreateBulk {
	return &SpeciesCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Species.
func (c *SpeciesClient) Update() *SpeciesUpdate {
	mutation := newSpeciesMutation(c.config, OpUpdate)
	return &SpeciesUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *SpeciesClient) UpdateOne(s *Species) *SpeciesUpdateOne {
	mutation := newSpeciesMutation(c.config, OpUpdateOne, withSpecies(s))
	return &SpeciesUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *SpeciesClient) UpdateOneID(id int) *SpeciesUpdateOne {
	mutation := newSpeciesMutation(c.config, OpUpdateOne, withSpeciesID(id))
	return &SpeciesUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Species.
func (c *SpeciesClient) Delete() *SpeciesDelete {
	mutation := newSpeciesMutation(c.config, OpDelete)
	return &SpeciesDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *SpeciesClient) DeleteOne(s *Species) *SpeciesDeleteOne {
	return c.DeleteOneID(s.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *SpeciesClient) DeleteOneID(id int) *SpeciesDeleteOne {
	builder := c.Delete().Where(species.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &SpeciesDeleteOne{builder}
}

// Query returns a query builder for Species.
func (c *SpeciesClient) Query() *SpeciesQuery {
	return &SpeciesQuery{
		config: c.config,
	}
}

// Get returns a Species entity by its id.
func (c *SpeciesClient) Get(ctx context.Context, id int) (*Species, error) {
	return c.Query().Where(species.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *SpeciesClient) GetX(ctx context.Context, id int) *Species {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryOriginatesFrom queries the originates_from edge of a Species.
func (c *SpeciesClient) QueryOriginatesFrom(s *Species) *PlanetQuery {
	query := &PlanetQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := s.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(species.Table, species.FieldID, id),
			sqlgraph.To(planet.Table, planet.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, species.OriginatesFromTable, species.OriginatesFromPrimaryKey...),
		)
		fromV = sqlgraph.Neighbors(s.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryAppearedIn queries the appeared_in edge of a Species.
func (c *SpeciesClient) QueryAppearedIn(s *Species) *FilmQuery {
	query := &FilmQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := s.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(species.Table, species.FieldID, id),
			sqlgraph.To(film.Table, film.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, species.AppearedInTable, species.AppearedInPrimaryKey...),
		)
		fromV = sqlgraph.Neighbors(s.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryIncludesPerson queries the includes_person edge of a Species.
func (c *SpeciesClient) QueryIncludesPerson(s *Species) *PersonQuery {
	query := &PersonQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := s.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(species.Table, species.FieldID, id),
			sqlgraph.To(person.Table, person.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, species.IncludesPersonTable, species.IncludesPersonPrimaryKey...),
		)
		fromV = sqlgraph.Neighbors(s.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *SpeciesClient) Hooks() []Hook {
	return c.hooks.Species
}

// StarshipClient is a client for the Starship schema.
type StarshipClient struct {
	config
}

// NewStarshipClient returns a client for the Starship from the given config.
func NewStarshipClient(c config) *StarshipClient {
	return &StarshipClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `starship.Hooks(f(g(h())))`.
func (c *StarshipClient) Use(hooks ...Hook) {
	c.hooks.Starship = append(c.hooks.Starship, hooks...)
}

// Create returns a create builder for Starship.
func (c *StarshipClient) Create() *StarshipCreate {
	mutation := newStarshipMutation(c.config, OpCreate)
	return &StarshipCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Starship entities.
func (c *StarshipClient) CreateBulk(builders ...*StarshipCreate) *StarshipCreateBulk {
	return &StarshipCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Starship.
func (c *StarshipClient) Update() *StarshipUpdate {
	mutation := newStarshipMutation(c.config, OpUpdate)
	return &StarshipUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *StarshipClient) UpdateOne(s *Starship) *StarshipUpdateOne {
	mutation := newStarshipMutation(c.config, OpUpdateOne, withStarship(s))
	return &StarshipUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *StarshipClient) UpdateOneID(id int) *StarshipUpdateOne {
	mutation := newStarshipMutation(c.config, OpUpdateOne, withStarshipID(id))
	return &StarshipUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Starship.
func (c *StarshipClient) Delete() *StarshipDelete {
	mutation := newStarshipMutation(c.config, OpDelete)
	return &StarshipDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *StarshipClient) DeleteOne(s *Starship) *StarshipDeleteOne {
	return c.DeleteOneID(s.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *StarshipClient) DeleteOneID(id int) *StarshipDeleteOne {
	builder := c.Delete().Where(starship.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &StarshipDeleteOne{builder}
}

// Query returns a query builder for Starship.
func (c *StarshipClient) Query() *StarshipQuery {
	return &StarshipQuery{
		config: c.config,
	}
}

// Get returns a Starship entity by its id.
func (c *StarshipClient) Get(ctx context.Context, id int) (*Starship, error) {
	return c.Query().Where(starship.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *StarshipClient) GetX(ctx context.Context, id int) *Starship {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryAppearedIn queries the appeared_in edge of a Starship.
func (c *StarshipClient) QueryAppearedIn(s *Starship) *FilmQuery {
	query := &FilmQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := s.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(starship.Table, starship.FieldID, id),
			sqlgraph.To(film.Table, film.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, starship.AppearedInTable, starship.AppearedInPrimaryKey...),
		)
		fromV = sqlgraph.Neighbors(s.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryPilotedBy queries the piloted_by edge of a Starship.
func (c *StarshipClient) QueryPilotedBy(s *Starship) *PersonQuery {
	query := &PersonQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := s.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(starship.Table, starship.FieldID, id),
			sqlgraph.To(person.Table, person.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, starship.PilotedByTable, starship.PilotedByPrimaryKey...),
		)
		fromV = sqlgraph.Neighbors(s.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *StarshipClient) Hooks() []Hook {
	return c.hooks.Starship
}

// VehicleClient is a client for the Vehicle schema.
type VehicleClient struct {
	config
}

// NewVehicleClient returns a client for the Vehicle from the given config.
func NewVehicleClient(c config) *VehicleClient {
	return &VehicleClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `vehicle.Hooks(f(g(h())))`.
func (c *VehicleClient) Use(hooks ...Hook) {
	c.hooks.Vehicle = append(c.hooks.Vehicle, hooks...)
}

// Create returns a create builder for Vehicle.
func (c *VehicleClient) Create() *VehicleCreate {
	mutation := newVehicleMutation(c.config, OpCreate)
	return &VehicleCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Vehicle entities.
func (c *VehicleClient) CreateBulk(builders ...*VehicleCreate) *VehicleCreateBulk {
	return &VehicleCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Vehicle.
func (c *VehicleClient) Update() *VehicleUpdate {
	mutation := newVehicleMutation(c.config, OpUpdate)
	return &VehicleUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *VehicleClient) UpdateOne(v *Vehicle) *VehicleUpdateOne {
	mutation := newVehicleMutation(c.config, OpUpdateOne, withVehicle(v))
	return &VehicleUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *VehicleClient) UpdateOneID(id int) *VehicleUpdateOne {
	mutation := newVehicleMutation(c.config, OpUpdateOne, withVehicleID(id))
	return &VehicleUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Vehicle.
func (c *VehicleClient) Delete() *VehicleDelete {
	mutation := newVehicleMutation(c.config, OpDelete)
	return &VehicleDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *VehicleClient) DeleteOne(v *Vehicle) *VehicleDeleteOne {
	return c.DeleteOneID(v.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *VehicleClient) DeleteOneID(id int) *VehicleDeleteOne {
	builder := c.Delete().Where(vehicle.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &VehicleDeleteOne{builder}
}

// Query returns a query builder for Vehicle.
func (c *VehicleClient) Query() *VehicleQuery {
	return &VehicleQuery{
		config: c.config,
	}
}

// Get returns a Vehicle entity by its id.
func (c *VehicleClient) Get(ctx context.Context, id int) (*Vehicle, error) {
	return c.Query().Where(vehicle.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *VehicleClient) GetX(ctx context.Context, id int) *Vehicle {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryAppearedIn queries the appeared_in edge of a Vehicle.
func (c *VehicleClient) QueryAppearedIn(v *Vehicle) *FilmQuery {
	query := &FilmQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := v.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(vehicle.Table, vehicle.FieldID, id),
			sqlgraph.To(film.Table, film.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, vehicle.AppearedInTable, vehicle.AppearedInPrimaryKey...),
		)
		fromV = sqlgraph.Neighbors(v.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryPilotedBy queries the piloted_by edge of a Vehicle.
func (c *VehicleClient) QueryPilotedBy(v *Vehicle) *PersonQuery {
	query := &PersonQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := v.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(vehicle.Table, vehicle.FieldID, id),
			sqlgraph.To(person.Table, person.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, vehicle.PilotedByTable, vehicle.PilotedByPrimaryKey...),
		)
		fromV = sqlgraph.Neighbors(v.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *VehicleClient) Hooks() []Hook {
	return c.hooks.Vehicle
}
