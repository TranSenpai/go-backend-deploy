package schema

import (
	softdelete "go-backend/ent/soft-delete"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// ChatMessages holds the schema definition for the ChatMessages entity.
type ChatMessages struct {
	ent.Schema
}

// Fields of the ChatMessages.
func (ChatMessages) Fields() []ent.Field {
	return []ent.Field{
		field.Int("chat_group_id").StructTag("chatGroupId"),
		field.Int("user_id").StructTag("userId"),
		field.Text("message_text").StructTag("messageText"),

		field.Time("created_at").Default(time.Now).Immutable(),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
	}
}

// Edges of the ChatMessages.
func (ChatMessages) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("ChatGroups", ChatGroups.Type).Ref("ChatMessages").Field("chat_group_id").Unique().Required(),
		edge.From("Users", Users.Type).Ref("ChatMessages").Field("user_id").Unique().Required(),
	}
}

func (ChatMessages) Mixin() []ent.Mixin {
	return []ent.Mixin{
		softdelete.SoftDeleteMixin{},
	}
}
