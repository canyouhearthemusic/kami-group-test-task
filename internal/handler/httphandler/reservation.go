package httphandler

import (
	"encoding/json"
	"net/http"

	"github.com/canyouhearthemusic/kamigr/internal/dto"
	"github.com/canyouhearthemusic/kamigr/internal/service/reservation"
	"github.com/canyouhearthemusic/kamigr/pkg/server/response"
	"github.com/go-chi/chi/v5"
)

type ReservationHandler struct {
	service *reservation.Service
}

func NewReservationHandler(service *reservation.Service) *ReservationHandler {
	return &ReservationHandler{
		service: service,
	}
}

func (h *ReservationHandler) Routes() *chi.Mux {
	r := chi.NewRouter()

	r.Post("/", h.create)
	r.Get("/{room_id}", h.listByRoomID)

	return r
}

func (h *ReservationHandler) create(w http.ResponseWriter, r *http.Request) {
	req := new(dto.ReservationDTO)

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.Render(w, r, http.StatusBadRequest, err.Error(), req)
		return
	}

	if err := req.Validate(); err != nil {
		response.Render(w, r, http.StatusBadRequest, err.Error(), req)
		return
	}

	sqlcreq, err := req.ConvertToSQLC()
	if err != nil {
		response.Render(w, r, http.StatusBadRequest, err.Error(), req)
		return
	}

	reservation, err := h.service.CreateReservation(r.Context(), sqlcreq)
	if err != nil {
		response.Render(w, r, http.StatusConflict, err.Error(), sqlcreq)
		return
	}

	response.Render(w, r, http.StatusCreated, "reservation created", reservation)
}

func (h *ReservationHandler) listByRoomID(w http.ResponseWriter, r *http.Request) {
	roomID := chi.URLParam(r, "room_id")
	if roomID == "" {
		response.Render(w, r, http.StatusBadRequest, "room_id required", roomID)
		return
	}

	reservations, err := h.service.GetAllReservationsByRoomID(r.Context(), roomID)
	if err != nil {
		response.Render(w, r, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	response.Render(w, r, http.StatusOK, "", reservations)
}
