package app

import "github.com/6rayWa1cher/shedevr-todo/backend/pkg/oas"

type Service struct {
	oas.UnimplementedHandler
}

func NewService() *Service {
	return &Service{}
}
