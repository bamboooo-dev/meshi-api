package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Restaurant holds the schema definition for the Restaurant entity.
type Restaurant struct {
	ent.Schema
}

// Fields of the Restaurant.
func (Restaurant) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").NotEmpty().Unique().Immutable(),
		field.String("name").NotEmpty(),
		field.String("url"),
		field.String("phone"),
		field.String("price"),
	}
}

// Edges of the Restaurant.
func (Restaurant) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("likes", Like.Type),
	}
}
