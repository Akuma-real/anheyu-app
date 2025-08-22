// ent/schema/linktag.go
package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// LinkTag holds the schema definition for the LinkTag entity.
// This will create a non-conflicting "link_tags" table.
type LinkTag struct {
	ent.Schema
}

// Fields of the LinkTag.
func (LinkTag) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").
			Comment("标签名称").
			Unique().
			NotEmpty(),
		field.String("color").
			Comment("标签颜色 (e.g., #ff0000)").
			Default("#666666"),
	}
}

// Edges of the LinkTag.
func (LinkTag) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("links", Link.Type).
			Ref("tags"),
	}
}
