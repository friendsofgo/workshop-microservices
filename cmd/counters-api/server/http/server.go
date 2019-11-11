package http

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

type Server struct {
	host string
	port uint

	srv *http.Server
}

// NewServer return a new HTTP server
func NewServer(host string, port uint) *Server {
	s := &Server{
		host: host,
		port: port,
	}

	router := mux.NewRouter()
	router.HandleFunc("/health", s.healthHandler)

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

type healthResponse struct {
	Kind    string `json:"kind"`
	Message string `json:"message"`
}

func (s Server) healthHandler(w http.ResponseWriter, r *http.Request) {
	response := map[string]healthResponse{
		"data": healthResponse{Kind: "health", Message: "everything is fine"},
	}

	b, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "error proccessing the response", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write(b)
}
