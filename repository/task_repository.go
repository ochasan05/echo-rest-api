package repository

import (
	"fmt"
	"go-rest-api-todo/model"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type ITaskRepository interface {
	GetAllTasks(tasks *[]model.Task, userId uint) error
	GetTaskById(task *model.Task, userId uint, taskId uint) error
	CreateTask(task *model.Task) error
	UpdateTask(task *model.Task, userId uint, taskId uint) error
	DeleteTask(userId uint, taskId uint) error
}

type taskRepository struct {
	db *gorm.DB
}

func NewTaskRepository(db *gorm.DB) ITaskRepository {
	return &taskRepository{db}
}

func (repository *taskRepository) GetAllTasks(tasks *[]model.Task, userId uint) error {
	if err := repository.db.Joins("User").Where("user_id=?", userId).Order("created_at").Find(tasks).Error; err != nil {
		return err
	}

	return nil
}

func (repository *taskRepository) GetTaskById(task *model.Task, userId uint, taskId uint) error {
	if err := repository.db.Joins("User").Where("user_id=?", userId).First(task, taskId).Error; err != nil {
		return err
	}

	return nil
}

func (repository *taskRepository) CreateTask(task *model.Task) error {
	if err := repository.db.Create(task).Error; err != nil {
		return err
	}

	return nil
}

func (repository *taskRepository) UpdateTask(task *model.Task, userId uint, taskId uint) error {
	result := repository.db.Model(task).Clauses(clause.Returning{}).Where("id=? AND user_id=?", taskId, userId).Update("title", task.Title)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return fmt.Errorf("record does not exist")
	}

	return nil
}

func (repository *taskRepository) DeleteTask(userId uint, taskId uint) error {
	result := repository.db.Where("id=? AND user_id=?", taskId, userId).Delete(&model.Task{})

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return fmt.Errorf("record does not exist")
	}

	return nil
}
