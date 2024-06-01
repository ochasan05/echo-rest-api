package usecase

import (
	"go-rest-api-todo/model"
	"go-rest-api-todo/repository"
	"go-rest-api-todo/validator"
)

type ITaskUsecase interface {
	GetAllTasks(userId uint) ([]model.TaskResponse, error)
	GetTaskById(userId uint, taskId uint) (model.TaskResponse, error)
	CreatedTask(task model.Task) (model.TaskResponse, error)
	UpdateTask(task model.Task, userId uint, taskId uint) (model.TaskResponse, error)
	DeleteTask(userId uint, taskId uint) error
}

type taskUsecase struct {
	repository repository.ITaskRepository
	validator validator.ITaskValidator
}

func NewTaskUsecase(repository repository.ITaskRepository, validator validator.ITaskValidator) ITaskUsecase {
	return &taskUsecase{repository, validator}
}

func (usecase *taskUsecase) GetAllTasks(userId uint) ([]model.TaskResponse, error) {
	tasks := []model.Task{}
	if err := usecase.repository.GetAllTasks(&tasks, userId); err != nil {
		return nil, err
	}

	resTasks := []model.TaskResponse{}
	for _, v := range tasks {
		t := model.TaskResponse{
			ID:        v.ID,
			Title:     v.Title,
			CreatedAt: v.CreatedAt,
			UpdatedAt: v.UpdatedAt,
		}
		resTasks = append(resTasks, t)
	}

	return resTasks, nil
}

func (usecase *taskUsecase) GetTaskById(userId uint, taskId uint) (model.TaskResponse, error) {
	task := model.Task{}
	if err := usecase.repository.GetTaskById(&task, userId, taskId); err != nil {
		return model.TaskResponse{}, err
	}

	resTask := model.TaskResponse{
		ID:        task.ID,
		Title:     task.Title,
		CreatedAt: task.CreatedAt,
		UpdatedAt: task.UpdatedAt,
	}

	return resTask, nil
}

func (usecase *taskUsecase) CreatedTask(task model.Task) (model.TaskResponse, error) {
	if err := usecase.validator.TaskValidate(task); err != nil {
		return model.TaskResponse{}, err
	}

	if err := usecase.repository.CreateTask(&task); err != nil {
		return model.TaskResponse{}, err
	}

	resTask := model.TaskResponse{
		ID:        task.ID,
		Title:     task.Title,
		CreatedAt: task.CreatedAt,
		UpdatedAt: task.UpdatedAt,
	}

	return resTask, nil
}

func (usecase *taskUsecase) UpdateTask(task model.Task, userId uint, taskId uint) (model.TaskResponse, error) {
	if err := usecase.validator.TaskValidate(task); err != nil {
		return model.TaskResponse{}, err
	}
	
	if err := usecase.repository.UpdateTask(&task, userId, taskId); err != nil {
		return model.TaskResponse{}, err
	}

	resTask := model.TaskResponse{
		ID:        task.ID,
		Title:     task.Title,
		CreatedAt: task.CreatedAt,
		UpdatedAt: task.UpdatedAt,
	}

	return resTask, nil
}

func (usecase *taskUsecase) DeleteTask(userId uint, taskId uint) error {
	if err := usecase.repository.DeleteTask(userId, taskId); err != nil {
		return err
	}

	return nil
}