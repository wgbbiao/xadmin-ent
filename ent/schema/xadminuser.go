package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// XadminUser holds the schema definition for the XadminUser entity.
type XadminUser struct {
	ent.Schema
}

// Fields of the XadminUser.
func (XadminUser) Fields() []ent.Field {
	return []ent.Field{
		field.String("username").Unique(),
		field.String("password").Sensitive(),
		field.String("salt").Sensitive(),
		field.Bool("is_super").Optional().
			Default(false).StructTag(`json:"is_super"`),
		field.Time("last_login_at").Optional().Nillable(),
	}
}

// Edges of the XadminUser.
func (XadminUser) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("roles", XadminRole.Type).Ref("users"),
		edge.From("permissions", XadminPermission.Type).Ref("users"),
	}
}

func (XadminUser) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("is_super"),
	}
}

func (XadminUser) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}
