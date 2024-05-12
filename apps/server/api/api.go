package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jaka-k/apps/server/ticket-broker/rabbitmq"
	"github.com/rs/cors"
)

type OrderRequest struct {
	ID          string  `json:"id"`
	CountryCode string  `json:"countryCode"`
	OrderAmount float32 `json:"orderAmount"`
}

type APIServer struct {
	listenAddr string
	service    rabbitmq.Service
}

func NewAPIServer(listenAddr string, service rabbitmq.Service) *APIServer {
	return &APIServer{
		listenAddr: listenAddr,
		service:    service,
	}
}

func (s *APIServer) Run() {
	router := mux.NewRouter()
	router.HandleFunc("/buy", makeHTTPHandleFunc(s.handleBuyCall)).Methods("POST")
	router.HandleFunc("/confirm", makeHTTPHandleFunc(s.handleConfirmCall)).Methods("POST")

	corsMiddleware := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
	})

	handler := corsMiddleware.Handler(router)
	log.Println("JSON API server running on", s.listenAddr)
	log.Fatal(http.ListenAndServe(s.listenAddr, handler))
}

func (s *APIServer) handleBuyCall(w http.ResponseWriter, r *http.Request) error {
	var req OrderRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return fmt.Errorf("error decoding request body: %v", err)
	}

	queue, ok := AllowedCountries[req.CountryCode]
	if !ok {
		http.Error(w, "Forbidden country code", http.StatusForbidden)
		return fmt.Errorf("forbidden country code: %s", req.CountryCode)
	}

	err = s.service.PublishMessage(queue, req.ID)
	if err != nil {
		http.Error(w, "Failed to publish message", http.StatusInternalServerError)
		return fmt.Errorf("failed to publish message: %v", err)
	}

	log.Printf("Received a buy request for ID: %s from country: %s", req.ID, req.CountryCode)
	return WriteJSON(w, http.StatusOK, "Buy request received")
}

func (s *APIServer) handleConfirmCall(w http.ResponseWriter, r *http.Request) error {
	id := mux.Vars(r)["id"]
	log.Printf("Received a confirmation for ID: %s", id)
	return WriteJSON(w, http.StatusOK, "Confirmation received")
}

func WriteJSON(w http.ResponseWriter, status int, v any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(v)
}

type apiFunc func(http.ResponseWriter, *http.Request) error

func makeHTTPHandleFunc(f apiFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {
			WriteJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		}
	}
}

var AllowedCountries = map[string]string{
	"fr": "order.france",
	"de": "order.germany",
	"gb": "order.greatbritain",
	"es": "order.spain",
}
