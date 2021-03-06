package handler

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"taskbuilder/internal/core/domain"
	"taskbuilder/internal/handler/dto"
	"taskbuilder/internal/task"
)

type taskHandler struct {
	service task.TaskService
}

func NewTaskHandler(service task.TaskService) *taskHandler {
	return &taskHandler{service}
}

func (hdl *taskHandler) Get(c echo.Context) error {

	result, err := hdl.service.Get(c.Param("id"))
	if err != nil {
		return c.JSONPretty(http.StatusInternalServerError, map[string]string{"error": err.Error()}, " ")

	}
	response := dto.TaskResponse{
		Id:          result.ID,
		Title:       result.Title,
		IsCompleted: result.IsCompleted,
	}
	return c.JSONPretty(http.StatusOK, response, " ")
}

func (hdl *taskHandler) GetAll(c echo.Context) error {

	tasks, err := hdl.service.GetAll()
	if err != nil {
		return c.JSONPretty(http.StatusInternalServerError, map[string]string{"error": err.Error()}, " ")

	}

	return c.JSONPretty(http.StatusOK, tasks, " ")
}

func (hdl *taskHandler) Create(c echo.Context) error {
	taskRequest := dto.TaskRequest{}
	err := c.Bind(&taskRequest)
	if err != nil {
		return c.JSONPretty(http.StatusBadRequest, map[string]string{"error": err.Error()}, " ")
	}
	req := domain.Task{
		Title:       taskRequest.Title,
		IsCompleted: taskRequest.IsCompleted,
	}
	result, err := hdl.service.Create(req)
	if err != nil {
		return c.JSONPretty(http.StatusInternalServerError, map[string]string{"error": err.Error()}, " ")
	}
	response := dto.TaskResponse{
		Id:          result.ID,
		Title:       result.Title,
		IsCompleted: result.IsCompleted,
	}
	return c.JSONPretty(http.StatusOK, response, " ")
}
