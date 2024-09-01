package httphandler

import "github.com/go-chi/chi/v5"

type ReservationHandler struct {
}

func NewReservationHandler() *ReservationHandler {
	return &ReservationHandler{}
}

func (h *ReservationHandler) Routes() *chi.Mux {
	r := chi.NewRouter()

	return r
}
