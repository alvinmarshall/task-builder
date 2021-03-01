package repository

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"log"
)

type Persistence struct {
	DB *gorm.DB
}

func NewPersistence(cfg Config) (*Persistence, error) {
	persistence := new(Persistence)
	connStr := fmt.Sprintf("host=%s port=%s dbname=%s user=%s sslmode=disable password=%s",
		cfg.DBHost, cfg.DBPort, cfg.DBName, cfg.DBUsername, cfg.DBPassword)
	db, err := gorm.Open(cfg.DBDriver, connStr)
	if err != nil {
		log.Fatalf("failed establishing db connection: %v", err)
	}
	fmt.Println("success establishing db connection")
	persistence.DB = db
	return persistence, nil
}
