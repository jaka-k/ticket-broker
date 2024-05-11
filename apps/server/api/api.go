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
	publisher  *rabbitmq.RabbitMQPublisher
}

var AllowedCountries = map[string]string{
	"fr": "FranceQueue",
	"de": "GermanyQueue",
	"gb": "GreatBritainQueue",
	"es": "SpainQueue",
}

func NewAPIServer(listenAddr string, publisher *rabbitmq.RabbitMQPublisher) *APIServer {
	return &APIServer{
		listenAddr: listenAddr,
		publisher:  publisher,
	}

}

func (s *APIServer) Run() {
	router := mux.NewRouter()

	router.HandleFunc("/buy", makeHTTPHandleFunc(s.handleBuyCall))
	router.HandleFunc("/confirm", makeHTTPHandleFunc(s.handleConfirmCall))

	corsMiddleware := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
	})

	handler := corsMiddleware.Handler(router)
	log.Println(handler)

	log.Println("JSON API server running on port" + s.listenAddr + "\nhttp://localhost" + s.listenAddr)

	err := http.ListenAndServe(s.listenAddr, handler)
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

func (s *APIServer) handleBuyCall(w http.ResponseWriter, r *http.Request) error {
	if r.Method == "POST" {
		return s.handleBuy(w, r)
	}

	return fmt.Errorf("method now allowed %v", r.Method)
}

func (s *APIServer) handleConfirmCall(w http.ResponseWriter, r *http.Request) error {
	if r.Method == "POST" {
		return s.handleConfirm(w, r)
	}

	return fmt.Errorf("method now allowed %v", r.Method)
}

func (s *APIServer) handleBuy(w http.ResponseWriter, r *http.Request) error {
	var req OrderRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return fmt.Errorf("error decoding request body: %v", err)
	}

	queue, ok := AllowedCountries[req.CountryCode]
	if !ok {
		http.Error(w, "Forbidden country code", http.StatusForbidden)
		return fmt.Errorf("forbidden country code %s", req.CountryCode)
	}

	s.publisher.PublishMessage(queue, []byte(req.ID))

	log.Printf("Received a buy request for ID: %s from country: %s", req.ID, req.CountryCode)
	return WriteJSON(w, http.StatusOK, "Buy request received")
}

func (s *APIServer) handleConfirm(w http.ResponseWriter, r *http.Request) error {
	id := mux.Vars(r)["id"]
	log.Printf("Received a confirmation for ID: %s", id)
	return WriteJSON(w, http.StatusOK, "Confirmation received")
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
