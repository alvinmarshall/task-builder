package port

import "taskbuilder/internal/core/domain"

type UserRepository interface {
	Save(user domain.User) (*domain.User, error)
	FindOne(id string) (*domain.User, error)
	FindByEmail(email string) (*domain.User, error)
	Find() (*domain.Users, error)
	Remove(user *domain.User) error
	Update(user domain.User) error
}

type UserService interface {
	Create(user domain.User) (*domain.User, error)
	Get(id string) (*domain.User, error)
	GetAll() (*domain.Users, error)
	Delete(id string) error
	Update(data domain.User) error
	FindByEmail(email string) (*domain.User, error)
}
