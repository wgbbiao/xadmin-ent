package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// XadminContenttype holds the schema definition for the XadminContenttype entity.
type XadminContenttype struct {
	ent.Schema
}

// Fields of the XadminContenttype.
func (XadminContenttype) Fields() []ent.Field {
	return []ent.Field{
		field.String("app_label").NotEmpty(),
		field.String("model").NotEmpty(),
	}
}

// Edges of the XadminContenttype.
func (XadminContenttype) Edges() []ent.Edge {
	return nil
}

// Indexes of the ContentType.
func (XadminContenttype) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("app_label"),
		index.Fields("model"),
	}
}

func (XadminContenttype) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}
