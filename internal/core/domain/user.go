package domain

type User struct {
	BaseEntity
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"-"`
	Contact  string `json:"contact"`
	//@one-to-many
	Tasks Tasks `gorm:"foreignKey:UserID"`
	//@one-to-one
	Role Role `json:"role"`
}

type Users []User

func (User) TableName() string {
	return "users"
}
