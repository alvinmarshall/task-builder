package domain

import "golang.org/x/crypto/bcrypt"

type User struct {
	BaseEntity
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"-"`
	Contact  string `json:"contact"`
}

type Users []User

func (User) TableName() string {
	return "users"
}

func (user *User) HashPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return err
	}
	user.Password = string(bytes)
	return nil
}

func (user *User) checkPassword(password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return err
	}
	return nil
}
