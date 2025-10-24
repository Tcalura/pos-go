package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewUser(t *testing.T) {
	user, err := NewUser("John Doe", "john@example.com", "password123")

	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.NotEmpty(t, user.ID)
	assert.NotEmpty(t, user.Password)
	assert.Equal(t, "John Doe", user.Name)
	assert.Equal(t, "john@example.com", user.Email)

}

func TestValidatePassword(t *testing.T) {
	user, err := NewUser("John Doe", "john@example.com", "password123")

	assert.Nil(t, err)
	assert.True(t, user.ValidatePassword("password123"))
	assert.False(t, user.ValidatePassword("wrongpassword"))
	assert.NotEqual(t, "password123", user.Password)
}