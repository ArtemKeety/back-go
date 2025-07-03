package handler

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
	"net/http"
)

type CustomError struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}

func sendOk(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		sendError(w, http.StatusInternalServerError, err.Error())
	}
}

func sendError(w http.ResponseWriter, code int, message string) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(code)
	if err := json.NewEncoder(w).Encode(CustomError{Message: message, Code: code}); err != nil {
		logrus.Errorf("Faled to Encode JSON: %v", err)
	}
}
