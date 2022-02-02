package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Planet holds the schema definition for the Planet entity.
type Planet struct {
	ent.Schema
}

// Fields of the Planet.
func (Planet) Fields() []ent.Field {
	return []ent.Field{
		field.String("climate"),
		field.Int("diameter").Optional(),
		field.String("gravity"),
		field.String("name"),
		field.String("orbital_period"),
		field.Int("population").Optional(),
		field.String("rotation_period"),
		field.String("surface_water"),
		field.String("terrain"),
	}
}

// Edges of the Planet.
func (Planet) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("home_to", Person.Type),
		edge.From("appeared_in", Film.Type).Ref("has_planet"),
		edge.From("origin_of", Species.Type).Ref("originates_from"),
	}
}
