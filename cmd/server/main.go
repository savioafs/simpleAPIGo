package main

import (
	"encoding/json"
	"net/http"

	"github.com/savioafs/simpleAPIGo/configs"
	"github.com/savioafs/simpleAPIGo/infra/database"
	"github.com/savioafs/simpleAPIGo/internal/dto"
	"github.com/savioafs/simpleAPIGo/internal/entity"
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
	productHandler := NewProductHandler(productDB)

	http.HandleFunc("/products", productHandler.CreateProduct)
	http.ListenAndServe(":8000", nil)

}

type ProductHandler struct {
	ProductDB database.ProductStorer
}

func NewProductHandler(db database.ProductStorer) *ProductHandler {
	return &ProductHandler{
		ProductDB: db,
	}
}

func (ph *ProductHandler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var product dto.CreateProductInput
	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}

	p, err := entity.NewProduct(product.Name, product.Price)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}

	err = ph.ProductDB.Create(p)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusCreated)
}
