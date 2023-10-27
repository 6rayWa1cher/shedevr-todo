package repository

import (
	"context"
	"database/sql"
	"github.com/6rayWa1cher/shedevr-todo/backend/internal/entities"
	sq "github.com/Masterminds/squirrel"
	"github.com/go-faster/errors"
	"github.com/jmoiron/sqlx"
)

type TaskQuery interface {
	GetTask(ctx context.Context, id int64) (*entities.Task, error)
	CreateTask(ctx context.Context, task entities.Task) (*int64, error)
	UpdateTask(ctx context.Context, task entities.Task) (*entities.Task, error)
	DeleteTask(ctx context.Context, id int64) error
}

type taskQuery struct {
	TaskQuery
	qb sq.StatementBuilderType
	db *sqlx.DB
}

var allFields = []string{
	"id", "title", "text", "completed", "counter_value", "counter_max_value", "counter_scale", "counter_exist",
}

const taskTable = "todo.task"

func (t *taskQuery) GetTask(ctx context.Context, id int64) (*entities.Task, error) {
	task := entities.Task{}
	q, args, err := t.qb.Select(allFields...).
		From(taskTable).
		Where(sq.Eq{"id": id}).
		ToSql()
	if err != nil {
		return nil, errors.Wrap(err, "[repository.GetTask] sql build error")
	}
	stmt, err := t.db.PreparexContext(ctx, q)
	if err != nil {
		return nil, errors.Wrap(err, "[repository.GetTask] prepare context error")
	}
	err = stmt.GetContext(ctx, &task, args...)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, ErrNoResult
	} else if err != nil {
		return nil, errors.Wrap(err, "[repository.GetTask] select error")
	}
	return &task, nil
}

func (t *taskQuery) CreateTask(ctx context.Context, task entities.Task) (*int64, error) {
	q, args, err := t.qb.Insert(taskTable).
		Columns("title", "text", "completed", "counter_value",
			"counter_max_value", "counter_scale", "counter_exist").
		Values(task.Title, task.Text, task.Completed, task.CounterValue,
			task.CounterMaxValue, task.CounterScale, task.CounterExist).
		Suffix("returning id").
		ToSql()
	if err != nil {
		return nil, errors.Wrap(err, "[repository.CreateTask] sql build error")
	}
	stmt, err := t.db.PreparexContext(ctx, q)
	if err != nil {
		return nil, errors.Wrap(err, "[repository.CreateTask] prepare context error")
	}
	var id int64
	err = stmt.GetContext(ctx, &id, args...)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, ErrNoResult
	} else if err != nil {
		return nil, errors.Wrap(err, "[repository.CreateTask] insert or deserialize error")
	}
	return &id, nil
}

func (t *taskQuery) UpdateTask(ctx context.Context, task entities.Task) (*entities.Task, error) {
	//TODO implement me
	panic("implement me")
}

func (t *taskQuery) DeleteTask(ctx context.Context, id int64) error {
	//TODO implement me
	panic("implement me")
}
