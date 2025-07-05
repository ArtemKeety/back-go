package handler

import (
	"context"
	"encoding/base64"
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

	res, err := h.service.Login(ctx, r.RemoteAddr, model.UserRequest{UserData: userRequest})
	if err != nil {
		sendError(w, http.StatusBadRequest, err.Error())
		return
	}

	sendOk(w, res)
}

func (h *Handler) Change(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 10*time.Second)
	defer cancel()

	var t model.RequestToken
	if err := json.NewDecoder(r.Body).Decode(&t); err != nil {
		sendError(w, http.StatusInternalServerError, err.Error())
		return
	}

	tokenDecoder, err := base64.StdEncoding.DecodeString(t.Token)
	if err != nil {
		sendError(w, http.StatusInternalServerError, err.Error())
		return
	}

	data, err := h.service.ChangeToken(ctx, r.RemoteAddr, string(tokenDecoder))
	if err != nil {
		sendError(w, http.StatusBadRequest, err.Error())
		return
	}

	sendOk(w, data)

}
