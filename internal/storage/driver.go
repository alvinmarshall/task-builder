package storage

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"taskbuilder/internal/config"
)

func NewDataSource(c *config.Config) (*gorm.DB, error) {
	if c.DataSource.Use == "postgres" {
		if c.DataSource.Postgres.Enabled {
			return newPostgres(c)
		}
	}
	return nil, fmt.Errorf("DataSource Not Implemented %s", c.DataSource.Use)
}

func newPostgres(c *config.Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s port=%s dbname=%s user=%s sslmode=disable password=%s",
		c.DataSource.Postgres.Host,
		c.DataSource.Postgres.Port,
		c.DataSource.Postgres.Database,
		c.DataSource.Postgres.Username,
		c.DataSource.Postgres.Password)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, err
}
