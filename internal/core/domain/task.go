package domain

type Task struct {
	BaseEntity
	Title       string
	IsCompleted bool   `json:"is_completed"`
	UserID      string `gorm:"type:uuid" json:"-"`
}

type Tasks []Task

func (Task) TableName() string {
	return "tasks"
}
