package handler

import (
	"github.com/labstack/echo/v4"
	"go.uber.org/dig"
	"taskbuilder/internal/core/port"
)

func UserRouter(e *echo.Group, c *dig.Container) *echo.Group {
	var userSvc port.UserService
	c.Invoke(func(svc port.UserService) { userSvc = svc })
	route := e.Group("/users")
	{
		h := NewUserHandler(userSvc)
		route.GET("", h.Index)
		route.POST("", h.Register)
	}
	return route

}
