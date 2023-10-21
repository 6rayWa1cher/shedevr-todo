package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type DbConfig struct {
	Host     string
	Port     int
	Name     string
	Username string
	Password string
}

func NewDb(c *DbConfig) (*sqlx.DB, error) {
	// Starting a database
	psqlInfo := fmt.Sprintf(
		"host=%s port=%d user=%s dbname=%s password=%s sslmode=disable",
		c.Host, c.Port, c.Username, c.Name, c.Password,
	)
	var err error
	Db, err = sqlx.Open("postgres", psqlInfo)
	if err != nil {
		return nil, err
	}
	return Db, nil
}

//func txFromCtx(db sqlx.DB, ctx context.Context) *sqlx.Tx {
//	tx := ctx.Value("tx").(*sqlx.Tx)
//	if tx != nil {
//		return tx
//	}
//	tx, err := db.BeginTx()
//}
