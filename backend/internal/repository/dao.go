package repository

import (
	"database/sql"
	"github.com/Masterminds/squirrel"
)

type Dao interface {
	NewTaskQuery() TaskQuery
}

type dao struct {
	db *sql.DB
}

var Db *sql.DB

func pgQbFactory(db *sql.DB) func() squirrel.StatementBuilderType {
	return func() squirrel.StatementBuilderType {
		return squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar).RunWith(db)
	}
}

func NewDao(db *sql.DB) Dao {
	return &dao{
		db: db,
	}
}

func (d *dao) NewTaskQuery() TaskQuery {
	return &taskQuery{
		qb: pgQbFactory(d.db),
	}
}
