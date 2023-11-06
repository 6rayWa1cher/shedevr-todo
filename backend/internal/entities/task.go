package entities

import "database/sql"

type Task struct {
	ID              int64           `db:"id"`
	Title           string          `db:"title"`
	Text            sql.NullString  `db:"text"`
	Completed       string          `db:"completed"`
	CounterExist    bool            `db:"counter_exist"`
	CounterValue    sql.NullFloat64 `db:"counter_value"`
	CounterMaxValue sql.NullFloat64 `db:"counter_max_value"`
	CounterScale    sql.NullString  `db:"counter_scale"`
	UserId          string          `db:"user_id"`
}
