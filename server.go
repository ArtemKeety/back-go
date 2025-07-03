package back_go

import (
	"github.com/sirupsen/logrus"
	"net/http"
	"time"
)

type Server struct {
	server *http.Server
}

func (s *Server) Run(adr string, h http.Handler) error {
	s.server = &http.Server{
		Addr:           adr,
		Handler:        h,
		MaxHeaderBytes: 1 << 20,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
	}

	logrus.Infof("server listening at %s", s.server.Addr)
	return s.server.ListenAndServe()
}

func (s *Server) Close() error {
	return s.server.Close()
}
