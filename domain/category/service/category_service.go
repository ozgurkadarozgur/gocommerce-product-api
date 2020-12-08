package service

import (
	"github.com/ozgurkadarozgur/gocommerce-product-api/domain/category/form"
	"github.com/ozgurkadarozgur/gocommerce-product-api/domain/category/repository"
	"github.com/ozgurkadarozgur/gocommerce-product-api/ent"
)

type CategoryService struct {
	repository repository.CategoryRepository
}

func (service CategoryService) All() []*ent.Category {
	return service.repository.All()
}

func (service CategoryService) Create(form *form.CreateCategoryForm) *ent.Category {
	return service.repository.Create(form)
}