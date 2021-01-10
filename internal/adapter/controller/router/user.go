package router

import (
	"net/http"
	"strconv"

	"github.com/Pranc1ngPegasus/go-api-practice/internal/adapter/infrastructure"
	"github.com/Pranc1ngPegasus/go-api-practice/internal/usecase"
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
)

type (
	UserHandler http.Handler

	userHandler struct {
		router      *chi.Mux
		connector   infrastructure.RDBConnector
		userUsecase usecase.GetUser
	}
)

func NewUser(
	connector infrastructure.RDBConnector,
) UserHandler {
	router := chi.NewRouter()

	handler := &userHandler{
		router:      router,
		connector:   connector,
		userUsecase: usecase.NewGetUser(connector),
	}

	handler.router.Get("/{userID}", handler.getUser)
	handler.router.Post("/", handler.create)

	return handler
}

func (h *userHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.router.ServeHTTP(w, r)
}

type (
	UserResponse struct {
		ID        int64  `json:"id"`
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
		Email     string `json:"email"`
	}

	ErrorResponse struct {
		Message string `json:"message"`
	}
)

func (h *userHandler) getUser(w http.ResponseWriter, r *http.Request) {
	pathParamUserID := chi.URLParam(r, "userID")
	userID, err := strconv.ParseInt(pathParamUserID, 10, 32)
	if err != nil {
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, ErrorResponse{
			Message: err.Error(),
		})
	}

	user, err := h.userUsecase.Execute(
		usecase.GetUserInput{
			ID: userID,
		},
	)
	if err != nil {
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, ErrorResponse{
			Message: err.Error(),
		})
	}

	render.Status(r, http.StatusOK)
	render.JSON(w, r, UserResponse{
		ID:        user.ID(),
		FirstName: user.FirstName(),
		LastName:  user.LastName(),
		Email:     user.Email(),
	})
}

func (h *userHandler) create(w http.ResponseWriter, r *http.Request) {
	render.Status(r, http.StatusOK)
	render.JSON(w, r, UserResponse{
		ID:        999,
		FirstName: "dummy",
		LastName:  "dummy",
		Email:     "dummy",
	})
}
