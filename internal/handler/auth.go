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
		return
	}

	guid, err := h.service.CreateUser(ctx, user)
	if err != nil {
		sendError(w, http.StatusBadRequest, err.Error())
		return
	}

	sendOk(w, map[string]interface{}{"guid": guid})
}

func (h *Handler) singIn(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 10*time.Second)
	defer cancel()

	var userRequest model.UserData

	if err := json.NewDecoder(r.Body).Decode(&userRequest); err != nil {
		sendError(w, http.StatusInternalServerError, err.Error())
		return
	}

	var user model.UserRequest
	user.UserData = userRequest

	res, err := h.service.Login(ctx, r.RemoteAddr, user)
	if err != nil {
		sendError(w, http.StatusBadRequest, err.Error())
		return
	}

	sendOk(w, res)
}
