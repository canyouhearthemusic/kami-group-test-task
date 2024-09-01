package config

import (
	"context"

	"github.com/joho/godotenv"
	"github.com/sethvargo/go-envconfig"
)

type Config struct {
}

func MustLoad(ctx context.Context) (*Config, error) {
	cfg := new(Config)

	if err := godotenv.Load(); err != nil {
		return nil, err
	}

	if err := envconfig.Process(ctx, &cfg); err != nil {
		return nil, err
	}

	return cfg, nil
}
