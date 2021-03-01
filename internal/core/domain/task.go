package domain

type Task struct {
	BaseEntity
	Title       string
	IsCompleted bool `json:"is_completed"`
}

type Tasks []Task
