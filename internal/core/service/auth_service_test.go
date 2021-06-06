package service

import (
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"taskbuilder/internal/core/domain"
	mock_port "taskbuilder/internal/mock"
	"taskbuilder/internal/server/rest/dto"
	"testing"
)

func fakeAuthUser() *domain.User {
	return &domain.User{
		BaseEntity: domain.BaseEntity{ID: "1"},
		Name:       "any name",
		Contact:    "any contact",
		Password:   "any password",
		Email:      "any email",
	}
}

func Test_authService_SignInWithEmailAndPassword(t *testing.T) {
	signInCredential := dto.SignInCredential{
		Email:    "any email",
		Password: "any password",
	}
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockUserService := mock_port.NewMockUserService(ctrl)
	mockEncryptionService := mock_port.NewMockEncryptionService(ctrl)
	mockUserService.EXPECT().FindByEmail(gomock.Any()).Return(fakeAuthUser(), nil)
	mockEncryptionService.EXPECT().ValidateEncryptedPassword(gomock.Any(), gomock.Any()).Return(true, nil)
	newAuthService := NewAuthService(mockUserService, mockEncryptionService)
	user, err := newAuthService.SignInWithEmailAndPassword(signInCredential)
	assert.NoError(t, err)
	assert.NotNil(t, user)
}

func Test_authService_SignInWithEmailAndPassword_Invalid_Password(t *testing.T) {
	signInCredential := dto.SignInCredential{
		Email:    "any email",
		Password: "invalid password",
	}
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockUserService := mock_port.NewMockUserService(ctrl)
	mockEncryptionService := mock_port.NewMockEncryptionService(ctrl)
	mockUserService.EXPECT().FindByEmail(gomock.Any()).Return(fakeAuthUser(), nil)
	mockEncryptionService.EXPECT().ValidateEncryptedPassword(gomock.Any(), gomock.Any()).Return(false, nil)
	newAuthService := NewAuthService(mockUserService, mockEncryptionService)
	user, err := newAuthService.SignInWithEmailAndPassword(signInCredential)
	if err != nil {
		assert.Equal(t, err.Error(), "invalid credentials")
	}
	assert.Nil(t, user)
}

func Test_authService_SignInWithEmailAndPassword_Invalid_Email(t *testing.T) {
	signInCredential := dto.SignInCredential{
		Email:    "invalid email",
		Password: "any password",
	}
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockUserService := mock_port.NewMockUserService(ctrl)
	mockUserService.EXPECT().FindByEmail(gomock.Any()).Return(nil, errors.New("user not found"))
	newAuthService := NewAuthService(mockUserService, nil)
	user, err := newAuthService.SignInWithEmailAndPassword(signInCredential)
	if err != nil {
		assert.Equal(t, err.Error(), "invalid credentials")
	}
	assert.Nil(t, user)
}

func Test_authService_SignInWithEmailAndPassword_ValidationError(t *testing.T) {
	signInCredential := dto.SignInCredential{
		Email:    "invalid email",
		Password: "any password",
	}
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockUserService := mock_port.NewMockUserService(ctrl)
	mockEncryptionService := mock_port.NewMockEncryptionService(ctrl)
	mockUserService.EXPECT().FindByEmail(gomock.Any()).Return(fakeAuthUser(), nil)
	mockEncryptionService.EXPECT().
		ValidateEncryptedPassword(gomock.Any(), gomock.Any()).Return(false, errors.New("an error occurred"))
	newAuthService := NewAuthService(mockUserService, mockEncryptionService)
	user, err := newAuthService.SignInWithEmailAndPassword(signInCredential)
	if err != nil {
		assert.Equal(t, err.Error(), "invalid credentials")
	}
	assert.Nil(t, user)
}
