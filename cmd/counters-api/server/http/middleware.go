package http

import (
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

type loggingResponseWriter struct {
	http.ResponseWriter
	statusCode int
}

func (lrw *loggingResponseWriter) WriteHeader(code int) {
	lrw.statusCode = code
	lrw.ResponseWriter.WriteHeader(code)
}

func (s *Server) loggerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			lrw := &loggingResponseWriter{w, http.StatusOK}
			next.ServeHTTP(lrw, r)

			path, _ := mux.CurrentRoute(r).GetPathTemplate()
			s.logger.Printf("%d|%s: %s\n", lrw.statusCode, r.Method, path)
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
