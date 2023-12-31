package app

import (
	"context"
	"github.com/6rayWa1cher/shedevr-todo/backend/internal/mapper"
	"github.com/6rayWa1cher/shedevr-todo/backend/internal/service"
	"github.com/6rayWa1cher/shedevr-todo/backend/pkg/oas"
	"github.com/go-faster/errors"
)

func (s Service) GetTaskById(ctx context.Context, params oas.GetTaskByIdParams) (oas.GetTaskByIdRes, error) {
	user, ok := s.securityService.GetCurrentRemoteUser(ctx)
	if !ok {
		return nil, errors.New("[app.GetTaskById] user isn't presented")
	}

	task, err := s.taskService.GetTask(ctx, params.ID, user)
	if errors.Is(err, service.ErrNotFound) {
		return nil, err
	} else if err != nil {
		return nil, errors.Wrap(err, "[app.GetTaskById] get task unexpected error")
	}

	taskOas := mapper.TaskDtoToOas(*task)
	return &taskOas, nil
}
