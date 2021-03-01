package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Router struct {
	engine *echo.Echo
}

func NewRouter(engine *echo.Echo) *Router {
	engine.Use(middleware.Recover())

	return &Router{engine}
}
