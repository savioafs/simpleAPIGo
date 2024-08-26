package database

import "github.com/savioafs/simpleAPIGo/internal/entity"

type UserStorer interface {
	Create(user *entity.User) error
	FindByEmail(email string) (*entity.User, error)
}

type ProductStorer interface {
	Create(product *entity.Product) error
	FindAll(page, limit int, sort string) ([]entity.Product, error)
	FindByID(id string) (*entity.Product, error)
	Update(product *entity.Product) error
	Delete(id string) error
}
