package user

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
	"taskbuilder/internal/core/domain"
	"taskbuilder/internal/core/port"
	"taskbuilder/internal/core/service"
	"taskbuilder/internal/handler/dto"
	"taskbuilder/internal/types"
)

type userHandler struct {
	service     port.UserService
	authService service.JwtService
}

func NewUserHandler(service port.UserService, authService service.JwtService) *userHandler {
	return &userHandler{service, authService}
}

func (hdl *userHandler) Register(context echo.Context) error {
	req := &dto.UserRequest{}
	err := context.Bind(req)
	if err != nil {
		return err
	}
	u := domain.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
		Contact:  req.Contact,
		Tasks:    nil,
		Role:     domain.Role{Name: req.Role.String()},
	}

	user, err := hdl.service.Create(u)
	if err != nil {
		return context.JSONPretty(http.StatusBadRequest, map[string]interface{}{"message": err.Error()}, "")
	}
	payload := types.JwtPayload{
		Id:    user.ID,
		Email: user.Email,
		Role:  user.Role.Name,
	}
	token, err := hdl.authService.GenerateToken(payload)
	if err != nil {
		return context.JSONPretty(http.StatusInternalServerError, map[string]interface{}{"message": err.Error()}, "")
	}
	response := dto.UserCreationResponse{
		User:  user,
		Token: token,
	}
	return context.JSONPretty(http.StatusCreated, response, " ")
}

func (hdl *userHandler) Index(context echo.Context) error {
	users, err := hdl.service.GetAll()
	if err != nil {
		fmt.Println("user index: ", err.Error())
		return context.JSONPretty(http.StatusInternalServerError, nil, "")
	}
	return context.JSONPretty(http.StatusOK, users, " ")
}
