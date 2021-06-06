package service

import "golang.org/x/crypto/bcrypt"

type EncryptionService interface {
	EncryptPassword(password string, salt int) (string, error)
	ValidateEncryptedPassword(password, hashPassword string) (bool, error)
}

type encryptionService struct {
}

func (e encryptionService) EncryptPassword(password string, salt int) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), salt)
	if err != nil {
		return "", err
	}
	hashed := string(bytes)
	return hashed, nil
}

func (e encryptionService) ValidateEncryptedPassword(password, hashPassword string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(hashPassword), []byte(password))
	if err != nil {
		return false, err
	}
	return true, nil
}

func NewEncryptionService() EncryptionService {
	return &encryptionService{}
}
