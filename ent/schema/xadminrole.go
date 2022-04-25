package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// XadminRole holds the schema definition for the XadminRole entity.
type XadminRole struct {
	ent.Schema
}

// Fields of the XadminRole.
func (XadminRole) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").NotEmpty(),
		field.Time("created_at").
			Default(time.Now),
		field.Time("updated_at").
			Default(time.Now).
			UpdateDefault(time.Now),
	}
}

// Edges of the XadminRole.
func (XadminRole) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("users", XadminUser.Type),
		edge.From("permissions", XadminPermission.Type).Ref("roles"),
	}
}

// Indexes of the XadminRole.
func (XadminRole) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("name"),
	}
}
