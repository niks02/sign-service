package transaction

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/signature_service/utils"

	"github.com/gorilla/mux"
)

// Service struct
type Service struct {
	repo repository
}

func newTransactionService(trepo repository) *Service {
	return &Service{
		repo: trepo,
	}
}

// SaveTransaction handles the storage of transactions
func (ts *Service) SaveTransaction(w http.ResponseWriter, r *http.Request) {
	request := Request{}
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		log.Printf("Error decoding the body: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	id := utils.GetID()
	ts.repo.SaveTransaction(id, request.Transaction)

	w.Header().Set("Content-Type", "application/json")
	res := Response{
		ID: id,
	}
	if err := json.NewEncoder(w).Encode(&res); err != nil {
		log.Printf("Error encoding the body: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

// AttachRoute adds transaction service related route to the mux router
func (ts *Service) AttachRoute(router *mux.Router) {
	router.HandleFunc("/transaction", ts.SaveTransaction).Methods("PUT")
}
