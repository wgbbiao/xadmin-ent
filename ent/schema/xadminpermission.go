package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// XadminPermission holds the schema definition for the XadminPermission entity.
type XadminPermission struct {
	ent.Schema
}

// Fields of the XadminPermission.
func (XadminPermission) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").NotEmpty(),
		field.String("code").NotEmpty(),
		field.Time("created_at").
			Default(time.Now),
		field.Time("updated_at").
			Default(time.Now).
			UpdateDefault(time.Now),
		field.Int("xadmin_permission_content_type").Optional(),
	}
}

// Edges of the XadminPermission.
func (XadminPermission) Edges() []ent.Edge {

	return []ent.Edge{
		edge.To("ContentType", XadminContenttype.Type).
			Field("xadmin_permission_content_type").Unique(),
		edge.To("users", XadminUser.Type),
		edge.To("roles", XadminRole.Type),
	}
}

// Indexes of the Permission.
func (XadminPermission) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("name"),
		index.Fields("code"),
	}
}
