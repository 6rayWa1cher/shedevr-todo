package app

import (
	"context"
	"github.com/6rayWa1cher/shedevr-todo/backend/internal/mapper"
	"github.com/6rayWa1cher/shedevr-todo/backend/pkg/oas"
	"github.com/go-faster/errors"
)

func (s Service) CreateTask(ctx context.Context, req *oas.NewTask) (*oas.Task, error) {
	taskDto, err := mapper.NewTaskToDto(*req)
	if err != nil {
		return nil, errors.Wrap(err, "[app.CreateTask] bad completed enum")
	}
	task, err := s.taskService.CreateTask(ctx, taskDto)
	if err != nil {
		return nil, errors.Wrap(err, "[app.CreateTask] unexpected create error")
	}
	taskOas := mapper.TaskDtoToOas(*task)
	return &taskOas, nil
}
