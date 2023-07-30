package server

import (
	"fmt"
	"net/http"
	"time"
)

type Server struct {
	srv http.Server
}

func (s *Server) Start(handlers http.Handler) error {
	s.srv = http.Server{
		Addr:         ":8080",
		Handler:      handlers,
		WriteTimeout: 10 * time.Second,
		ReadTimeout:  10 * time.Second,
	}
	fmt.Printf("Listen the server http://localhost:%d\n", 8080)
	return s.srv.ListenAndServe()
}
