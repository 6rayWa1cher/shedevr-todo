package service

import (
	"context"
	"github.com/6rayWa1cher/shedevr-todo/backend/internal/dto"
	"github.com/6rayWa1cher/shedevr-todo/backend/internal/mapper"
	"github.com/6rayWa1cher/shedevr-todo/backend/internal/repository"
	"github.com/go-faster/errors"
)

type TaskService interface {
	GetTask(ctx context.Context, taskId int64) (*dto.Task, error)
	CreateTask(ctx context.Context, task dto.Task) (*int64, error)
	UpdateTask(ctx context.Context, task dto.Task) (*dto.Task, error)
	DeleteTask(ctx context.Context, taskId int64, userId int64) error
}

type taskService struct {
	TaskService
	dao repository.Dao
}

func NewTaskService(dao repository.Dao) TaskService {
	return &taskService{dao: dao}
}

func (t *taskService) GetTask(ctx context.Context, taskId int64) (*dto.Task, error) {
	task, err := t.dao.NewTaskQuery().GetTask(ctx, taskId)
	if errors.Is(err, repository.ErrNoResult) {
		return nil, ErrNotFound
	} else if err != nil {
		return nil, errors.Wrap(err, "[service.GetTask] sql error")
	}
	taskDto, err := mapper.TaskEntityToDto(*task)
	if err != nil {
		return nil, errors.Wrap(err, "[service.GetTask] mapping error")
	}
	return &taskDto, nil
}

func (t *taskService) CreateTask(ctx context.Context, task dto.Task) (*int64, error) {
	//TODO implement me
	panic("implement me")
}

func (t *taskService) UpdateTask(ctx context.Context, task dto.Task) (*dto.Task, error) {
	//TODO implement me
	panic("implement me")
}

func (t *taskService) DeleteTask(ctx context.Context, taskId int64, userId int64) error {
	//TODO implement me
	panic("implement me")
}
