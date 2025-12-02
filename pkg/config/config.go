package config

import (
	"github.com/kelseyhightower/envconfig"
	"time"
)

type Config struct {
	App struct {
		Name    string `env:"APP_NAME,required"`
		Env     string `env:"APP_ENV" default:"dev"`
		Version string `env:"APP_VERSION" default:"local"`
	}

	HTTP struct {
		Port            int           `env:"HTTP_PORT" default:"8080"`
		ReadTimeout     time.Duration `env:"HTTP_READ_TIMEOUT" default:"5s"`
		WriteTimeout    time.Duration `env:"HTTP_WRITE_TIMEOUT" default:"5s"`
		IdleTimeout     time.Duration `env:"HTTP_IDLE_TIMEOUT" default:"60s"`
		ShutdownTimeout time.Duration `env:"HTTP_SHUTDOWN_TIMEOUT" default:"5s"`
	}

	GRPC struct {
		Port            int           `env:"GRPC_PORT" default:"50051"`
		ShutdownTimeout time.Duration `env:"GRPC_SHUTDOWN_TIMEOUT" default:"5s"`
	}

	Log struct {
		Level  string `env:"LOG_LEVEL" default:"info"`
		Format string `env:"LOG_FORMAT" default:"json"`
	}

	Metrics struct {
		Port int `env:"METRICS_PORT" default:"9090"`
	}

	OTEL struct {
		Enabled      bool    `env:"OTEL_ENABLED" default:"false"`
		Endpoint     string  `env:"OTEL_EXPORTER_OTLP_ENDPOINT" default:""`
		SamplerRatio float64 `env:"OTEL_SAMPLER_RATIO" default:"1.0"`
	}
}

func Load() (*Config, error) {
	var cfg Config
	if err := envconfig.Process("", &cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}

func MustLoad() *Config {
	cfg, err := Load()
	if err != nil {
		panic(err)
	}
	return cfg
}
