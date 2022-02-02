package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Starship holds the schema definition for the Starship entity.
type Starship struct {
	ent.Schema
}

// Fields of the Starship.
func (Starship) Fields() []ent.Field {
	return []ent.Field{
		field.Int("cargo_capacity").Optional(),
		field.String("class"),
		field.String("consumables"),
		field.Int("cost_in_credits"),
		field.String("crew"),
		field.String("hyperdrive_rating"),
		field.Float("length"),
		field.String("manufacturer"),
		field.String("max_atmosphering_speed"),
		field.String("maximum_megalights").Optional(),
		field.String("model"),
		field.String("name"),
		field.Int("passenger_capacity").Optional(),
	}
}

// Edges of the Starship.
func (Starship) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("appeared_in", Film.Type).Ref("has_starship"),
		edge.From("piloted_by", Person.Type).Ref("piloted_starship"),
	}
}
