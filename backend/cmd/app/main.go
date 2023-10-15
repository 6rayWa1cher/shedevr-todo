package main

import (
	"fmt"

	"github.com/6rayWa1cher/shedevr-todo/backend/internal/app"
	"github.com/6rayWa1cher/shedevr-todo/backend/pkg"
)

func main() {
	fmt.Println("Hello, world!")
	var s pkg.ServerInterface = &app.Service{}
	s.GetTasks(nil)
}
