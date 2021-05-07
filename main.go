package main

import (
	"database/sql"

	db2 "github.com/deividroger/go-hexagonal/adapters/db"
	"github.com/deividroger/go-hexagonal/application"
)

func main() {

	db, _ := sql.Open("sqlite3", "sqlite3.db")

	productAdapter := db2.NewProductDb(db)

	productService := application.NewProductService(productAdapter)

	product, _ := productService.Create("product example", 30)

	product.Enable()

	productService.Persistence.Save(product)

}
