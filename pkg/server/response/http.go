package response

import (
	"net/http"

	"github.com/go-chi/render"
)

type Response struct {
	Message string `json:"message,omitempty"`
	Data    any    `json:"data,omitempty"`
}

func Render(w http.ResponseWriter, r *http.Request, statusCode int, msg string, data any) {
	render.Status(r, statusCode)

	v := Response{
		Message: msg,
		Data:    data,
	}
	render.JSON(w, r, v)
}
