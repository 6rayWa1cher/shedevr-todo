package mapper

import (
	"database/sql"
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
		Completed: status,
	}
	if task.Text.Valid {
		output.Text = &task.Text.String
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
		Completed: oas.CompletedEnum(task.Completed),
	}
	if task.Text != nil {
		output.Text = *task.Text
	}
	if task.Counter != nil {
		output.Counter = oas.NewOptCounter(oas.Counter{
			Value:    task.Counter.Value,
			Scale:    task.Counter.Scale,
			MaxValue: task.Counter.MaxValue,
		})
	}
	return output
}

func TaskDtoToEntity(task dto.Task) entities.Task {
	output := entities.Task{
		ID:        task.ID,
		Title:     task.Title,
		Completed: string(task.Completed),
	}
	if task.Text != nil {
		output.Text = sql.NullString{String: *task.Text, Valid: true}
	}
	if task.Counter != nil {
		output.CounterExist = true
		output.CounterValue = sql.NullFloat64{Float64: task.Counter.Value, Valid: true}
		output.CounterMaxValue = sql.NullFloat64{Float64: task.Counter.MaxValue, Valid: true}
		output.CounterScale = sql.NullString{String: task.Counter.Scale, Valid: true}
	}
	return output
}

type oasTask interface {
	GetTitle() string
	GetText() oas.OptString
	GetCompleted() oas.OptCompletedEnum
	GetCounter() oas.OptCounter
}

func OasTaskToDto(task oasTask) (dto.Task, error) {
	output := dto.Task{
		Title: task.GetTitle(),
	}
	if value, ok := task.GetText().Get(); ok {
		output.Text = &value
	}
	if value, ok := task.GetCompleted().Get(); ok {
		completedStatus, err := completedStatusToType(string(value))
		if err != nil {
			return dto.Task{}, err
		}
		output.Completed = completedStatus
	} else {
		output.Completed = dto.CompletedNo
	}

	if value, ok := task.GetCounter().Get(); ok {
		counter := dto.Counter{
			Value:    value.Value,
			MaxValue: value.MaxValue,
			Scale:    value.Scale,
		}
		output.Counter = &counter
	}
	return output, nil
}
