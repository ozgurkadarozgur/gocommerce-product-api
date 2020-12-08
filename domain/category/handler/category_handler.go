package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/ozgurkadarozgur/gocommerce-product-api/domain/category/form"
	"github.com/ozgurkadarozgur/gocommerce-product-api/domain/category/service"
)

type CategoryHandler struct {
	service service.CategoryService
}

func (handler CategoryHandler) Index(ctx *gin.Context) {
	response := handler.service.All()
	ctx.JSON(200, response)
}

func (handler CategoryHandler) Store(ctx *gin.Context) {

	createForm := new(form.CreateCategoryForm)
	err := ctx.ShouldBindJSON(createForm)

	if err != nil {
		ctx.JSON(422, err.Error())
		return
	}

	result := handler.service.Create(createForm)
	ctx.JSON(200, result)
}

func (handler CategoryHandler) Show(ctx *gin.Context) {


	/*category := &presenter.Category{
		Id:    1,
		Title: "naber",
	}

	ctx.JSON(200, category.Serialize())*/
}
