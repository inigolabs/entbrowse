// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/inigolabs/entbrowse/starwars/ent/gen/ent/starship"
)

// Starship is the model entity for the Starship schema.
type Starship struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// CargoCapacity holds the value of the "cargo_capacity" field.
	CargoCapacity int `json:"cargo_capacity,omitempty"`
	// Class holds the value of the "class" field.
	Class string `json:"class,omitempty"`
	// Consumables holds the value of the "consumables" field.
	Consumables string `json:"consumables,omitempty"`
	// CostInCredits holds the value of the "cost_in_credits" field.
	CostInCredits int `json:"cost_in_credits,omitempty"`
	// Crew holds the value of the "crew" field.
	Crew string `json:"crew,omitempty"`
	// HyperdriveRating holds the value of the "hyperdrive_rating" field.
	HyperdriveRating string `json:"hyperdrive_rating,omitempty"`
	// Length holds the value of the "length" field.
	Length float64 `json:"length,omitempty"`
	// Manufacturer holds the value of the "manufacturer" field.
	Manufacturer string `json:"manufacturer,omitempty"`
	// MaxAtmospheringSpeed holds the value of the "max_atmosphering_speed" field.
	MaxAtmospheringSpeed string `json:"max_atmosphering_speed,omitempty"`
	// MaximumMegalights holds the value of the "maximum_megalights" field.
	MaximumMegalights string `json:"maximum_megalights,omitempty"`
	// Model holds the value of the "model" field.
	Model string `json:"model,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// PassengerCapacity holds the value of the "passenger_capacity" field.
	PassengerCapacity int `json:"passenger_capacity,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the StarshipQuery when eager-loading is set.
	Edges StarshipEdges `json:"edges"`
}

