package service

import (
	"context"
	"github.com/6rayWa1cher/shedevr-todo/backend/internal/dto"
	"github.com/6rayWa1cher/shedevr-todo/backend/internal/mapper"
	"github.com/6rayWa1cher/shedevr-todo/backend/internal/repository"
	"github.com/6rayWa1cher/shedevr-todo/backend/internal/utils"
	"github.com/go-faster/errors"
)

type TaskService interface {
	GetTasks(ctx context.Context, user string) ([]dto.Task, error)
	GetTask(ctx context.Context, taskId int64, user string) (*dto.Task, error)
	CreateTask(ctx context.Context, task dto.Task, user string) (*dto.Task, error)
	UpdateTask(ctx context.Context, task dto.Task, user string) (*dto.Task, error)
	DeleteTask(ctx context.Context, taskId int64, user string) error
}

type taskService struct {
	TaskService
	dao repository.Dao
}

func NewTaskService(dao repository.Dao) TaskService {
	return &taskService{dao: dao}
}

func (t *taskService) GetTasks(ctx context.Context, user string) ([]dto.Task, error) {
	tasks, err := t.dao.NewTaskQuery().GetTasks(ctx, user)
	if err != nil {
		return nil, errors.Wrap(err, "[service.GetTasks] sql error")
	}

	taskDtos, err := utils.MapWithError(tasks, mapper.TaskEntityToDto)
	if err != nil {
		return nil, errors.Wrap(err, "[service.GetTasks] mapping error")
	}

	return taskDtos, nil
}

func (t *taskService) GetTask(ctx context.Context, taskId int64, user string) (*dto.Task, error) {
	task, err := t.dao.NewTaskQuery().GetTask(ctx, taskId, user)
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

func (t *taskService) CreateTask(ctx context.Context, task dto.Task, user string) (*dto.Task, error) {
	taskEntity := mapper.TaskDtoToEntity(task)

	taskEntity.UserId = user

	taskId, err := t.dao.NewTaskQuery().CreateTask(ctx, taskEntity)
	if err != nil {
		return nil, errors.Wrap(err, "[service.CreateTask] sql error")
	}

	taskEntity.ID = *taskId

	taskDto, err := mapper.TaskEntityToDto(taskEntity)
	if err != nil {
		return nil, errors.Wrap(err, "[service.CreateTask] mapping error")
	}

	return &taskDto, nil
}

func (t *taskService) UpdateTask(ctx context.Context, task dto.Task, user string) (*dto.Task, error) {
	taskEntity := mapper.TaskDtoToEntity(task)

	updatedTaskEntity, err := t.dao.NewTaskQuery().UpdateTask(ctx, taskEntity, user)
	if errors.Is(err, repository.ErrNoResult) {
		return nil, ErrNotFound
	} else if err != nil {
		return nil, errors.Wrap(err, "[service.UpdateTask] sql error")
	}

	taskDto, err := mapper.TaskEntityToDto(*updatedTaskEntity)
	if err != nil {
		return nil, errors.Wrap(err, "[service.UpdateTask] mapping error")
	}

	return &taskDto, nil
}

func (t *taskService) DeleteTask(ctx context.Context, taskId int64, user string) error {
	err := t.dao.NewTaskQuery().DeleteTask(ctx, taskId, user)
	if err != nil {
		return errors.Wrap(err, "[service.DeleteTask] sql error")
	}

	return nil
}
