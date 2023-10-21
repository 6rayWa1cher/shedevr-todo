package mapper

import (
	"fmt"
	"github.com/6rayWa1cher/shedevr-todo/backend/internal/dto"
	"github.com/6rayWa1cher/shedevr-todo/backend/internal/entities"
	"github.com/6rayWa1cher/shedevr-todo/backend/pkg/oas"
)

func TaskEntityToDto(task entities.Task) (dto.Task, error) {
	status, err := completedStatusToType(task.Completed)
	if err != nil {
		return dto.Task{}, err
	}
	output := dto.Task{
		ID:        task.ID,
		Title:     task.Title,
		Text:      task.Text.String,
		Completed: status,
	}
	if task.CounterExist {
		counter := dto.Counter{}
		if task.CounterValue.Valid {
			counter.Value = task.CounterValue.Float64
		}
		if task.CounterMaxValue.Valid {
			counter.MaxValue = task.CounterMaxValue.Float64
		}
		if task.CounterScale.Valid {
			counter.Scale = task.CounterScale.String
		}
		output.Counter = &counter
	}
	return output, nil
}

func completedStatusToType(completedStatus string) (dto.CompletedStatusType, error) {
	switch completedStatus {
	case "no":
		return dto.CompletedNo, nil
	case "yes":
		return dto.CompletedYes, nil
	case "cancelled":
		return dto.CompletedCancelled, nil
	default:
		return "", fmt.Errorf("unknown value %s", completedStatus)
	}
}

func TaskDtoToOas(task dto.Task) oas.Task {
	output := oas.Task{
		ID:        oas.NewOptInt64(task.ID),
		Title:     task.Title,
		Text:      task.Text,
		Completed: oas.CompletedEnum(task.Completed),
	}
	if task.Counter != nil {
		output.Counter = oas.NewOptCounter(oas.Counter{
			Value:    oas.NewOptFloat64(task.Counter.Value),
			Scale:    oas.NewOptString(task.Counter.Scale),
			MaxValue: oas.NewOptFloat64(task.Counter.MaxValue),
		})
	}
	return output
}
