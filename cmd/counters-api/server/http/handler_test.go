package http

import (
	"bytes"
	"context"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	creatingmock "github.com/friendsofgo/workshop-microservices/test/creating"
	"github.com/stretchr/testify/mock"
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
	body := []byte(`{
		"name": "test",
		"belongs_to": "01DSKGZBMHQTNTVPRST0ASXTC7"
	}`)

	req, err := http.NewRequest("POST", "/counters", bytes.NewBuffer(body))
	if err != nil {
		t.Fatalf("could not created request: %v", err)
	}
	testCtx := context.Background()

	rec := httptest.NewRecorder()
	service := &creatingmock.Service{}
	service.On("CreateCounter",
		testCtx,
		mock.AnythingOfType("string"),
		mock.AnythingOfType("string"),
	).Return(nil).Once()

	srv := NewServer(testCtx, "", 0, service, &log.Logger{})
	srv.createCounterHandler(testCtx).ServeHTTP(rec, req)

	res := rec.Result()
	defer res.Body.Close()

	if res.StatusCode != http.StatusCreated {
		t.Errorf("POST /counters want %d, got: %d", http.StatusCreated, res.StatusCode)
	}

}
