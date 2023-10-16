package entities

type Task struct {
	ID              int64   `db:"id"`
	Title           string  `db:"title"`
	Text            string  `db:"text"`
	Completed       string  `db:"completed"`
	CounterValue    float64 `db:"counter_value"`
	CounterMaxValue float64 `db:"counter_max_value"`
	CounterScale    string  `db:"counter_scale"`
}
