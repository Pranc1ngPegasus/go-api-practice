package router

import (
	"net/http"
	"time"

	"github.com/Pranc1ngPegasus/go-api-practice/internal/adapter/infrastructure"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func NewRouter(connector infrastructure.RDBConnector) http.Handler {
	router := chi.NewRouter()

	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(middleware.Timeout(60 * time.Second))

	router.Mount("/healthcheck", NewHealthcheck())
	router.Mount("/users", NewUser(connector))

	return router
}
