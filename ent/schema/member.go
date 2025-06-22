package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Member holds the schema definition for the Member entity.
type Member struct {
	ent.Schema
}

// Fields of the Member.
func (Member) Fields() []ent.Field {
	return []ent.Field{
		field.String("email"),
		field.String("nickname"),
		field.String("password_hash"),
		field.String("status"),
	}
}

// Edges of the Member.
func (Member) Edges() []ent.Edge {
	return nil
}
