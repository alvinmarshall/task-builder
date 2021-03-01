package taskhdl

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"taskbuilder/internal/core/domain"
	"taskbuilder/internal/core/service"
	"taskbuilder/internal/handler/dto"
)

type httpHandler struct {
	service *service.TaskService
}

func New(service *service.TaskService) *httpHandler {
	return &httpHandler{service}
}

func (hdl *httpHandler) Get(c echo.Context) error {

	task, err := hdl.service.Get(c.Param("id"))
	if err != nil {
		return c.JSONPretty(http.StatusInternalServerError, map[string]string{"error": err.Error()}, " ")

	}
	response := dto.TaskResponse{
		Id:          task.ID,
		Title:       task.Title,
		IsCompleted: task.IsCompleted,
	}
	c.JSONPretty(http.StatusOK, response, " ")
	return nil
}

func (hdl *httpHandler) GetAll(c echo.Context) error {

	tasks, err := hdl.service.GetAll()
	if err != nil {
		return c.JSONPretty(http.StatusInternalServerError, map[string]string{"error": err.Error()}, " ")

	}

	c.JSONPretty(http.StatusOK, tasks, " ")
	return nil
}

func (hdl *httpHandler) Create(c echo.Context) error {
	taskRequest := dto.TaskRequest{}
	err := c.Bind(&taskRequest)
	if err != nil {
		return c.JSONPretty(http.StatusBadRequest, map[string]string{"error": err.Error()}, " ")
	}
	task := domain.Task{
		Title:       taskRequest.Title,
		IsCompleted: taskRequest.IsCompleted,
	}
	task, err = hdl.service.Create(task)
	if err != nil {
		return c.JSONPretty(http.StatusInternalServerError, map[string]string{"error": err.Error()}, " ")
	}
	response := dto.TaskResponse{
		Id:          task.ID,
		Title:       task.Title,
		IsCompleted: task.IsCompleted,
	}
	c.JSONPretty(http.StatusOK, response, " ")
	return nil
}
