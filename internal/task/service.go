package task

import (
	"errors"
	"taskbuilder/internal/core/domain"
)

type taskService struct {
	taskRepository TaskRepository
}

func (t *taskService) Create(task domain.Task) (*domain.Task, error) {
	return t.taskRepository.Save(task)
}

func (t *taskService) Get(id string) (*domain.Task, error) {
	return t.taskRepository.FindOne(id)
}

func (t *taskService) Remove(id string) error {
	task, err := t.Get(id)
	if err != nil {
		return err
	}
	if task != nil {
		return t.taskRepository.Delete(task)
	}
	return errors.New("task not found")
}

func (t *taskService) GetAll() (*domain.Tasks, error) {
	return t.taskRepository.Find()
}

func (t *taskService) Update(data domain.Task) error {
	task, err := t.Get(data.ID)
	if err != nil {
		return err
	}
	if task != nil {
		return errors.New("task not found")
	}
	return t.taskRepository.Update(data)
}

func NewTaskService(taskRepository TaskRepository) TaskService {
	return &taskService{taskRepository}
}
