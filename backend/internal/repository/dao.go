package repository

import (
	"database/sql"
	"fmt"
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

type DaoConfig struct {
	Host     string
	Port     int
	Name     string
	Username string
	Password string
}

func NewDb(c DaoConfig) (*sql.DB, error) {
	// Starting a database
	psqlInfo := fmt.Sprintf(
		"host=%s port=%d user=%s dbname=%s password=%s sslmode=prefer",
		c.Host, c.Port, c.Username, c.Name, c.Password,
	)
	var err error
	Db, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, err
	}
	return Db, nil
}

func (d *dao) NewTaskQuery() TaskQuery {
	return &taskQuery{
		qb: pgQbFactory(d.db),
	}
}
