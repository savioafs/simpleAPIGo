package main

import (
	"net/http"

	"github.com/savioafs/simpleAPIGo/configs"
	"github.com/savioafs/simpleAPIGo/internal/entity"
	"github.com/savioafs/simpleAPIGo/internal/infra/database"
	"github.com/savioafs/simpleAPIGo/internal/infra/webserver/handlers"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	_, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}

	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&entity.User{}, &entity.Product{})

	productDB := database.NewProduct(db)
	productHandler := handlers.NewProductHandler(productDB)

	http.HandleFunc("/products", productHandler.CreateProduct)
	http.ListenAndServe(":8000", nil)
}
