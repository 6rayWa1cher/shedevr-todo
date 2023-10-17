package main

import (
	"fmt"
	"github.com/go-faster/errors"
	"github.com/go-playground/validator/v10"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"go.uber.org/zap"
	"os"
	"strconv"
)

type DbConfig struct {
	Host     string `mapstructure:"DB_HOST" validate:"required"`
	Port     int    `mapstructure:"DB_PORT" validate:"required,gt=0"`
	Name     string `mapstructure:"DB_NAME" validate:"required"`
	Username string `mapstructure:"DB_USERNAME" validate:"required"`
	Password string `mapstructure:"DB_PASSWORD" validate:"required"`
}

func readConfig() (*DbConfig, error) {
	port, err := strconv.Atoi(os.Getenv("DB_PORT"))
	if err != nil {
		return nil, err
	}
	config := DbConfig{
		Host:     os.Getenv("DB_HOST"),
		Port:     port,
		Name:     os.Getenv("DB_NAME"),
		Username: os.Getenv("DB_USERNAME"),
		Password: os.Getenv("DB_PASSWORD"),
	}
	v := validator.New()
	if err := v.Struct(&config); err != nil {
		return nil, err
	}
	return &config, nil
}

func migrateDb(source string, config *DbConfig) error {
	databaseUrl := fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s?sslmode=disable",
		config.Username, config.Password, config.Host, config.Port, config.Name,
	)
	zap.L().Info("Connecting to the database", zap.String("databaseUrl", databaseUrl))
	m, err := migrate.New(
		source,
		databaseUrl)
	if err != nil {
		return errors.Wrap(err, "creation of a migrator is failed")
	}
	defer m.Close()
	zap.L().Info("Connected to the database", zap.String("databaseUrl", databaseUrl))

	version, dirty, err := m.Version()
	zap.L().Info("Starting database migration",
		zap.Uint("version", version), zap.Bool("dirty", dirty), zap.Error(err))
	if err := m.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return errors.Wrap(err, "migration of the database is failed")
	}
	version, dirty, err = m.Version()
	zap.L().Info("Database migration completed successfully",
		zap.Uint("version", version), zap.Bool("dirty", dirty), zap.Error(err))

	return nil
}

func initLogger() {
	logger := zap.Must(zap.NewProduction())
	if os.Getenv("APP_ENV") == "development" {
		logger = zap.Must(zap.NewDevelopment())
	}
	zap.ReplaceGlobals(logger)
}

func main() {
	initLogger()

	zap.L().Info("db-migrate is booting")

	config, err := readConfig()
	if err != nil {
		zap.L().Fatal("fatal error on config read", zap.Error(err))
	}

	zap.L().Info("config ready", zap.Any("config", config))

	if err := migrateDb("file://db/", config); err != nil {
		zap.L().Fatal("fatal error on migrating", zap.Error(err))
	}
}
