package http

import (
	"context"
	"encoding/json"
	"net/http"
)

type healthResponse struct {
	Kind    string `json:"kind"`
	Message string `json:"message"`
}

func (s Server) healthHandler(ctx context.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
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
}

type createRequest struct {
	Name      string `json:"name"`
	BelongsTo string `json:"belongs_to"`
}

func (s Server) createCounterHandler(ctx context.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req createRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "the body can't be parsed, check that is a valid json", http.StatusBadRequest)
			return
		}

		if err := s.creating.CreateCounter(ctx, req.Name, req.BelongsTo); err != nil {
			http.Error(w, "some error occurred when creating process was executed", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
	}
}
