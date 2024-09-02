package handler

import (
	"github.com/canyouhearthemusic/kamigr/internal/handler/httphandler"
	"github.com/canyouhearthemusic/kamigr/internal/service/reservation"
	"github.com/canyouhearthemusic/kamigr/pkg/server/router"
	"github.com/go-chi/chi/v5"
)

type Dependencies struct {
	ReservationService *reservation.Service
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

		reservationHandler := httphandler.NewReservationHandler(h.deps.ReservationService)

		h.Mux.Route("/api/v1", func(r chi.Router) {
			r.Mount("/reservations", reservationHandler.Routes())
		})

		return nil
	}
}
