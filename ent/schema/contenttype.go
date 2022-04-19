package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// ContentType holds the schema definition for the ContentType entity.
type ContentType struct {
	ent.Schema
}

// Fields of the ContentType.
func (ContentType) Fields() []ent.Field {
	return []ent.Field{
		field.String("app_label").NotEmpty(),
		field.String("model").NotEmpty(),
		field.Time("created_at").
			Default(time.Now),
		field.Time("updated_at").
			Default(time.Now).
			UpdateDefault(time.Now),
	}
}

// Edges of the ContentType.
func (ContentType) Edges() []ent.Edge {
	return nil
}

// Indexes of the ContentType.
func (ContentType) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("app_label"),
		index.Fields("model"),
	}
}
