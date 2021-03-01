package port

import "taskbuilder/internal/core/domain"

type TaskRepository interface {
	FindOne(id string) (domain.Task, error)
	Find() (domain.Tasks, error)
	Save(task domain.Task) (domain.Task, error)
	Delete(task domain.Task) error
	Update(data domain.Task) error
}

type TaskService interface {
	Create(task domain.Task) (domain.Task, error)
	Get(id string) (domain.Task, error)
	Remove(id string) error
	GetAll() (domain.Tasks, error)
	Update(data domain.Task) error
}
