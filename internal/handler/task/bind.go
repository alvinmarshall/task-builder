package task

import (
	"github.com/labstack/echo/v4"
	"go.uber.org/dig"
	"log"
	"taskbuilder/internal/core/port"
)

func Bind(e *echo.Group, c *dig.Container) *echo.Group {
	var taskSvc port.TaskService
	err := c.Invoke(func(svc port.TaskService) { taskSvc = svc })
	if err != nil {
		log.Fatal(err)
	}
	routes := e.Group("/tasks")
	{
		h := NewTaskHandler(taskSvc)
		routes.POST("", h.Create)
		routes.GET("", h.GetAll)
		routes.GET("/:id", h.Get)
	}
	return routes
}
