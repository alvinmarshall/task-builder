package orm

import (
	"github.com/jinzhu/gorm"
	"log"
	"taskbuilder/internal/core/domain"
	"taskbuilder/internal/core/port"
)

type userRepo struct {
	db *gorm.DB
}

func (u userRepo) Save(user domain.User) (*domain.User, error) {
	create := u.db.Create(&user)
	err := create.Save(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (u userRepo) FindOne(id string) (*domain.User, error) {
	user := &domain.User{}
	err := u.db.Where("id = ?", id).Find(user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u userRepo) FindByEmail(email string) (*domain.User, error) {
	user := &domain.User{}
	err := u.db.Where("email = ?", email).Find(user).Error
	if err != nil {
		log.Printf("email not found: %v", err.Error())
		return nil, nil
	}
	return user, nil
}

func (u userRepo) Find() (*domain.Users, error) {
	users := &domain.Users{}
	err := u.db.Preload("Tasks").Debug().Find(users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (u userRepo) Remove(user *domain.User) error {
	return u.db.Delete(&user).Error
}

func (u userRepo) Update(user domain.User) error {
	return u.db.Update(&user).Error
}

func NewUserRepo(db *gorm.DB) port.UserRepository {
	return &userRepo{db}
}
