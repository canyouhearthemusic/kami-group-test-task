package httphandler

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

type RoomHandler struct {
}

func NewRoomHandler() *RoomHandler {
	return &RoomHandler{}
}

func (h *RoomHandler) Routes() *chi.Mux {
	r := chi.NewRouter()

	return r
}

func (h *RoomHandler) create(w http.Response, r *http.Request) {
}

func (h *RoomHandler) get(w http.Response, r *http.Request) {

}

func (h *RoomHandler) update(w http.Response, r *http.Request) {

}

func (h *RoomHandler) delete(w http.Response, r *http.Request) {

}
