package service

import (
	"errors"
	"fmt"
	"taskbuilder/internal/core/domain"
	"taskbuilder/internal/core/port"
)

type userService struct {
	userRepository    port.UserRepository
	encryptionService EncryptionService
}

func (u *userService) FindByEmail(email string) (*domain.User, error) {
	return u.userRepository.FindByEmail(email)
}

func (u *userService) Create(user domain.User) (*domain.User, error) {
	result, err := u.userRepository.FindByEmail(user.Email)
	if err != nil {
		fmt.Println(err)
	}
	if result != nil {
		return nil, fmt.Errorf("email already in use: %s", user.Email)
	}
	salt := 14
	hashPassword, err := u.encryptionService.EncryptPassword(user.Password, salt)
	if err != nil {
		return nil, err
	}
	user.Password = hashPassword
	return u.userRepository.Save(user)
}

func (u *userService) Get(id string) (*domain.User, error) {
	return u.userRepository.FindOne(id)
}

func (u *userService) GetAll() (*domain.Users, error) {
	return u.userRepository.Find()
}

func (u *userService) Delete(id string) error {
	user, err := u.Get(id)
	if err != nil {
		return errors.New("user not found")
	}
	return u.userRepository.Remove(user)
}

func (u *userService) Update(data domain.User) error {
	_, err := u.Get(data.ID)
	if err != nil {
		return errors.New("user not found")
	}

	return u.userRepository.Update(data)
}

func NewUserService(userRepository port.UserRepository, encryptionService EncryptionService) port.UserService {
	return &userService{userRepository: userRepository, encryptionService: encryptionService}
}
