package app

import (
	"context"
	"github.com/6rayWa1cher/shedevr-todo/backend/internal/service"
	"github.com/6rayWa1cher/shedevr-todo/backend/pkg/oas"
	"github.com/go-faster/errors"
)

func (s Service) NewError(ctx context.Context, err error) *oas.ErrorStatusCode {
	if errors.Is(err, service.ErrNotFound) {
		return &oas.ErrorStatusCode{
			StatusCode: 404,
			Response: oas.Error{
				Code:    1,
				Message: "not found",
			},
		}
	}
	return &oas.ErrorStatusCode{
		StatusCode: 500,
		Response: oas.Error{
			Code:    -1,
			Message: errors.Wrap(err, "internal server error").Error(),
		},
	}
}
