package repository

import (
	"context"
	"database/sql"
	"github.com/6rayWa1cher/shedevr-todo/backend/internal/entities"
	sq "github.com/Masterminds/squirrel"
	"github.com/go-faster/errors"
	"github.com/jmoiron/sqlx"
	"strings"
)

type TaskQuery interface {
	GetTasks(ctx context.Context) ([]entities.Task, error)
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

var allFieldsExceptId = []string{
	"title", "text", "completed", "counter_value", "counter_max_value", "counter_scale", "counter_exist",
}

var allFields = append([]string{"id"}, allFieldsExceptId...)

const taskTable = "todo.task"

func (t *taskQuery) GetTasks(ctx context.Context) ([]entities.Task, error) {
	tasks := make([]entities.Task, 10)
	q, args, err := t.qb.Select(allFields...).
		From(taskTable).
		ToSql()
	if err != nil {
		return nil, errors.Wrap(err, "[repository.GetTasks] sql build error")
	}
	stmt, err := t.db.PreparexContext(ctx, q)
	if err != nil {
		return nil, errors.Wrap(err, "[repository.GetTasks] prepare context error")
	}
	err = stmt.SelectContext(ctx, &tasks, args...)
	if errors.Is(err, sql.ErrNoRows) {
		return tasks, nil
	} else if err != nil {
		return nil, errors.Wrap(err, "[repository.GetTasks] select error")
	}
	return tasks, nil
}

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
		Columns(allFieldsExceptId...).
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
	q, args, err := t.qb.Update(taskTable).
		Set("title", task.Title).
		Set("text", task.Text).
		Set("completed", task.Completed).
		Set("counter_value", task.CounterValue).
		Set("counter_max_value", task.CounterMaxValue).
		Set("counter_scale", task.CounterScale).
		Set("counter_exist", task.CounterExist).
		Where(sq.Eq{"id": task.ID}).
		Suffix("returning " + strings.Join(allFields, ",")).
		ToSql()
	if err != nil {
		return nil, errors.Wrap(err, "[repository.UpdateTask] sql build error")
	}
	stmt, err := t.db.PreparexContext(ctx, q)
	if err != nil {
		return nil, errors.Wrap(err, "[repository.UpdateTask] prepare context error")
	}
	var outputTask entities.Task
	err = stmt.GetContext(ctx, &outputTask, args...)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, ErrNoResult
	} else if err != nil {
		return nil, errors.Wrap(err, "[repository.UpdateTask] update or deserialize error")
	}
	return &outputTask, nil
}

func (t *taskQuery) DeleteTask(ctx context.Context, id int64) error {
	q, args, err := t.qb.Delete(taskTable).
		Where(sq.Eq{"id": id}).
		ToSql()
	if err != nil {
		return errors.Wrap(err, "[repository.DeleteTask] sql build error")
	}
	stmt, err := t.db.PreparexContext(ctx, q)
	if err != nil {
		return errors.Wrap(err, "[repository.DeleteTask] prepare context error")
	}
	_, err = stmt.ExecContext(ctx, args...)
	if err != nil {
		return errors.Wrap(err, "[repository.DeleteTask] delete error")
	}
	return nil
}
