package domain

import (
	"taskbuilder/internal/core"
)

type Task struct {
	core.BaseEntity
	Title       string
	IsCompleted bool `json:"is_completed"`
}

type Tasks []Task
