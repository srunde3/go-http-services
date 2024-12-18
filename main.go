package main

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
)


type server struct {
	router chi.Router
}


func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}


func newServer() *server {
	s := &server{router: chi.NewRouter()}
	s.routes()
	return s
}


func main() {
	s := newServer()
	http.ListenAndServe("localhost:8000", s)
}

func HandleHealthCheck(w http.ResponseWriter, r *http.Request) {
	response := map[string]string{"status": "ok"}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	data, _ := json.Marshal(response)
	w.Write(data)
}


func HandleHiMom(w http.ResponseWriter, r *http.Request) {
	response := map[string]string{"message": "Hi, Mom!"}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	data, _ := json.Marshal(response)
	w.Write(data)
}