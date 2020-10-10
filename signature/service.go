package signature

import (
	"crypto"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/signature_service/transaction"
	"github.com/signature_service/utils"
)

// Service struct
type Service struct {
	signer crypto.Signer
	trepo  transaction.Reader
}

func newSignatureService(signer crypto.Signer, trepo transaction.Reader) *Service {
	return &Service{
		signer: signer,
		trepo:  trepo,
	}
}

// GetSignature handles the signing of transactions in request
func (ss *Service) GetSignature(w http.ResponseWriter, r *http.Request) {
	//unmarshal request body to Request struct
	request := Request{}
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		log.Printf("Error decoding the body: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	data, err := ss.trepo.GetTransactionsByIDs(request.IDs)
	if err != nil {
		log.Printf("Error while fetching data: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	signature, err := ss.signData(data)
	if err != nil {
		return
	}

	w.Header().Set("Content-Type", "application/json")
	res := Response{
		Message:   data,
		Signature: signature,
	}
	if err := json.NewEncoder(w).Encode(&res); err != nil {
		log.Printf("Error encoding the body: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (ss *Service) signData(data []string) (string, error) {
	bytes, err := json.Marshal(data)
	if err != nil {
		log.Printf("Error while data preperation for signing: %v", err)
		return "", err
	}
	signature, err := ss.signer.Sign(nil, bytes, crypto.Hash(0))
	if err != nil {
		log.Printf("Error while signing data: %v", err)
		return "", err
	}
	return string(utils.EncodeBase64(signature)), nil
}

// AttachRoute adds signature service related route to the mux router
func (ss *Service) AttachRoute(router *mux.Router) {
	router.HandleFunc("/signature", ss.GetSignature).Methods("POST")
}
