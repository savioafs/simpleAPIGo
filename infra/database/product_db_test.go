package database

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/savioafs/simpleAPIGo/internal/entity"
	"github.com/test-go/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestCreateNewProduct(t *testing.T) {
	as := assert.New(t)

	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}
	db.AutoMigrate(&entity.Product{})

	product, err := entity.NewProduct("product 1", 10.50)
	as.Nil(err)

	productDB := NewProduct(db)
	err = productDB.Create(product)
	as.Nil(err)
	as.NotEmpty(product.ID)
}

func TestFindAllProducts(t *testing.T) {
	as := assert.New(t)

	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}
	db.AutoMigrate(&entity.Product{})

	for i := 1; i < 24; i++ {
		product, err := entity.NewProduct(fmt.Sprintf("product %d", i), rand.Float64()*100)
		as.Nil(err)
		db.Create(product)
	}

	productDB := NewProduct(db)
	products, err := productDB.FindAll(1, 10, "asc")
	as.Nil(err)
	as.Len(products, 10)
	as.Equal("product 1", products[0].Name)
	as.Equal("product 10", products[9].Name)

	products, err = productDB.FindAll(2, 10, "asc")
	as.Nil(err)
	as.Len(products, 10)
	as.Equal("product 11", products[0].Name)
	as.Equal("product 20", products[9].Name)

	products, err = productDB.FindAll(3, 10, "asc")
	as.Nil(err)
	as.Len(products, 3)
	as.Equal("product 21", products[0].Name)
	as.Equal("product 23", products[2].Name)
}

func TestFindProductByID(t *testing.T) {
	as := assert.New(t)
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}

	db.AutoMigrate(&entity.Product{})

	product, err := entity.NewProduct("Product 1", 20.00)
	as.Nil(err)

	db.Create(product)

	productDB := NewProduct(db)

	product, err = productDB.FindByID(product.ID.String())
	as.Nil(err)
	as.Equal("Product 1", product.Name)
}

func TestUpdateProduct(t *testing.T) {
	as := assert.New(t)

	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}

	db.AutoMigrate(&entity.Product{})

	product, err := entity.NewProduct("product 1", 300.00)
	as.Nil(err)

	db.Create(product)

	productDB := NewProduct(db)
	product.Name = "product 2"
	err = productDB.Update(product)
	as.Nil(err)

	product, err = productDB.FindByID(product.ID.String())
	as.Nil(err)
	as.Equal("product 2", product.Name)
}

func TestDeleteProduct(t *testing.T) {
	as := assert.New(t)

	db, err := gorm.Open(sqlite.Open("file::memory"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}

	db.AutoMigrate(&entity.Product{})

	product, err := entity.NewProduct("product for delete", 200.00)
	as.Nil(err)

	db.Create(product)

	productDB := NewProduct(db)

	err = productDB.Delete(product.ID.String())
	as.Nil(err)

	_, err = productDB.FindByID(product.ID.String())
	as.Error(err)
}
