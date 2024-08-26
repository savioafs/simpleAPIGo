package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/savioafs/simpleAPIGo/internal/dto"
	"github.com/savioafs/simpleAPIGo/internal/entity"
	"github.com/savioafs/simpleAPIGo/internal/infra/database"
)

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
