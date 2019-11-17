package http

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/friendsofgo/workshop-microservices/internal/creating"
)

type Server struct {
	host string
	port uint
	srv  *http.ServeMux

	creating creating.Service
	logger   *log.Logger
}

// NewServer return a new HTTP server
func NewServer(ctx context.Context, host string, port uint, c creating.Service, logger *log.Logger) *Server {
	s := &Server{
		host:   host,
		port:   port,
		logger: logger,
		srv:    http.NewServeMux(),

		creating: c,
	}
	s.srv.HandleFunc("/health", s.healthHandler(ctx))
	s.srv.HandleFunc("/counters", s.createCounterHandler(ctx))
	return s
}

// Serve execute the server on the host and port indicated previously
func (s Server) Serve() error {
	httpAddr := fmt.Sprintf("%s:%d", s.host, s.port)
	log.Println("The server is on tap now:", httpAddr)

	return http.ListenAndServe(httpAddr, s.srv)
}
