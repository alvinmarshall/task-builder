package user

import (
	"github.com/labstack/echo/v4"
	"go.uber.org/dig"
	"log"
	"taskbuilder/internal/core/port"
	"taskbuilder/internal/core/service"
)

func Bind(e *echo.Group, c *dig.Container) *echo.Group {
	var userSvc port.UserService
	var authSvc service.JwtService
	err := c.Invoke(func(svc port.UserService) { userSvc = svc })
	if err != nil {
		log.Fatal(err)
	}
	err = c.Invoke(func(svc service.JwtService) { authSvc = svc })
	if err != nil {
		log.Fatal(err)
	}

	routes := e.Group("/users")
	{
		h := NewUserHandler(userSvc, authSvc)
		routes.GET("", h.Index)
		routes.POST("", h.Register)
	}
	return routes

}
