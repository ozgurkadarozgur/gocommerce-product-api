package app

import (
	"github.com/gin-gonic/gin"
	"github.com/ozgurkadarozgur/gocommerce-product-api/domain/category/router"
	"log"
)

var (
	r = gin.Default()
)

func InitRouter() {


	new(router.CategoryRouter).Init(r)

	err := r.Run()
	if err != nil {
		log.Fatal(err)
	}
}