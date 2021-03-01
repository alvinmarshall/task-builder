package service

import (
	"errors"
	"taskbuilder/internal/core/domain"
	"taskbuilder/internal/core/port"
)

type TaskService struct {
	taskRepository port.TaskRepository
}

func (s *TaskService) Create(task domain.Task) (domain.Task, error) {
	return s.taskRepository.Save(task)
}

func (s *TaskService) Get(id string) (domain.Task, error) {
	return s.taskRepository.FindOne(id)
}

func (s *TaskService) Remove(id string) error {
	task, err := s.Get(id)
	if err != nil {
		return err
	}
	if (domain.Task{}) != task {
		return s.taskRepository.Delete(task)
	}
	return errors.New("task not found")
}

func (s *TaskService) GetAll() (domain.Tasks, error) {
	return s.taskRepository.Find()
}

func (s *TaskService) Update(data domain.Task) error {
	task, err := s.Get(data.ID)
	if err != nil {
		return err
	}
	if (domain.Task{}) == task {
		return errors.New("task not found")
	}
	return s.taskRepository.Update(data)
}

func New(taskRepository port.TaskRepository) *TaskService {
	return &TaskService{taskRepository}
}

var _ port.TaskService = &TaskService{}
