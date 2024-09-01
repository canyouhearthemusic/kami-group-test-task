package app

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/canyouhearthemusic/kamigr/internal/handler"
	"github.com/canyouhearthemusic/kamigr/pkg/server"
)

func Run(ctx context.Context) {
	// cfg, err := config.MustLoad(ctx)
	// if err != nil {
	// 	return
	// }
	handlers := handler.New(handler.Dependencies{}, handler.WithHTTPHandler())

	server, err := server.New(server.WithHTTPServer(handlers.Mux, "8081"))
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
