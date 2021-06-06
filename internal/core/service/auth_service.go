package service

import (
	"errors"
	"taskbuilder/internal/core/domain"
	"taskbuilder/internal/core/port"
	"taskbuilder/internal/server/rest/dto"
)

type AuthService interface {
	SignInWithEmailAndPassword(credential dto.SignInCredential) (*domain.User, error)
}

type authService struct {
	userService       port.UserService
	encryptionService EncryptionService
}

func (a *authService) SignInWithEmailAndPassword(credential dto.SignInCredential) (*domain.User, error) {
	user, err := a.userService.FindByEmail(credential.Email)
	if err != nil {
		return nil, errors.New("invalid credentials")
	}
	validateEncryptedPassword, err := a.encryptionService.ValidateEncryptedPassword(credential.Password, user.Password)
	if err != nil {
		return nil, errors.New("invalid credentials")
	}
	if !validateEncryptedPassword {
		return nil, errors.New("invalid credentials")
	}
	return user, nil
}

func NewAuthService(userService port.UserService, encryptionService EncryptionService) AuthService {
	return &authService{userService: userService, encryptionService: encryptionService}
}
