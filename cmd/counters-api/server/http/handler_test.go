package http

import (
	"context"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	creatingmock "github.com/friendsofgo/workshop-microservices/test/creating"
)

func TestServer_healthHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/health", nil)
	if err != nil {
		t.Fatalf("could not created request: %v", err)
	}

	rec := httptest.NewRecorder()
	srv := NewServer(context.Background(), "", 0, &creatingmock.Service{}, &log.Logger{})
	srv.healthHandler(context.Background()).ServeHTTP(rec, req)

	res := rec.Result()
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		t.Errorf("GET /health want %d, got: %d", http.StatusOK, res.StatusCode)
	}

	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Fatalf("GET /health has a unreadble response; %v", err)
	}

	expectedResponse := `{"data":{"kind":"health","message":"everything is fine"}}`
	strResponse := string(b)
	if strResponse != expectedResponse {
		t.Errorf("GET /health want response %s, got: %s", expectedResponse, strResponse)
	}
}

func TestServer_createHandler(t *testing.T) {
	t.Fatal("To be implemented...")
}
