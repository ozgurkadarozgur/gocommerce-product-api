package repository

import (
	"context"
	"github.com/ozgurkadarozgur/gocommerce-product-api/database"
	"github.com/ozgurkadarozgur/gocommerce-product-api/domain/category/form"
	"github.com/ozgurkadarozgur/gocommerce-product-api/ent"
	"github.com/ozgurkadarozgur/gocommerce-product-api/ent/category"
)

type CategoryRepository struct{}

var (
	db = database.Client()
)

func (repository CategoryRepository) All() []*ent.Category {
	_categories, err := db.Category.Query().
		Order(ent.Desc(category.FieldID)).
		All(context.Background())

	if err != nil {
		panic(err)
	}

	return _categories

}

func (repository CategoryRepository) Create(form *form.CreateCategoryForm) *ent.Category {
	_category, err := db.Category.Create().
		SetTitle(form.Title).
		Save(context.Background())

	if err != nil {
		panic(err)
	}

	return _category
}
