package router

import (
	"github.com/gin-gonic/gin"
	"github.com/ozgurkadarozgur/gocommerce-product-api/domain/category/handler"
)

type CategoryRouter struct {
	handler handler.CategoryHandler
}

func (categoryRouter CategoryRouter) Init(r *gin.Engine) {

	r.GET("/categories", categoryRouter.handler.Index)

	r.POST("categories", categoryRouter.handler.Store)

}