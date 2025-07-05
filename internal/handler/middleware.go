package handler

import (
	"context"
	"github.com/ArtemKeety/back-go.git/pkg/token"
	"github.com/sirupsen/logrus"
	"net/http"
	"strings"
)

func (h *Handler) getLogger(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			logrus.Infof("method: %s; point: %s; address: %s", r.Method, r.URL.Path, r.RemoteAddr) //r.URL.Query()
			next.ServeHTTP(w, r)
		},
	)
}

func (h *Handler) userIdentity(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		auth := r.Header.Get("Authorization")
		if auth == "" {
			sendError(w, http.StatusUnauthorized, "Authorization header missing")
			return
		}

		t := strings.Split(auth, "Bearer ")[1]

		guid, err := token.ParseToken(t)
		if err != nil {
			sendError(w, http.StatusUnauthorized, "Invalid token")
			return
		}

		ctx := context.WithValue(r.Context(), "guid", guid)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
