package app

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/canyouhearthemusic/kamigr/config"
	"github.com/canyouhearthemusic/kamigr/db/sqlc"
	"github.com/canyouhearthemusic/kamigr/internal/handler"
	"github.com/canyouhearthemusic/kamigr/internal/service/reservation"
	"github.com/canyouhearthemusic/kamigr/pkg/server"
	"github.com/jackc/pgx/v5"
)

func Run(ctx context.Context) {
	cfg, err := config.MustLoad(ctx)
	if err != nil {
		log.Fatalf("failed to load config: %s", err)
		return
	}

	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		cfg.Database.User, cfg.Database.Password, cfg.Database.Host, cfg.Database.Port, cfg.Database.Name)

	conn, err := pgx.Connect(ctx, dsn)
	if err != nil {
		log.Fatalf("failed to establish connection to db: %s\n", err)
	}
	defer conn.Close(ctx)

	DAO := sqlc.New(conn)

	reservationService := reservation.New(DAO)

	handlers := handler.New(handler.Dependencies{
		ReservationService: reservationService,
	}, handler.WithHTTPHandler())

	server, err := server.New(server.WithHTTPServer(handlers.Mux, cfg.App.Port))
	if err != nil {
		return
	}

	if err := server.Start(); err != nil {
		return
	}

	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)

	<-shutdown

	ctx2, cancel := context.WithTimeout(ctx, 15*time.Second)
	defer cancel()

	if err := server.Stop(ctx2); err != nil {
		return
	}
}
