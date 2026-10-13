package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHealth(t *testing.T) {
	response := httptest.NewRecorder()
	handler().ServeHTTP(response, httptest.NewRequest(http.MethodGet, "/health", nil))

	if response.Code != http.StatusOK {
		t.Fatalf("status = %d, want %d", response.Code, http.StatusOK)
	}
	if response.Body.String() != "ok\n" {
		t.Fatalf("body = %q, want %q", response.Body.String(), "ok\n")
	}
	if response.Header().Get("Content-Type") != "text/plain; charset=utf-8" {
		t.Fatalf("content type = %q", response.Header().Get("Content-Type"))
	}
}
