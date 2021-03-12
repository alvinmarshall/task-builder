package dto

type UserRequest struct {
	Name     string `form:"name"`
	Email    string `form:"email"`
	Password string `form:"password"`
	Contact  string `form:"contact"`
}
