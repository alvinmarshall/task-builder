package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"os"
	"taskbuilder/internal/di"
	"taskbuilder/internal/server"
)

func main() {
	err := run()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v", err)
		os.Exit(-1)
	}
}

func run() error {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	buildContainer := di.BuildContainer()
	s := server.NewServer(e, buildContainer)
	s.MapRoutes()

	err := s.InitializeDB()
	if err != nil {
		return err
	}
	return s.Start()
}
