package rest

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"go.uber.org/dig"
	"gorm.io/gorm"
	"taskbuilder/internal/config"
	"taskbuilder/internal/storage/migration"
)

type server struct {
	engine    *echo.Echo
	container *dig.Container
}

func NewServer(e *echo.Echo, c *dig.Container) *server {
	return &server{
		engine:    e,
		container: c,
	}
}

func (s *server) InitializeDB() error {
	var db *gorm.DB
	err := s.container.Invoke(func(d *gorm.DB) { db = d })
	if err != nil {
		return err
	}
	err = migration.Migrate(db)
	if err != nil {
		return err
	}
	return nil

}

func (s *server) Start() error {
	var cfg *config.Config
	err := s.container.Invoke(func(c *config.Config) { cfg = c })
	if err != nil {
		return err
	}
	return s.engine.Start(fmt.Sprintf(":%s", cfg.Server.Port))
}
