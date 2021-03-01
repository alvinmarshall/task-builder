package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"os"
	"taskbuilder/internal/core/domain"
	"taskbuilder/internal/core/service"
	"taskbuilder/internal/handler/taskhdl"
	"taskbuilder/internal/repository"
)

func main() {
	config := repository.Config{
		DBDriver:   "postgres",
		DBPassword: os.Getenv("DB_PASSWORD"),
		DBUsername: os.Getenv("DB_USER"),
		DBName:     os.Getenv("DB_NAME"),
		DBPort:     os.Getenv("DB_PORT"),
		DBHost:     os.Getenv("DB_HOST"),
	}
	persistence, _ := repository.NewPersistence(config)
	persistence.DB.AutoMigrate(domain.Task{})
	taskRepo := repository.NewTaskRepo(persistence)
	s := service.New(taskRepo)
	handler := taskhdl.New(s)

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.GET("/tasks/{id}", handler.Get)
	e.GET("/tasks", handler.GetAll)
	e.POST("/tasks", handler.Create)
	e.Logger.Fatal(e.Start(":3000"))

}
