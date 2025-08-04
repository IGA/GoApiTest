package handler

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestPingHandler(t *testing.T) {
	req := httptest.NewRequest("GET", "/ping", nil)
	w := httptest.NewRecorder()

	PingHandler(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("Expected status 200, got %d", w.Code)
	}

	expected := `{"message":"pong"}`
	actual := w.Body.String()

	if actual != expected+"\n" {
		t.Errorf("Expected %s, got %s", expected, actual)
	}
}
