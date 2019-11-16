package http

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"

	"github.com/friendsofgo/workshop-microservices/internal/creating"
	"github.com/friendsofgo/workshop-microservices/internal/fetching"
)

type Server struct {
	host string
	port uint
	srv  *http.Server

	creating creating.Service
	fetching fetching.Service
	logger   *log.Logger
}

// NewServer return a new HTTP server
func NewServer(ctx context.Context, host string, port uint, c creating.Service, f fetching.Service, logger *log.Logger) *Server {
	s := &Server{
		host:   host,
		port:   port,
		logger: logger,

		fetching: f,
		creating: c,
	}

	router := mux.NewRouter()
	router.Use(s.loggerMiddleware, s.requestTimeMiddleware)

	router.HandleFunc("/health", s.healthHandler(ctx)).Methods(http.MethodGet)
	router.HandleFunc("/counters", s.createCounterHandler(ctx)).Methods(http.MethodPost)
	router.HandleFunc("/counters/belongs-to/{belongs_to:[a-zA-Z0-9]+}", s.fetchAllCountersHandler(ctx)).Methods(http.MethodGet)

	s.srv = &http.Server{
		Handler:      router,
		Addr:         fmt.Sprintf("%s:%d", host, port),
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	return s
}

// Serve execute the server on the host and port indicated previously
func (s Server) Serve() error {
	log.Println("The server is on tap now:", s.srv.Addr)
	return s.srv.ListenAndServe()
}
