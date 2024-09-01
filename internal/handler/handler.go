package handler

import (
	"github.com/canyouhearthemusic/kamigr/internal/handler/httphandler"
	"github.com/canyouhearthemusic/kamigr/pkg/server/router"
	"github.com/go-chi/chi/v5"
)

type Dependencies struct {
}

type Handler struct {
	deps Dependencies
	Mux  *chi.Mux
}

type Configuration func(h *Handler) error

func New(deps Dependencies, cfgs ...Configuration) Handler {
	h := Handler{
		deps: deps,
	}

	for _, cfg := range cfgs {
		cfg(&h)
	}

	return h
}

func WithHTTPHandler() Configuration {
	return func(h *Handler) error {
		h.Mux = router.New()

		roomHandler := httphandler.NewRoomHandler()
		reservationHandler := httphandler.NewReservationHandler()

		h.Mux.Route("/api/v1", func(r chi.Router) {
			r.Mount("/rooms", roomHandler.Routes())

			r.Mount("/reservations", reservationHandler.Routes())
		})

		return nil
	}
}
