package back_go

import (
	"fmt"
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

	fmt.Println("server is start")

	return s.server.ListenAndServe()
}

func (s *Server) Close() error {
	return s.server.Close()
}
