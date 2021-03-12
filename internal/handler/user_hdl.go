package handler

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
	"taskbuilder/internal/core/domain"
	"taskbuilder/internal/core/port"
	"taskbuilder/internal/handler/dto"
)

type userHandler struct {
	service port.UserService
}

func NewUserHandler(service port.UserService) *userHandler {
	return &userHandler{service}
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
	}

	user, err := hdl.service.Create(u)
	if err != nil {
		return context.JSONPretty(http.StatusBadRequest, map[string]interface{}{"message": err.Error()}, "")
	}
	return context.JSONPretty(http.StatusCreated, user, " ")
}

func (hdl *userHandler) Index(context echo.Context) error {
	users, err := hdl.service.GetAll()
	if err != nil {
		fmt.Println("user index: ", err.Error())
		return context.JSONPretty(http.StatusInternalServerError, nil, "")
	}
	return context.JSONPretty(http.StatusOK, users, " ")
}
