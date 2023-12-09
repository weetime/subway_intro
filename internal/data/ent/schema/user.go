package schema

import (
	"entgo.io/ent/dialect"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id"),
		field.String("name").
			Default("unknown").
			Comment("姓名"),
		field.Uint32("age").
			Positive().
			Comment("年龄"),
		field.Time("created_at").
			Default(time.Now).SchemaType(map[string]string{
			dialect.MySQL: "datetime",
		}).Comment("创建时间"),
		field.Time("updated_at").
			Default(time.Now).SchemaType(map[string]string{
			dialect.MySQL: "datetime",
		}).Comment("更新时间"),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return nil
}
