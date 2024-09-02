package config

import (
	"context"

	"github.com/joho/godotenv"
	"github.com/sethvargo/go-envconfig"
)

type App struct {
	Host string `env:"APP_HOST"`
	Port string `env:"APP_PORT"`
}

type Database struct {
	User     string `env:"DATABASE_USER"`
	Password string `env:"DATABASE_PASSWORD"`
	Host     string `env:"DATABASE_HOST"`
	Port     string `env:"DATABASE_PORT"`
	Name     string `env:"DATABASE_NAME"`
}

type Config struct {
	App      App
	Database Database
}

func MustLoad(ctx context.Context) (*Config, error) {
	cfg := new(Config)

	if err := godotenv.Load(); err != nil {
		return nil, err
	}

	if err := envconfig.Process(ctx, cfg); err != nil {
		return nil, err
	}

	return cfg, nil
}
