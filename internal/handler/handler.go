package handler

import (
	"github.com/gorilla/mux"
)

type Handler struct {
	service interface{}
}

func NewHandler(service interface{}) *Handler {
	return &Handler{service: service}
}

func (h *Handler) InitRouter() *mux.Router {
	router := mux.NewRouter()
	router.Use(h.getLogger)

	auth := router.PathPrefix("/auth").Subrouter()
	{
		auth.HandleFunc("/sign-up", nil).Methods("POST")
		auth.HandleFunc("/sign-in", nil).Methods("POST")
		auth.HandleFunc("/test", h.hand_test).Methods("GET")
	}

	return router
}
