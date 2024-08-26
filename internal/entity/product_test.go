package entity

import (
	"testing"

	"github.com/test-go/testify/assert"
)

func TestNewProduct(t *testing.T) {
	as := assert.New(t)
	product, err := NewProduct("product 1", 15.90)
	as.Nil(err)
	as.NotNil(product)
	as.NotEmpty(product.ID)
	as.Equal("product 1", product.Name)
	as.Equal(15.90, product.Price)
}

func TestProductWhenNameIsRequired(t *testing.T) {
	as := assert.New(t)
	product, err := NewProduct("", 15.90)
	as.Empty(product)
	as.Equal(ErrNameIsRequired, err)

}

func TestProductWhenPriceIsRequired(t *testing.T) {
	as := assert.New(t)
	product, err := NewProduct("product 2", 0.0)
	as.Nil(product)
	as.Equal(ErrPriceIsRequired, err)
}
