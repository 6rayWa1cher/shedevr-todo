package app

import (
	"context"
	"fmt"
	"github.com/6rayWa1cher/shedevr-todo/backend/pkg/oas"
)

func (s *Service) NewError(ctx context.Context, err error) *oas.ErrorStatusCode {
	return &oas.ErrorStatusCode{
		StatusCode: 500,
		Response: oas.Error{
			Code:    -1,
			Message: fmt.Errorf("internal server error: %w", err).Error(),
		},
	}
}
