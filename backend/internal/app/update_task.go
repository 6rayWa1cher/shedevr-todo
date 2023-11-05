package app

import (
	"context"
	"github.com/6rayWa1cher/shedevr-todo/backend/internal/mapper"
	"github.com/6rayWa1cher/shedevr-todo/backend/internal/service"
	"github.com/6rayWa1cher/shedevr-todo/backend/pkg/oas"
	"github.com/go-faster/errors"
)

func (s Service) UpdateTask(
	ctx context.Context,
	req *oas.UpdateTask,
	params oas.UpdateTaskParams,
) (oas.UpdateTaskRes, error) {
	taskDto, err := mapper.OasTaskToDto(req)
	if err != nil {
		return nil, errors.Wrap(err, "[app.UpdateTask] bad completed enum")
	}
	taskDto.ID = params.ID
	task, err := s.taskService.UpdateTask(ctx, taskDto)
	if errors.Is(err, service.ErrNotFound) {
		return nil, err
	} else if err != nil {
		return nil, errors.Wrap(err, "[app.UpdateTask] unexpected update error")
	}
	taskOas := mapper.TaskDtoToOas(*task)
	return &taskOas, nil
}
