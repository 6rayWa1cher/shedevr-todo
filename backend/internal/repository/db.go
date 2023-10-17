package repository

import (
	"database/sql"
	"fmt"
)

type DbConfig struct {
	Host     string `mapstructure:"DB_HOST" validate:"required"`
	Port     int    `mapstructure:"DB_PORT" validate:"required,gt=0"`
	Name     string `mapstructure:"DB_NAME" validate:"required"`
	Username string `mapstructure:"DB_USERNAME" validate:"required"`
	Password string `mapstructure:"DB_PASSWORD" validate:"required"`
}

func NewDb(c *DbConfig) (*sql.DB, error) {
	// Starting a database
	psqlInfo := fmt.Sprintf(
		"host=%s port=%d user=%s dbname=%s password=%s",
		c.Host, c.Port, c.Username, c.Name, c.Password,
	)
	var err error
	Db, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, err
	}
	return Db, nil
}
