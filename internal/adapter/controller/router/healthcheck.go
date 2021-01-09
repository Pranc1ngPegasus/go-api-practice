package router

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
)

func NewHealthcheck() http.Handler {
	router := chi.NewRouter()

	router.Get("/", get)

	return router
}

type (
	HealthcheckResponse struct {
		Message string `json:"message"`
	}
)

func get(w http.ResponseWriter, r *http.Request) {
	render.Status(r, http.StatusOK)
	render.JSON(w, r, HealthcheckResponse{
		Message: "ok",
	})
}
