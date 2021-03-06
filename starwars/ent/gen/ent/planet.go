// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/inigolabs/entbrowse/starwars/ent/gen/ent/planet"
)

// Planet is the model entity for the Planet schema.
type Planet struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Climate holds the value of the "climate" field.
	Climate string `json:"climate,omitempty"`
	// Diameter holds the value of the "diameter" field.
	Diameter int `json:"diameter,omitempty"`
	// Gravity holds the value of the "gravity" field.
	Gravity string `json:"gravity,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// OrbitalPeriod holds the value of the "orbital_period" field.
	OrbitalPeriod string `json:"orbital_period,omitempty"`
	// Population holds the value of the "population" field.
	Population int `json:"population,omitempty"`
	// RotationPeriod holds the value of the "rotation_period" field.
	RotationPeriod string `json:"rotation_period,omitempty"`
	// SurfaceWater holds the value of the "surface_water" field.
	SurfaceWater string `json:"surface_water,omitempty"`
	// Terrain holds the value of the "terrain" field.
	Terrain string `json:"terrain,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the PlanetQuery when eager-loading is set.
	Edges PlanetEdges `json:"edges"`
}

// PlanetEdges holds the relations/edges for other nodes in the graph.
type PlanetEdges struct {
	// HomeTo holds the value of the home_to edge.
	HomeTo []*Person `json:"home_to,omitempty"`
	// AppearedIn holds the value of the appeared_in edge.
	AppearedIn []*Film `json:"appeared_in,omitempty"`
	// OriginOf holds the value of the origin_of edge.
	OriginOf []*Species `json:"origin_of,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [3]bool
}

// HomeToOrErr returns the HomeTo value or an error if the edge
// was not loaded in eager-loading.
func (e PlanetEdges) HomeToOrErr() ([]*Person, error) {
	if e.loadedTypes[0] {
		return e.HomeTo, nil
	}
	return nil, &NotLoadedError{edge: "home_to"}
}

// AppearedInOrErr returns the AppearedIn value or an error if the edge
// was not loaded in eager-loading.
func (e PlanetEdges) AppearedInOrErr() ([]*Film, error) {
	if e.loadedTypes[1] {
		return e.AppearedIn, nil
	}
	return nil, &NotLoadedError{edge: "appeared_in"}
}

// OriginOfOrErr returns the OriginOf value or an error if the edge
// was not loaded in eager-loading.
func (e PlanetEdges) OriginOfOrErr() ([]*Species, error) {
	if e.loadedTypes[2] {
		return e.OriginOf, nil
	}
	return nil, &NotLoadedError{edge: "origin_of"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Planet) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case planet.FieldID, planet.FieldDiameter, planet.FieldPopulation:
			values[i] = new(sql.NullInt64)
		case planet.FieldClimate, planet.FieldGravity, planet.FieldName, planet.FieldOrbitalPeriod, planet.FieldRotationPeriod, planet.FieldSurfaceWater, planet.FieldTerrain:
			values[i] = new(sql.NullString)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Planet", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Planet fields.
func (pl *Planet) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case planet.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			pl.ID = int(value.Int64)
		case planet.FieldClimate:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field climate", values[i])
			} else if value.Valid {
				pl.Climate = value.String
			}
		case planet.FieldDiameter:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field diameter", values[i])
			} else if value.Valid {
				pl.Diameter = int(value.Int64)
			}
		case planet.FieldGravity:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field gravity", values[i])
			} else if value.Valid {
				pl.Gravity = value.String
			}
		case planet.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				pl.Name = value.String
			}
		case planet.FieldOrbitalPeriod:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field orbital_period", values[i])
			} else if value.Valid {
				pl.OrbitalPeriod = value.String
			}
		case planet.FieldPopulation:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field population", values[i])
			} else if value.Valid {
				pl.Population = int(value.Int64)
			}
		case planet.FieldRotationPeriod:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field rotation_period", values[i])
			} else if value.Valid {
				pl.RotationPeriod = value.String
			}
		case planet.FieldSurfaceWater:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field surface_water", values[i])
			} else if value.Valid {
				pl.SurfaceWater = value.String
			}
		case planet.FieldTerrain:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field terrain", values[i])
			} else if value.Valid {
				pl.Terrain = value.String
			}
		}
	}
	return nil
}

// QueryHomeTo queries the "home_to" edge of the Planet entity.
func (pl *Planet) QueryHomeTo() *PersonQuery {
	return (&PlanetClient{config: pl.config}).QueryHomeTo(pl)
}

// QueryAppearedIn queries the "appeared_in" edge of the Planet entity.
func (pl *Planet) QueryAppearedIn() *FilmQuery {
	return (&PlanetClient{config: pl.config}).QueryAppearedIn(pl)
}

// QueryOriginOf queries the "origin_of" edge of the Planet entity.
func (pl *Planet) QueryOriginOf() *SpeciesQuery {
	return (&PlanetClient{config: pl.config}).QueryOriginOf(pl)
}

// Update returns a builder for updating this Planet.
// Note that you need to call Planet.Unwrap() before calling this method if this Planet
// was returned from a transaction, and the transaction was committed or rolled back.
func (pl *Planet) Update() *PlanetUpdateOne {
	return (&PlanetClient{config: pl.config}).UpdateOne(pl)
}

// Unwrap unwraps the Planet entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (pl *Planet) Unwrap() *Planet {
	tx, ok := pl.config.driver.(*txDriver)
	if !ok {
		panic("ent: Planet is not a transactional entity")
	}
	pl.config.driver = tx.drv
	return pl
}

// String implements the fmt.Stringer.
func (pl *Planet) String() string {
	var builder strings.Builder
	builder.WriteString("Planet(")
	builder.WriteString(fmt.Sprintf("id=%v", pl.ID))
	builder.WriteString(", climate=")
	builder.WriteString(pl.Climate)
	builder.WriteString(", diameter=")
	builder.WriteString(fmt.Sprintf("%v", pl.Diameter))
	builder.WriteString(", gravity=")
	builder.WriteString(pl.Gravity)
	builder.WriteString(", name=")
	builder.WriteString(pl.Name)
	builder.WriteString(", orbital_period=")
	builder.WriteString(pl.OrbitalPeriod)
	builder.WriteString(", population=")
	builder.WriteString(fmt.Sprintf("%v", pl.Population))
	builder.WriteString(", rotation_period=")
	builder.WriteString(pl.RotationPeriod)
	builder.WriteString(", surface_water=")
	builder.WriteString(pl.SurfaceWater)
	builder.WriteString(", terrain=")
	builder.WriteString(pl.Terrain)
	builder.WriteByte(')')
	return builder.String()
}

// Planets is a parsable slice of Planet.
type Planets []*Planet

func (pl Planets) config(cfg config) {
	for _i := range pl {
		pl[_i].config = cfg
	}
}
