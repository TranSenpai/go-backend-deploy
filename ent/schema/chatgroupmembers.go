package schema

import (
	softdelete "go-backend/ent/soft-delete"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// ChatGroupMembers holds the schema definition for the ChatGroupMembers entity.
type ChatGroupMembers struct {
	ent.Schema
}

// Fields of the ChatGroupMembers.
func (ChatGroupMembers) Fields() []ent.Field {
	return []ent.Field{
		field.Int("chat_group_id").StructTag("chatGroupId"),
		field.Int("user_id").StructTag("userId"),

		field.Time("created_at").Default(time.Now).Immutable(),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
	}
}

// Edges of the ChatGroupMembers.
func (ChatGroupMembers) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("ChatGroups", ChatGroups.Type).Ref("ChatGroupMembers").Field("chat_group_id").Unique().Required(),
		edge.From("Users", Users.Type).Ref("ChatGroupMembers").Field("user_id").Unique().Required(),
	}
}

func (ChatGroupMembers) Mixin() []ent.Mixin {
	return []ent.Mixin{
		softdelete.SoftDeleteMixin{},
	}
}
