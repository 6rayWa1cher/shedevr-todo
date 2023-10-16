package repository

import (
	"github.com/6rayWa1cher/shedevr-todo/backend/internal/entities"
	sq "github.com/Masterminds/squirrel"
)

type TaskQuery interface {
	GetTask(id int64) (*entities.Task, error)
	CreateTask(task entities.Task) (*int64, error)
	UpdateTask(task entities.Task) (*entities.Task, error)
	DeleteTask(id int64) error
}

type taskQuery struct {
	qb func() sq.StatementBuilderType
}

var allFields = []string{
	"id", "title", "text", "completed", "counter_value", "counter_max_value", "counter_scale",
}

const taskTable = "public.task"

func (t taskQuery) GetTask(id int64) (*entities.Task, error) {
	t.qb().Select(allFields...).From(taskTable).Where(sq.Eq{"id": id})
}

func (t taskQuery) CreateTask(task entities.Task) (*int64, error) {
	//TODO implement me
	panic("implement me")
}

func (t taskQuery) UpdateTask(task entities.Task) (*entities.Task, error) {
	//TODO implement me
	panic("implement me")
}

func (t taskQuery) DeleteTask(id int64) error {
	//TODO implement me
	panic("implement me")
}
