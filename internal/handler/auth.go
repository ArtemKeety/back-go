package handler

import (
	"context"
	"encoding/json"
	"github.com/ArtemKeety/back-go.git/internal/model"
	"net/http"
	"time"
)

func (h *Handler) hand_test(w http.ResponseWriter, r *http.Request) {
	data := make(map[string]string)
	data["status"] = "ok"
	sendOk(w, data)
}

func (h *Handler) singUp(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 10*time.Second)
	defer cancel()

	var user model.UserRequest

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		sendError(w, http.StatusInternalServerError, err.Error())
	}

	id, err := h.service.CreateUser(ctx, user)
	if err != nil {
		sendError(w, http.StatusBadRequest, err.Error())
	}

	sendOk(w, map[string]interface{}{"id": id})
}
