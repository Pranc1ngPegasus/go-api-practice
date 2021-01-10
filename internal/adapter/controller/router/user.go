package router

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
)

func NewUser() http.Handler {
	router := chi.NewRouter()

	router.Post("/", create)

	return router
}

type (
	CreateUserResponse struct {
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
		Email     string `json:"email"`
	}
)

func create(w http.ResponseWriter, r *http.Request) {
	render.Status(r, http.StatusOK)
	render.JSON(w, r, CreateUserResponse{
		FirstName: "dummy",
		LastName:  "dummy",
		Email:     "dummy",
	})
}
