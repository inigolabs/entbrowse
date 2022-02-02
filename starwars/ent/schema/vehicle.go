package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Vehicle holds the schema definition for the Vehicle entity.
type Vehicle struct {
	ent.Schema
}

// Fields of the Vehicle.
func (Vehicle) Fields() []ent.Field {
	return []ent.Field{
		field.Int("cargo_capacity").Optional(),
		field.String("consumables"),
		field.Int("cost_in_credits"),
		field.String("crew"),
		field.Float("length").Optional().Nillable(),
		field.String("manufacturer"),
		field.String("max_atmosphering_speed"),
		field.String("model"),
		field.String("name"),
		field.Int("passenger_capacity").Optional(),
	}
}

// Edges of the Vehicle.
func (Vehicle) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("appeared_in", Film.Type).Ref("has_vehicle"),
		edge.From("piloted_by", Person.Type).Ref("piloted_vehicle"),
	}
}
