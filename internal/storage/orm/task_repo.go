package orm

import (
	"github.com/jinzhu/gorm"
	"taskbuilder/internal/core/domain"
	"taskbuilder/internal/task"
)

type taskRepo struct {
	db *gorm.DB
}

func (t *taskRepo) Find() (*domain.Tasks, error) {
	tasks := &domain.Tasks{}
	err := t.db.Find(&tasks).Error
	if err != nil {
		return nil, err
	}
	return tasks, nil
}

func (t *taskRepo) Save(task domain.Task) (*domain.Task, error) {
	err := t.db.Create(&task).Error
	if err != nil {
		return nil, err
	}
	return &task, nil
}

func (t *taskRepo) Delete(task *domain.Task) error {
	return t.db.Delete(task).Error
}

func (t *taskRepo) Update(data domain.Task) error {
	return t.db.Update(data).Error
}

func (t *taskRepo) FindOne(id string) (*domain.Task, error) {
	data := &domain.Task{}
	err := t.db.Find(data, id).Error
	if err != nil {
		return nil, err
	}
	return data, nil
}

//
//func (t *taskRepo) Find() (*task.Tasks, error) {
//	t.log.Debugf("fetch all task")
//
//	tasks := &task.Tasks{}
//	err := t.db.Find(&tasks).Error
//	if err != nil {
//		return nil, err
//	}
//	return tasks, nil
//}
//
//func (t *taskRepo) Save(task task.Task) (*task.Task, error) {
//	err := t.db.Create(&task).Error
//	if err != nil {
//		return nil, err
//	}
//	return &task, nil
//}
//
//func (t *taskRepo) Delete(task task.Task) error {
//	t.log.Debugf("deleting task with id: %s", task.ID)
//	return t.db.Delete(task).Error
//}
//
//func (t *taskRepo) Update(data task.Task) error {
//	t.log.Debugf("updating task with id: %s", data.ID)
//	return t.db.Update(data).Error
//}

func NewTaskRepo(db *gorm.DB) task.TaskRepository {
	return &taskRepo{db: db}
}
