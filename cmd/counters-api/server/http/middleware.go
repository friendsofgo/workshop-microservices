package http

import (
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func (s *Server) loggerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {

			if currentRoute := mux.CurrentRoute(r); currentRoute != nil {
				if routePath, err := currentRoute.GetPathTemplate(); err == nil {
					s.logger.Printf("%s: %s\n", r.Method, routePath)
				}
			}
			next.ServeHTTP(w, r)
		},
	)
}

func (s *Server) requestTimeMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()
			next.ServeHTTP(w, r)

			rt := time.Since(start)
			s.logger.Printf("Time request: %fs\n", rt.Seconds())
		},
	)
}
