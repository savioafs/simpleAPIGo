package entity

import (
	"testing"

	"github.com/test-go/testify/assert"
)

func TestNewUser(t *testing.T) {
	as := assert.New(t)
	user, err := NewUser("test", "test@test.com", "123456")
	as.Nil(err)
	as.NotNil(user)
	as.NotEmpty(user.ID)
	as.NotEmpty(user.Password)
	as.Equal("test", user.Name)
	as.Equal("test@test.com", user.Email)
}

func TestUserValidatePassword(t *testing.T) {
	as := assert.New(t)
	user, err := NewUser("test", "test@test.com", "123456")
	as.Nil(err)
	as.True(user.ValidatePassword("123456"))
	as.False(user.ValidatePassword("12345"))
	as.NotEqual("123456", user.Password)
}
