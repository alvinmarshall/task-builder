package orm

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"log"
	"taskbuilder/internal/core/domain"
	"taskbuilder/internal/core/port"
)

type userRepo struct {
	db *gorm.DB
}

func (u userRepo) Save(user domain.User) (*domain.User, error) {
	create := u.db.Model(&user).Create(&user)
	err := create.Save(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (u userRepo) FindOne(id string) (*domain.User, error) {
	user := &domain.User{}
	err := u.db.Model(user).Where("id = ?", id).Preload(clause.Associations).First(user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u userRepo) FindByEmail(email string) (*domain.User, error) {
	user := &domain.User{}
	err := u.db.Model(user).Where("email = ?", email).Preload(clause.Associations).First(user).Error
	if err != nil {
		log.Printf("email not found: %v", err.Error())
		return nil, err
	}
	return user, nil
}

func (u userRepo) Find() (*domain.Users, error) {
	users := &domain.Users{}
	err := u.db.Model(&domain.User{}).Preload(clause.Associations).Find(users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (u userRepo) Remove(user *domain.User) error {
	return u.db.Model(&user).Delete(&user).Error
}

func (u userRepo) Update(user domain.User) error {
	return u.db.Model(&user).Save(&user).Error
}

func NewUserRepo(db *gorm.DB) port.UserRepository {
	return &userRepo{db}
}
