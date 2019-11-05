package http

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Server struct {
	host string
	port uint

	srv *http.ServeMux
}

// NewServer return a new HTTP server
func NewServer(host string, port uint) *Server {
	s := &Server{
		host: host,
		port: port,

		srv: http.NewServeMux(),
	}

	s.srv.HandleFunc("/health", s.healthHandler)
	return s
}

// Serve execute the server on the host and port indicated previously
func (s Server) Serve() {
	httpAddr := fmt.Sprintf("%s:%d", s.host, s.port)
	log.Println("The server is on tap now:", httpAddr)

	_ = http.ListenAndServe(httpAddr, s.srv)
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
