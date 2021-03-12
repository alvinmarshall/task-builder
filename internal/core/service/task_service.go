package service

import (
	"errors"
	"taskbuilder/internal/core/domain"
	"taskbuilder/internal/core/port"
)

type taskService struct {
	taskRepository port.TaskRepository
}

func (t *taskService) Create(task domain.Task, user domain.User) (*domain.Task, error) {
	task.UserID = user.ID
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

func (t *taskService) GetAll(user domain.User) (*domain.Tasks, error) {
	return t.taskRepository.Find(user)
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

func NewTaskService(taskRepository port.TaskRepository) port.TaskService {
	return &taskService{taskRepository}
}
