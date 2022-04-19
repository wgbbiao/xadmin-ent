package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Permission holds the schema definition for the Permission entity.
type Permission struct {
	ent.Schema
}

// Fields of the Permission.
func (Permission) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").NotEmpty(),
		field.Uint("content_type_id"),
		field.String("model_code").NotEmpty(),
	}
}

// Edges of the Permission.
func (Permission) Edges() []ent.Edge {
	return nil
}
