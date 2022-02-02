package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Species holds the schema definition for the Species entity.
type Species struct {
	ent.Schema
}

// Fields of the Species.
func (Species) Fields() []ent.Field {
	return []ent.Field{
		field.Int("average_height"),
		field.String("average_lifespan"),
		field.String("classification"),
		field.String("designation"),
		field.String("name"),
		field.String("skin_color"),
		field.String("eye_color"),
		field.String("hair_color"),
		field.String("language"),
	}
}

// Edges of the Species.
func (Species) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("originates_from", Planet.Type),
		edge.From("appeared_in", Film.Type).Ref("has_species"),
		edge.From("includes_person", Person.Type).Ref("is_of_type"),
	}
}
