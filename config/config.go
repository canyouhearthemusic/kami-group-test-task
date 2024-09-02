package config

import (
	"context"
	"log"
	"sync"

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

var once sync.Once
var cfg *Config

func MustLoad(ctx context.Context) (*Config, error) {
	once.Do(func() {
		cfg = new(Config)

		if err := godotenv.Load(); err != nil {
			log.Fatalf("Error loading .env file: %v", err)
			return
		}

		if err := envconfig.Process(ctx, cfg); err != nil {
			log.Fatalf("Error processing environment variables: %v", err)
			return
		}
	})

	return cfg, nil
}
