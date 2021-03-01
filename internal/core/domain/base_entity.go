package domain

import (
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
	"time"
)

type BaseEntity struct {
	ID        string `gorm:"type:uuid;primary_key;"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
}

func (base *BaseEntity) BeforeCreate(scope *gorm.Scope) error {
	uid := uuid.NewV4()
	return scope.SetColumn("ID", uid.String())
}
