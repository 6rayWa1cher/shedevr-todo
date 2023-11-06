package app

import (
	"context"
	"github.com/6rayWa1cher/shedevr-todo/backend/internal/mapper"
	"github.com/6rayWa1cher/shedevr-todo/backend/internal/utils"
	"github.com/6rayWa1cher/shedevr-todo/backend/pkg/oas"
	"github.com/go-faster/errors"
)

func (s Service) GetTasks(ctx context.Context) ([]oas.Task, error) {
	user, ok := s.securityService.GetCurrentRemoteUser(ctx)
	if !ok {
		return nil, errors.New("[app.GetTasks] user isn't presented")
	}

	tasks, err := s.taskService.GetTasks(ctx, user)
	if err != nil {
		return nil, errors.Wrap(err, "[app.GetTasks] get task unexpected error")
	}

	tasksOas := utils.Map(tasks, mapper.TaskDtoToOas)
	return tasksOas, nil
}
