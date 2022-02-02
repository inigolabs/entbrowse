package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Person holds the schema definition for the Person entity.
type Person struct {
	ent.Schema
}

// Fields of the Person.
func (Person) Fields() []ent.Field {
	return []ent.Field{
		field.String("birth_year"),
		field.String("eye_color"),
		field.String("gender"),
		field.String("hair_color"),
		field.Int("height"),
		field.Float("mass").Optional().Nillable(),
		field.String("name"),
		field.String("skin_color"),
	}
}

// Edges of the Person.
func (Person) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("piloted_starship", Starship.Type),
		edge.To("piloted_vehicle", Vehicle.Type),
		edge.To("is_of_type", Species.Type),
		edge.From("appeared_in", Film.Type).Ref("has_person"),
		edge.From("from_planet", Planet.Type).Ref("home_to"),
	}
}
