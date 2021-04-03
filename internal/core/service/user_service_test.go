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
	repo := mock_port.NewMockUserRepository(ctrl)
	repo.EXPECT().FindByEmail(gomock.Eq(actual.Email)).Return(nil, nil).Times(1)
	repo.EXPECT().Save(gomock.Any()).Return(actual, nil).Times(1)
	svc := NewUserService(repo)
	expected, err := svc.Create(*actual)
	assert.NoError(t, err)
	assert.NotNil(t, expected)

}

func Test_userService_Create_failed(t *testing.T) {
	actual := fakeUser()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	repo := mock_port.NewMockUserRepository(ctrl)
	repo.EXPECT().FindByEmail(gomock.Eq(actual.Email)).Return(actual, nil).Times(1)
	repo.EXPECT().Save(gomock.Any()).Return(actual, nil).Times(0)
	svc := NewUserService(repo)
	_, err := svc.Create(*actual)
	assert.Error(t, err)

}

func Test_userService_Get(t *testing.T) {
	actual := fakeUser()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	repo := mock_port.NewMockUserRepository(ctrl)
	repo.EXPECT().FindOne(gomock.Any()).Return(actual, nil)
	svc := NewUserService(repo)
	expected, err := svc.Get(actual.ID)
	assert.NoError(t, err)
	assert.NotNil(t, expected)
}

func Test_userService_GetAll(t *testing.T) {
	actual := &domain.Users{*fakeUser()}
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	repo := mock_port.NewMockUserRepository(ctrl)
	repo.EXPECT().Find().Return(actual, nil)
	svc := NewUserService(repo)
	expected, err := svc.GetAll()
	assert.NoError(t, err)
	assert.NotEmpty(t, expected)
}

func Test_userService_Delete(t *testing.T) {
	actual := fakeUser()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	repo := mock_port.NewMockUserRepository(ctrl)
	repo.EXPECT().FindOne(gomock.Any()).Return(actual, nil).Times(1)
	repo.EXPECT().Remove(gomock.Any()).Return(nil).Times(1)
	svc := NewUserService(repo)
	err := svc.Delete(actual.ID)
	assert.NoError(t, err)
}

func Test_userService_Delete_failed(t *testing.T) {
	actual := fakeUser()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	repo := mock_port.NewMockUserRepository(ctrl)
	repo.EXPECT().FindOne(gomock.Any()).Return(nil, errors.New("user not found")).Times(1)
	repo.EXPECT().Remove(gomock.Any()).Return(nil).Times(0)
	svc := NewUserService(repo)
	err := svc.Delete(actual.ID)
	assert.Error(t, err)
}

func Test_userService_Update(t *testing.T) {
	user := fakeUser()
	actual := user
	actual.Name = "update name"
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	repo := mock_port.NewMockUserRepository(ctrl)
	repo.EXPECT().FindOne(user.ID).Return(user, nil).Times(1)
	repo.EXPECT().Update(gomock.Any()).Return(nil).Times(1)
	svc := NewUserService(repo)
	err := svc.Update(*actual)
	assert.NoError(t, err)
}

func Test_userService_Update_failed(t *testing.T) {
	user := fakeUser()
	actual := user
	actual.Name = "update name"
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	repo := mock_port.NewMockUserRepository(ctrl)
	repo.EXPECT().FindOne(user.ID).Return(nil, errors.New("user not found")).Times(1)
	repo.EXPECT().Update(gomock.Any()).Return(errors.New("an error occurred")).Times(0)
	svc := NewUserService(repo)
	err := svc.Update(*actual)
	assert.Error(t, err)
}
