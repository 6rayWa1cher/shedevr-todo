package main

import (
	"context"
	"fmt"
	"github.com/6rayWa1cher/shedevr-todo/backend/internal/app"
	"github.com/6rayWa1cher/shedevr-todo/backend/internal/server"
	"github.com/6rayWa1cher/shedevr-todo/backend/pkg/oas"
	"github.com/go-faster/errors"
	"github.com/go-playground/validator/v10"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"golang.org/x/sync/errgroup"
	"net/http"
)

type Config struct {
	ApiPort        int32 `mapstructure:"API_PORT" validate:"required,gt=0"`
	MetricsEnabled bool  `mapstructure:"METRICS_ENABLED"`
	MetricsPort    int32 `mapstructure:"METRICS_PORT" validate:"required_with=MetricsEnabled&gt=0"`
}

func NewConfig() Config {
	return Config{
		ApiPort:        8080,
		MetricsEnabled: false,
		MetricsPort:    0,
	}
}

func initConfig() (*Config, error) {
	viper.AddConfigPath("config")
	viper.SetConfigName("config.env")
	viper.SetConfigType("env")
	viper.SetEnvPrefix("SHEDEVR_BACKEND")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	config := NewConfig()
	if err := viper.Unmarshal(&config); err != nil {
		return nil, err
	}
	v := validator.New()
	if err := v.Struct(&config); err != nil {
		return nil, err
	}
	return &config, nil
}

func initMetrics(addr string, lg *zap.Logger) (*server.Metrics, []oas.ServerOption, error) {
	m, err := server.NewMetrics(lg, server.Config{
		Addr: addr,
		Name: "api",
	})
	if err != nil {
		return nil, make([]oas.ServerOption, 0), err
	}
	options := []oas.ServerOption{
		oas.WithTracerProvider(m.TracerProvider()),
		oas.WithMeterProvider(m.MeterProvider()),
	}
	return m, options, nil
}

func main() {
	fmt.Println("Hello Docker!")
	server.Run(func(ctx context.Context, lg *zap.Logger) error {
		appConfig, err := initConfig()
		if err != nil {
			return errors.Wrap(err, "config.env file")
		}

		apiAddr := fmt.Sprintf("127.0.0.1:%d", appConfig.ApiPort)
		metricsAddr := fmt.Sprintf("127.0.0.1:%d", appConfig.MetricsPort)

		lg.Info("Initializing",
			zap.String("http.addr", apiAddr),
			zap.String("metrics.addr", metricsAddr),
		)

		options := make([]oas.ServerOption, 0)
		var metrics *server.Metrics
		if appConfig.MetricsEnabled {
			m, opts, err := initMetrics(metricsAddr, lg)
			if err != nil {
				return errors.Wrap(err, "metrics init")
			}
			options = append(options, opts...)
			metrics = m
		} else {
			lg.Info("Metrics are disabled")
		}

		handler := app.NewService()

		oasServer, err := oas.NewServer(handler, options...)
		if err != nil {
			return errors.Wrap(err, "server init")
		}
		httpServer := http.Server{
			Addr:    apiAddr,
			Handler: oasServer,
		}

		g, ctx := errgroup.WithContext(ctx)
		if metrics != nil {
			g.Go(func() error {
				return metrics.Run(ctx)
			})
		}
		g.Go(func() error {
			<-ctx.Done()
			return httpServer.Shutdown(ctx)
		})
		g.Go(func() error {
			defer lg.Info("Server stopped")
			lg.Info("Starting serve")
			if err := httpServer.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
				return errors.Wrap(err, "http")
			}
			return nil
		})

		return g.Wait()
	})
}
