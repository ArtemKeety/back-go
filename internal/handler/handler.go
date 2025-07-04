package handler

import (
	"github.com/ArtemKeety/back-go.git/internal/service"
	"github.com/gorilla/mux"
)

type Handler struct {
	service *service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) InitRouter() *mux.Router {
	router := mux.NewRouter()
	router.Use(h.getLogger)

	auth := router.PathPrefix("/auth").Subrouter()
	{
		auth.HandleFunc("/sign-up", h.singUp).Methods("POST")
		auth.HandleFunc("/sign-in", h.singIn).Methods("POST")
		auth.HandleFunc("/test", h.hand_test).Methods("GET")
	}

	return router
}
