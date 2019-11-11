package http

import (
	"net/http"

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
