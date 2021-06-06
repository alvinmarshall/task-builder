package orm

import (
	"gorm.io/gorm"
	"taskbuilder/internal/core/domain"
	"taskbuilder/internal/core/port"
)

type taskRepo struct {
	db *gorm.DB
}

func (t *taskRepo) Find(user domain.User) (*domain.Tasks, error) {
	tasks := &domain.Tasks{}
	err := t.db.Where("user_id = ?", user.ID).Find(&tasks).Error
	if err != nil {
		return nil, err
	}
	return tasks, nil
}

func (t *taskRepo) Save(task domain.Task) (*domain.Task, error) {
	create := t.db.Model(&task).Create(&task)
	err := create.Save(&task).Error
	if err != nil {
		return nil, err
	}
	return &task, nil
}

func (t *taskRepo) Delete(task *domain.Task) error {
	return t.db.Model(&task).Delete(&task).Error
}

func (t *taskRepo) Update(data domain.Task) error {
	return t.db.Model(&data).Save(&data).Error
}

func (t *taskRepo) FindOne(id string) (*domain.Task, error) {
	data := &domain.Task{}
	err := t.db.Model(data).Where("id = ?", id).First(data).Error
	if err != nil {
		return nil, err
	}
	return data, nil
}

func NewTaskRepo(db *gorm.DB) port.TaskRepository {
	return &taskRepo{db: db}
}
