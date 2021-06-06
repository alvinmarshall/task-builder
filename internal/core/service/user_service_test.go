package service

import (
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"taskbuilder/internal/core/domain"
	mock_port "taskbuilder/internal/mock"
	"testing"
)

func fakeUser() *domain.User {
	return &domain.User{
		BaseEntity: domain.BaseEntity{ID: "1"},
		Name:       "any name",
		Contact:    "any contact",
		Password:   "any password",
		Email:      "any email",
	}
}

func Test_userService_Create(t *testing.T) {
	actual := fakeUser()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockUserRepository := mock_port.NewMockUserRepository(ctrl)
	mockEncryptionService := mock_port.NewMockEncryptionService(ctrl)
	mockUserRepository.EXPECT().FindByEmail(gomock.Eq(actual.Email)).Return(nil, nil).Times(1)
	mockUserRepository.EXPECT().Save(gomock.Any()).Return(actual, nil).Times(1)
	mockEncryptionService.EXPECT().EncryptPassword(gomock.Any(), gomock.Any()).Return("hash_password", nil)
	newUserService := NewUserService(mockUserRepository, mockEncryptionService)
	expected, err := newUserService.Create(*actual)
	assert.NoError(t, err)
	assert.NotNil(t, expected)

}

func Test_userService_Create_failed(t *testing.T) {
	actual := fakeUser()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockUserRepository := mock_port.NewMockUserRepository(ctrl)
	mockUserRepository.EXPECT().FindByEmail(gomock.Eq(actual.Email)).Return(actual, nil).Times(1)
	mockUserRepository.EXPECT().Save(gomock.Any()).Return(actual, nil).Times(0)
	newUserService := NewUserService(mockUserRepository, nil)
	_, err := newUserService.Create(*actual)
	assert.Error(t, err)

}

func Test_userService_Get(t *testing.T) {
	actual := fakeUser()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockUserRepository := mock_port.NewMockUserRepository(ctrl)
	mockUserRepository.EXPECT().FindOne(gomock.Any()).Return(actual, nil)
	svc := NewUserService(mockUserRepository, nil)
	expected, err := svc.Get(actual.ID)
	assert.NoError(t, err)
	assert.NotNil(t, expected)
}

func Test_userService_GetAll(t *testing.T) {
	actual := &domain.Users{*fakeUser()}
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockUserRepository := mock_port.NewMockUserRepository(ctrl)
	mockUserRepository.EXPECT().Find().Return(actual, nil)
	newUserService := NewUserService(mockUserRepository, nil)
	expected, err := newUserService.GetAll()
	assert.NoError(t, err)
	assert.NotEmpty(t, expected)
}

func Test_userService_Delete(t *testing.T) {
	actual := fakeUser()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockUserRepository := mock_port.NewMockUserRepository(ctrl)
	mockUserRepository.EXPECT().FindOne(gomock.Any()).Return(actual, nil).Times(1)
	mockUserRepository.EXPECT().Remove(gomock.Any()).Return(nil).Times(1)
	newUserService := NewUserService(mockUserRepository, nil)
	err := newUserService.Delete(actual.ID)
	assert.NoError(t, err)
}

func Test_userService_Delete_failed(t *testing.T) {
	actual := fakeUser()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockUserRepository := mock_port.NewMockUserRepository(ctrl)
	mockUserRepository.EXPECT().FindOne(gomock.Any()).Return(nil, errors.New("user not found")).Times(1)
	mockUserRepository.EXPECT().Remove(gomock.Any()).Return(nil).Times(0)
	newUserService := NewUserService(mockUserRepository, nil)
	err := newUserService.Delete(actual.ID)
	assert.Error(t, err)
}

func Test_userService_Update(t *testing.T) {
	user := fakeUser()
	actual := user
	actual.Name = "update name"
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockUserRepository := mock_port.NewMockUserRepository(ctrl)
	mockUserRepository.EXPECT().FindOne(user.ID).Return(user, nil).Times(1)
	mockUserRepository.EXPECT().Update(gomock.Any()).Return(nil).Times(1)
	newUserService := NewUserService(mockUserRepository, nil)
	err := newUserService.Update(*actual)
	assert.NoError(t, err)
}

func Test_userService_Update_failed(t *testing.T) {
	user := fakeUser()
	actual := user
	actual.Name = "update name"
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockUserRepository := mock_port.NewMockUserRepository(ctrl)
	mockUserRepository.EXPECT().FindOne(user.ID).Return(nil, errors.New("user not found")).Times(1)
	mockUserRepository.EXPECT().Update(gomock.Any()).Return(errors.New("an error occurred")).Times(0)
	svc := NewUserService(mockUserRepository, nil)
	err := svc.Update(*actual)
	assert.Error(t, err)
}
