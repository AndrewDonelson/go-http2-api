package api

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHeartbeatHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/heartbeat", bytes.NewBuffer(nil))
	if err != nil {
		t.Fatal(err)
	}

	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(HeartbeatHandler)

	handler.ServeHTTP(recorder, req)

	if status := recorder.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	expected := "Service is alive!"
	if recorder.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", recorder.Body.String(), expected)
	}
}
