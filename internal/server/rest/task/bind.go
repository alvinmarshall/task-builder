package task

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.uber.org/dig"
	"log"
	"taskbuilder/internal/core/port"
	"taskbuilder/internal/core/service"
)

func Bind(e *echo.Group, c *dig.Container) {
	var taskSvc port.TaskService
	var jwtSvc service.JwtService
	err := c.Invoke(func(svc port.TaskService) { taskSvc = svc })
	if err != nil {
		log.Fatal(err)
	}

	err = c.Invoke(func(svc service.JwtService) { jwtSvc = svc })
	if err != nil {
		log.Fatal(err)
	}

	h := NewTaskHandler(taskSvc)

	routesProtected := e.Group("/tasks")
	{
		routesProtected.Use(middleware.JWTWithConfig(jwtSvc.GetJWTConfig()))
		routesProtected.POST("", h.Create)
		routesProtected.GET("", h.GetAll)
		routesProtected.GET("/:id", h.Get)
	}
}
