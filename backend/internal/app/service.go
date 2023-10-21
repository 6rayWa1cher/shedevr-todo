package app

import (
	"github.com/6rayWa1cher/shedevr-todo/backend/internal/service"
	"github.com/6rayWa1cher/shedevr-todo/backend/pkg/oas"
)

type Service struct {
	oas.UnimplementedHandler
	taskService service.TaskService
}

func NewService(taskService service.TaskService) *Service {
	return &Service{
		taskService: taskService,
	}
}
