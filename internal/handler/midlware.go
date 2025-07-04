package handler

import (
	"github.com/sirupsen/logrus"
	"net/http"
)

func (h *Handler) getLogger(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			logrus.Infof("method: %s; point: %s; address: %s", r.Method, r.URL.Path, r.RemoteAddr) //r.URL.Query()
			next.ServeHTTP(w, r)
		},
	)
}
