package repository

import (
	"errors"
	"github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
)

type Dao interface {
	NewTaskQuery() TaskQuery
}

type QuerySupplier interface {
}

type dao struct {
	db *sqlx.DB
}

var Db *sqlx.DB

func pgQb() squirrel.StatementBuilderType {
	return squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar)
}

func NewDao(db *sqlx.DB) Dao {
	return &dao{
		db: db,
	}
}

func (d *dao) NewTaskQuery() TaskQuery {
	return &taskQuery{
		qb: pgQb(),
		db: d.db,
	}
}

var ErrNoResult = errors.New("no result found")
