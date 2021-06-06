package dto

import "taskbuilder/internal/core/domain"

type UserRequest struct {
	Name     string          `form:"name"`
	Email    string          `form:"email"`
	Password string          `form:"password"`
	Contact  string          `form:"contact"`
	Role     domain.RoleType `form:"role"`
}

type UserCreationResponse struct {
	User  *domain.User `json:"user"`
	Token string       `json:"token"`
}
