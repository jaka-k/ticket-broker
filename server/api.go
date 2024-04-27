package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type APIServer struct {
	listenAddr string
}

func NewAPIServer(listenAddr string) *APIServer {
	return &APIServer{
		listenAddr: listenAddr,
	}

}

func (s *APIServer) Run() {
	router := mux.NewRouter()

	router.HandleFunc("/buy", makeHTTPHandleFunc(s.handleCall))
	router.HandleFunc("/buy/{id}", makeHTTPHandleFunc(s.handleCall))

	log.Println("JSON API server running on port" + s.listenAddr + "\nhttp://localhost" + s.listenAddr)

	err := http.ListenAndServe(s.listenAddr, router)
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

func (s *APIServer) handleCall(w http.ResponseWriter, r *http.Request) error {
	if r.Method == "GET" {
		return s.handleBuy(w, r)
	}
	if r.Method == "GET" {
		return s.handleConfirm(w, r)
	}

	return fmt.Errorf("method now allowed %v", r.Method)
}

func (s *APIServer) handleBuy(w http.ResponseWriter, r *http.Request) error {

	fmt.Println(r)
	return WriteJSON(w, http.StatusOK, "Some oneee")
}

func (s *APIServer) handleConfirm(w http.ResponseWriter, r *http.Request) error {
	id := mux.Vars(r)["id"]

	fmt.Println(id)

	return WriteJSON(w, http.StatusOK, "Some oneee")
}

func WriteJSON(w http.ResponseWriter, status int, v any) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)

	return json.NewEncoder(w).Encode(v)
}

type apiFunc func(http.ResponseWriter, *http.Request) error

type ApiError struct {
	Error string `json:"error"`
}

func makeHTTPHandleFunc(f apiFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {
			WriteJSON(w, http.StatusBadRequest, ApiError{Error: err.Error()})
			// handle the error
		}
	}

}
