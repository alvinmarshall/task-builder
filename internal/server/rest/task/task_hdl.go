package task

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"os"
	"taskbuilder/internal/core/domain"
	"taskbuilder/internal/core/port"
	dto2 "taskbuilder/internal/server/rest/dto"
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
	user := domain.User{}
	// TODO Replace after authentication
	user.ID = os.Getenv("USER_ID")
	println("uuuusr", user.ID)
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
	// TODO Replace after authentication
	user.ID = os.Getenv("USER_ID")
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
