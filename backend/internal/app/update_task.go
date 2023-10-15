package app

import (
	"context"
	"errors"
	"github.com/6rayWa1cher/shedevr-todo/backend/pkg/oas"
)

func (s Service) UpdateTask(
	ctx context.Context,
	req *oas.UpdateTask,
	params oas.UpdateTaskParams,
) (oas.UpdateTaskRes, error) {
	return nil, errors.New("123")
}
