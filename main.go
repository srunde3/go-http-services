package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/pkg/errors"
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

	port := flag.String("port", "9100", "Local port to bind to")
	flag.Parse()

	if err := run(*port); err != nil {
		log.Fatalf("%v", err)
	}
}

func run(port string) error {
	s := &server{}
	router := chi.NewRouter()
	router.Use(middleware.Logger)
	s.router = router
	s.routes()

	bind := fmt.Sprintf("localhost:%v", port)
	if err := http.ListenAndServe(bind, s); err != nil {
		return errors.Wrap(err, "serve application")
	}

	return nil
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

func HandleEchoUrlParameter(w http.ResponseWriter, r *http.Request) {
	parameter := chi.URLParam(r, "parameter")
	response := map[string]string{"parameter": parameter}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	data, _ := json.Marshal(response)
	w.Write(data)
}
