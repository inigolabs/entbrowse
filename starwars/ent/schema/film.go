package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Film holds the schema definition for the Film entity.
type Film struct {
	ent.Schema
}

// Fields of the Film.
func (Film) Fields() []ent.Field {
	return []ent.Field{
		field.String("director"),
		field.Int("episode_id"),
		field.String("opening_crawl"),
		field.String("producer"),
		field.String("title"),
	}
}

// Edges of the Film.
func (Film) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("has_person", Person.Type),
		edge.To("has_planet", Planet.Type),
		edge.To("has_starship", Starship.Type),
		edge.To("has_vehicle", Vehicle.Type),
		edge.To("has_species", Species.Type),
	}
}
