package database

import (
	"testing"

	"github.com/savioafs/simpleAPIGo/internal/entity"
	"github.com/test-go/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestNewUser(t *testing.T) {
	as := assert.New(t)
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		as.Error(err)
	}

	db.AutoMigrate(&entity.User{})

	user, err := entity.NewUser("John Doe", "john@doe.com", "123456")
	as.Nil(err)
	as.NotNil(user)

	userDB := NewUser(db)

	err = userDB.Create(user)
	as.Nil(err)

	var userFound entity.User
	err = db.First(&userFound, "id = ?", user.ID).Error
	as.Nil(err)
	as.Equal(user.Name, userFound.Name)
	as.Equal(user.Email, userFound.Email)
	as.NotEmpty(user.Password, userFound.Password)
}

func TestFindByEmail(t *testing.T) {
	as := assert.New(t)
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		as.Error(err)
	}

	db.AutoMigrate(&entity.User{})

	user, _ := entity.NewUser("John Doe", "john@doe.com", "123456")
	userDB := NewUser(db)
	err = userDB.Create(user)
	as.Nil(err)

	userFound, err := userDB.FindByEmail(user.Email)
	as.Nil(err)
	as.Equal(user.Name, userFound.Name)
	as.Equal(user.Email, userFound.Email)
	as.NotEmpty(user.Password)
}
