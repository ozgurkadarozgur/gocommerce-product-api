package presenter

import "github.com/ozgurkadarozgur/gocommerce-product-api/ent"

type CategoryPresenter struct {
	entity *ent.Category
}

func (category CategoryPresenter) Serialize() map[string]interface{} {
	return map[string]interface{}{
		"id":    category.entity.ID,
		"title": category.entity.Title + " serialized",
	}
}