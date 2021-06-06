package user

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.uber.org/dig"
	"log"
	"taskbuilder/internal/core/port"
	"taskbuilder/internal/core/service"
)

func Bind(e *echo.Group, c *dig.Container) {
	var userSvc port.UserService
	var jwtSvc service.JwtService

	err := c.Invoke(func(svc port.UserService) { userSvc = svc })
	if err != nil {
		log.Fatal(err)
	}
	err = c.Invoke(func(svc service.JwtService) { jwtSvc = svc })
	if err != nil {
		log.Fatal(err)
	}

	h := NewUserHandler(userSvc, jwtSvc)
	routesProtected := e.Group("/users")
	{
		routesProtected.Use(middleware.JWTWithConfig(jwtSvc.GetJWTConfig()))
		routesProtected.GET("", h.Index)
	}

	routes := e.Group("/users/registration")
	{
		routes.POST("", h.Register)
	}

}
