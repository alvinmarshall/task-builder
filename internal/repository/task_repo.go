package repository

import (
	"taskbuilder/internal/core/domain"
	"taskbuilder/internal/core/port"
)

type taskRepo struct {
	persistence *Persistence
}

func (t *taskRepo) FindOne(id string) (domain.Task, error) {
	task := domain.Task{}
	err := t.persistence.DB.Find(&task, id).Error
	if err != nil {
		return domain.Task{}, err
	}
	return task, nil
}

func (t *taskRepo) Find() (domain.Tasks, error) {
	tasks := domain.Tasks{}
	err := t.persistence.DB.Find(&tasks).Error
	if err != nil {
		return domain.Tasks{}, err
	}
	return tasks, nil
}

func (t *taskRepo) Save(task domain.Task) (domain.Task, error) {
	err := t.persistence.DB.Create(&task).Error
	if err != nil {
		return domain.Task{}, err
	}
	return task, nil
}

func (t *taskRepo) Delete(task domain.Task) error {
	return t.persistence.DB.Delete(task).Error
}

func (t *taskRepo) Update(data domain.Task) error {
	return t.persistence.DB.Update(data).Error
}

func NewTaskRepo(persistence *Persistence) *taskRepo {
	return &taskRepo{persistence}
}

var _ port.TaskRepository = &taskRepo{}
