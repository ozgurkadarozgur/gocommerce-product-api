package app

import "github.com/ozgurkadarozgur/gocommerce-product-api/database"

func StartApp() {
	database.InitDatabase()
	InitRouter()
}
