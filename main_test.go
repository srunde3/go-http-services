package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHealthEndpoint(t *testing.T) {
	// The httptest.NewRequest (as opposed to http.NewRequest) does *not* return
	// an error. This is handy for tests, where I don't care about errors creating reqs.
	r := httptest.NewRequest(http.MethodGet, "/health", nil)
	w := httptest.NewRecorder()

	handler := http.HandlerFunc(handleHealthCheck)
	handler.ServeHTTP(w, r)

	if status := w.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v, want %v", status, http.StatusOK)
	}

	expected := `{"status":"ok"}`

	if w.Body.String() != expected {
		t.Errorf("Handler returned unexpected body: got %v, want %v", w.Body.String(), expected)
	}
}

func TestHiMomEndpoint(t *testing.T) {
	r := httptest.NewRequest(http.MethodGet, "/hi-mom", nil)
	w := httptest.NewRecorder()

	handler := http.HandlerFunc(handleHiMom)
	handler.ServeHTTP(w, r)

	if status := w.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v, want %v", status, http.StatusOK)
	}

	expected := `{"message":"Hi, Mom!"}`

	if w.Body.String() != expected {
		t.Errorf("Handler returned unexpected body: got %v, want %v", w.Body.String(), expected)
	}
}

func TestHealthRouter(t *testing.T) {
	r := httptest.NewRequest(http.MethodGet, "/health", nil)
	w := httptest.NewRecorder()

	s := newServer()

	s.ServeHTTP(w, r)

	expected := `{"status":"ok"}`

	if w.Body.String() != expected {
		t.Errorf("Handler returned unexpected body: got %v, want %v", w.Body.String(), expected)
	}
}

func TestHiMomRouter(t *testing.T) {
	r := httptest.NewRequest(http.MethodGet, "/hi-mom", nil)
	w := httptest.NewRecorder()

	s := newServer()

	s.ServeHTTP(w, r)

	expected := `{"message":"Hi, Mom!"}`

	if w.Body.String() != expected {
		t.Errorf("Handler returned unexpected body: got %v, want %v", w.Body.String(), expected)
	}
}

func TestEchoUrlParameterRouter(t *testing.T) {
	r := httptest.NewRequest(http.MethodGet, "/echo/12345", nil)
	w := httptest.NewRecorder()

	s := newServer()

	s.ServeHTTP(w, r)

	if status := w.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v, want %v", status, http.StatusOK)
	}

	expected := `{"parameter":"12345"}`

	if w.Body.String() != expected {
		t.Errorf("Handler returned unexpected body: got %v, want %v", w.Body.String(), expected)
	}
}
