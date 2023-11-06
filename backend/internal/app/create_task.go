package app

import (
	"context"
	"github.com/6rayWa1cher/shedevr-todo/backend/internal/mapper"
	"github.com/6rayWa1cher/shedevr-todo/backend/pkg/oas"
	"github.com/go-faster/errors"
)

func (s Service) CreateTask(ctx context.Context, req *oas.NewTask) (*oas.Task, error) {
	user, ok := s.securityService.GetCurrentRemoteUser(ctx)
	if !ok {
		return nil, errors.New("[app.CreateTask] user isn't presented")
	}

	taskDto, err := mapper.OasTaskToDto(req)
	if err != nil {
		return nil, errors.Wrap(err, "[app.CreateTask] mapping error")
	}

	task, err := s.taskService.CreateTask(ctx, taskDto, user)
	if err != nil {
		return nil, errors.Wrap(err, "[app.CreateTask] unexpected create error")
	}

	taskOas := mapper.TaskDtoToOas(*task)
	return &taskOas, nil
}
