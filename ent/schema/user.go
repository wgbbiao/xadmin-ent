package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("username"),
		field.String("password"),
		field.String("salt"),
		field.Bool("is_super"),
		field.Time("last_login_at"),
		field.Time("created_at").
			Default(time.Now),
		field.Time("updated_at").
			Default(time.Now).
			UpdateDefault(time.Now),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("roles", Role.Type).Ref("users"),
		edge.From("permissions", Permission.Type).Ref("users"),
	}
}

func (User) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("is_super"),
		index.Fields("username"),
	}
}
