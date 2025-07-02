package handler

import (
	_ "fmt"
	"github.com/sirupsen/logrus"
	"net/http"
)

func (h *Handler) getLogger(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			logrus.Infof("method: %s; point: %s; query: %s; address: %s", r.Method, r.URL.Path, r.RequestURI, r.RemoteAddr)
			next.ServeHTTP(w, r)
		},
	)
}
