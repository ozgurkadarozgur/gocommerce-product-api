package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/field"
)

// Category holds the schema definition for the Category presenter.
type Category struct {
	ent.Schema
}

// Fields of the Category.
func (Category) Fields() []ent.Field {
	return []ent.Field{
		field.String("title").
			MaxLen(255),
	}
}

// Edges of the Category.
func (Category) Edges() []ent.Edge {
	return nil
}
