package config

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Service struct
type Service struct {
	repo Repository
}

func newConfigService(repo Repository) *Service {
	return &Service{
		repo: repo,
	}
}

// GetPublicKey returns the configured publicKey
func (cs *Service) GetPublicKey(w http.ResponseWriter, r *http.Request) {
	pk := cs.repo.GetPublicKey()
	w.Header().Set("Content-Type", "application/json")
	res := PublicKeyResponse{
		PublicKey: pk,
	}
	if err := json.NewEncoder(w).Encode(&res); err != nil {
		log.Printf("Error encoding the body: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

// AttachRoute adds config service related route to the mux router
func (cs *Service) AttachRoute(router *mux.Router) {
	router.HandleFunc("/public_key", cs.GetPublicKey).Methods("GET")
}
