package main

import (
	"fmt"
	"github.com/6rayWa1cher/shedevr-todo/backend/pkg/oas"

	"github.com/6rayWa1cher/shedevr-todo/backend/internal/app"
)

func main() {
	fmt.Println("Hello, world!")
	var s oas.Handler = &app.Service{}
	s.GetTasks(nil)
}
