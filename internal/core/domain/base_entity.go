package domain

import (
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
	"time"
)

type BaseEntity struct {
	ID        string `json:"id" gorm:"type:uuid;primary_key;unique"`
	CreatedAt time.Time
	UpdatedAt time.Time  `json:"-"`
	DeletedAt *time.Time `sql:"index" json:"-"`
}

func (base *BaseEntity) BeforeCreate(tx *gorm.DB) error {
	uid := uuid.NewV4()
	tx.Statement.SetColumn("ID", uid.String())
	return nil
}
