package task

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"net/http"
	"taskbuilder/internal/core/domain"
	"taskbuilder/internal/core/port"
	dto2 "taskbuilder/internal/server/rest/dto"
	"taskbuilder/internal/types"
)

type taskHandler struct {
	service port.TaskService
}

func NewTaskHandler(service port.TaskService) *taskHandler {
	return &taskHandler{service}
}

func (hdl *taskHandler) Get(c echo.Context) error {
	param := c.Param("id")
	result, err := hdl.service.Get(param)
	if err != nil {
		return c.JSONPretty(http.StatusInternalServerError, map[string]string{"error": err.Error()}, " ")

	}
	response := dto2.TaskResponse{
		Id:          result.ID,
		Title:       result.Title,
		IsCompleted: result.IsCompleted,
	}
	return c.JSONPretty(http.StatusOK, response, " ")
}

func (hdl *taskHandler) GetAll(c echo.Context) error {
	token := c.Get("user").(*jwt.Token)
	claims := token.Claims.(*types.JwtClaim)
	user := domain.User{}
	user.ID = claims.Id
	tasks, err := hdl.service.GetAll(user)
	if err != nil {
		return c.JSONPretty(http.StatusInternalServerError, map[string]string{"error": err.Error()}, " ")
	}
	return c.JSONPretty(http.StatusOK, tasks, " ")
}

func (hdl *taskHandler) Create(c echo.Context) error {
	taskRequest := dto2.TaskRequest{}
	err := c.Bind(&taskRequest)
	if err != nil {
		return c.JSONPretty(http.StatusBadRequest, map[string]string{"error": err.Error()}, " ")
	}
	req := domain.Task{
		Title:       taskRequest.Title,
		IsCompleted: taskRequest.IsCompleted,
	}
	user := domain.User{}
	token := c.Get("user").(*jwt.Token)
	claim := token.Claims.(*types.JwtClaim)
	user.ID = claim.Id
	result, err := hdl.service.Create(req, user)
	if err != nil {
		return c.JSONPretty(http.StatusInternalServerError, map[string]string{"error": err.Error()}, " ")
	}
	response := dto2.TaskResponse{
		Id:          result.ID,
		Title:       result.Title,
		IsCompleted: result.IsCompleted,
	}
	return c.JSONPretty(http.StatusOK, response, " ")
}