// StarshipEdges holds the relations/edges for other nodes in the graph.
type StarshipEdges struct {
	// AppearedIn holds the value of the appeared_in edge.
	AppearedIn []*Film `json:"appeared_in,omitempty"`
	// PilotedBy holds the value of the piloted_by edge.
	PilotedBy []*Person `json:"piloted_by,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// AppearedInOrErr returns the AppearedIn value or an error if the edge
// was not loaded in eager-loading.
func (e StarshipEdges) AppearedInOrErr() ([]*Film, error) {
	if e.loadedTypes[0] {
		return e.AppearedIn, nil
	}
	return nil, &NotLoadedError{edge: "appeared_in"}
}

// PilotedByOrErr returns the PilotedBy value or an error if the edge
// was not loaded in eager-loading.
func (e StarshipEdges) PilotedByOrErr() ([]*Person, error) {
	if e.loadedTypes[1] {
		return e.PilotedBy, nil
	}
	return nil, &NotLoadedError{edge: "piloted_by"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Starship) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case starship.FieldLength:
			values[i] = new(sql.NullFloat64)
		case starship.FieldID, starship.FieldCargoCapacity, starship.FieldCostInCredits, starship.FieldPassengerCapacity:
			values[i] = new(sql.NullInt64)
		case starship.FieldClass, starship.FieldConsumables, starship.FieldCrew, starship.FieldHyperdriveRating, starship.FieldManufacturer, starship.FieldMaxAtmospheringSpeed, starship.FieldMaximumMegalights, starship.FieldModel, starship.FieldName:
			values[i] = new(sql.NullString)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Starship", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Starship fields.
func (s *Starship) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case starship.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			s.ID = int(value.Int64)
		case starship.FieldCargoCapacity:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field cargo_capacity", values[i])
			} else if value.Valid {
				s.CargoCapacity = int(value.Int64)
			}
		case starship.FieldClass:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field class", values[i])
			} else if value.Valid {
				s.Class = value.String
			}
		case starship.FieldConsumables:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field consumables", values[i])
			} else if value.Valid {
				s.Consumables = value.String
			}
		case starship.FieldCostInCredits:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field cost_in_credits", values[i])
			} else if value.Valid {
				s.CostInCredits = int(value.Int64)
			}
		case starship.FieldCrew:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field crew", values[i])
			} else if value.Valid {
				s.Crew = value.String
			}
		case starship.FieldHyperdriveRating:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field hyperdrive_rating", values[i])
			} else if value.Valid {
				s.HyperdriveRating = value.String
			}
		case starship.FieldLength:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field length", values[i])
			} else if value.Valid {
				s.Length = value.Float64
			}
		case starship.FieldManufacturer:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field manufacturer", values[i])
			} else if value.Valid {
				s.Manufacturer = value.String
			}
		case starship.FieldMaxAtmospheringSpeed:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field max_atmosphering_speed", values[i])
			} else if value.Valid {
				s.MaxAtmospheringSpeed = value.String
			}
		case starship.FieldMaximumMegalights:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field maximum_megalights", values[i])
			} else if value.Valid {
				s.MaximumMegalights = value.String
			}
		case starship.FieldModel:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field model", values[i])
			} else if value.Valid {
				s.Model = value.String
			}
		case starship.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				s.Name = value.String
			}
		case starship.FieldPassengerCapacity:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field passenger_capacity", values[i])
			} else if value.Valid {
				s.PassengerCapacity = int(value.Int64)
			}
		}
	}
	return nil
}

// QueryAppearedIn queries the "appeared_in" edge of the Starship entity.
func (s *Starship) QueryAppearedIn() *FilmQuery {
	return (&StarshipClient{config: s.config}).QueryAppearedIn(s)
}

// QueryPilotedBy queries the "piloted_by" edge of the Starship entity.
func (s *Starship) QueryPilotedBy() *PersonQuery {
	return (&StarshipClient{config: s.config}).QueryPilotedBy(s)
}

// Update returns a builder for updating this Starship.
// Note that you need to call Starship.Unwrap() before calling this method if this Starship
// was returned from a transaction, and the transaction was committed or rolled back.
func (s *Starship) Update() *StarshipUpdateOne {
	return (&StarshipClient{config: s.config}).UpdateOne(s)
}

// Unwrap unwraps the Starship entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (s *Starship) Unwrap() *Starship {
	tx, ok := s.config.driver.(*txDriver)
	if !ok {
		panic("ent: Starship is not a transactional entity")
	}
	s.config.driver = tx.drv
	return s
}

// String implements the fmt.Stringer.
func (s *Starship) String() string {
	var builder strings.Builder
	builder.WriteString("Starship(")
	builder.WriteString(fmt.Sprintf("id=%v", s.ID))
	builder.WriteString(", cargo_capacity=")
	builder.WriteString(fmt.Sprintf("%v", s.CargoCapacity))
	builder.WriteString(", class=")
	builder.WriteString(s.Class)
	builder.WriteString(", consumables=")
	builder.WriteString(s.Consumables)
	builder.WriteString(", cost_in_credits=")
	builder.WriteString(fmt.Sprintf("%v", s.CostInCredits))
	builder.WriteString(", crew=")
	builder.WriteString(s.Crew)
	builder.WriteString(", hyperdrive_rating=")
	builder.WriteString(s.HyperdriveRating)
	builder.WriteString(", length=")
	builder.WriteString(fmt.Sprintf("%v", s.Length))
	builder.WriteString(", manufacturer=")
	builder.WriteString(s.Manufacturer)
	builder.WriteString(", max_atmosphering_speed=")
	builder.WriteString(s.MaxAtmospheringSpeed)
	builder.WriteString(", maximum_megalights=")
	builder.WriteString(s.MaximumMegalights)
	builder.WriteString(", model=")
	builder.WriteString(s.Model)
	builder.WriteString(", name=")
	builder.WriteString(s.Name)
	builder.WriteString(", passenger_capacity=")
	builder.WriteString(fmt.Sprintf("%v", s.PassengerCapacity))
	builder.WriteByte(')')
	return builder.String()
}

// Starships is a parsable slice of Starship.
type Starships []*Starship

func (s Starships) config(cfg config) {
	for _i := range s {
		s[_i].config = cfg
	}
}