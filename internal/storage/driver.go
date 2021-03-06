package storage

import (
	"errors"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"taskbuilder/internal/config"
)

func NewDataSource(c *config.Config) (*gorm.DB, error) {
	if c.DataSource.Use == "postgres" {
		return newPostgres(c)
	}
	return nil, errors.New(fmt.Sprintf("DataSource Not Implemented %s", c.DataSource.Use))
}

func newPostgres(c *config.Config) (*gorm.DB, error) {
	connStr := fmt.Sprintf("host=%s port=%s dbname=%s user=%s sslmode=disable password=%s",
		c.DataSource.Postgres.Host,
		c.DataSource.Postgres.Port,
		c.DataSource.Postgres.Database,
		c.DataSource.Postgres.Username,
		c.DataSource.Postgres.Password)
	db, err := gorm.Open(c.DataSource.Postgres.Dialect, connStr)
	if err != nil {
		return nil, err
	}
	return db, err
}
