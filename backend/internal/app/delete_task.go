package app

import (
	"context"
	"github.com/6rayWa1cher/shedevr-todo/backend/pkg/oas"
	"github.com/go-faster/errors"
)

func (s Service) DeleteTask(ctx context.Context, params oas.DeleteTaskParams) (oas.DeleteTaskRes, error) {
	user, ok := s.securityService.GetCurrentRemoteUser(ctx)
	if !ok {
		return nil, errors.New("[app.DeleteTask] user isn't presented")
	}

	taskId := params.ID

	err := s.taskService.DeleteTask(ctx, taskId, user)
	if err != nil {
		return nil, errors.Wrap(err, "[app.DeleteTask] get task unexpected error")
	}

	return &oas.DeleteTaskNoContent{}, nil
}
