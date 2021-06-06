package auth

import (
	"github.com/labstack/echo/v4"
	"go.uber.org/dig"
	"log"
	"taskbuilder/internal/core/service"
)

func Bind(e *echo.Group, c *dig.Container) *echo.Group {
	var authSvc service.AuthService
	var jwtSvc service.JwtService
	err := c.Invoke(func(svc service.AuthService) { authSvc = svc })
	if err != nil {
		log.Fatal(err)
	}
	err = c.Invoke(func(svc service.JwtService) { jwtSvc = svc })
	if err != nil {
		log.Fatal(err)
	}

	routes := e.Group("/auth")
	{
		h := NewAuthHandler(authSvc, jwtSvc)
		routes.POST("", h.SignInWithEmailAndPassword)
	}
	return routes

}
