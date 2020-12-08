package database

import (
	"context"
	"fmt"
	"log"

	"github.com/ozgurkadarozgur/gocommerce-product-api/ent"

	_ "github.com/lib/pq"
)

const (
	HOST     = "postgres"
	PORT     = 5432
	USER     = "postgres"
	DB       = "gocommerce_product"
	PASSWORD = "123456"
)

var (
	client *ent.Client
)

func Client() *ent.Client {
	if client == nil {
		InitDatabase()
		log.Println("db:init")
	}
	return client
}

func InitDatabase() {
	var err error
	conn := fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=disable", HOST, PORT, USER, DB, PASSWORD)
	client, err = ent.Open("postgres", conn)
	if err != nil {
		log.Fatal(err)
	}

	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
}
