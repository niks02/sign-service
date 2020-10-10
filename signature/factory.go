package signature

import (
	"crypto"
	"crypto/ed25519"
	"log"

	"github.com/gorilla/mux"
	"github.com/signature_service/config"
	"github.com/signature_service/transaction"
	"github.com/signature_service/utils"
)

// GetSigner returns implementation of configured signing algorithm
func GetSigner() crypto.Signer {
	repo := config.GetRepo()

	privateKeyStr := repo.GetPrivateKey()
	bytes, err := utils.DecodeBase64([]byte(privateKeyStr))
	if err != nil {
		log.Fatalf("Error decoding private key: %v", err)
	}

	switch repo.GetSigningAlgo() {

	case "ed25519":
		privateKey := ed25519.PrivateKey(bytes)
		return privateKey
	}
	return nil
}

// InitService initialises the Signature Service
func InitService(router *mux.Router) {
	cs := newSignatureService(GetSigner(), transaction.GetTransactionReader())
	cs.AttachRoute(router)
}
