package database

import "github.com/savioafs/simpleAPIGo/internal/entity"

type UserStorer interface {
	Create(user *entity.User) error
	FindByEmail(email string) (*entity.User, error)
}
