package service

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewEncryptionService(t *testing.T) {
	service := NewEncryptionService()
	assert.NotNil(t, service)
}

func Test_encryptionService_EncryptPassword(t *testing.T) {
	password := "any password"
	salt := 14
	service := NewEncryptionService()
	encryptPassword, err := service.EncryptPassword(password, salt)
	assert.NoError(t, err)
	assert.NotEqualf(t, encryptPassword, password, "password hashed")
}

func Test_encryptionService_ValidateEncryptedPassword(t *testing.T) {
	password := "any password"
	salt := 14
	service := NewEncryptionService()
	encryptPassword, err := service.EncryptPassword(password, salt)
	assert.NoError(t, err)
	validateEncryptedPassword, err := service.ValidateEncryptedPassword(password, encryptPassword)
	assert.NoError(t, err)
	assert.True(t, validateEncryptedPassword)
}

func Test_encryptionService_ValidateEncryptedPassword_InvalidHash(t *testing.T) {
	password := "any password"
	service := NewEncryptionService()
	validateEncryptedPassword, err := service.ValidateEncryptedPassword(password, "any hash")
	assert.Error(t, err)
	assert.False(t, validateEncryptedPassword)
}
