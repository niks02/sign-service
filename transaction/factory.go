package transaction

import (
	"github.com/gorilla/mux"
)

var repo *trepo

func init() {
	repo = newTransactionRepository()
}

// Reader interface
type Reader interface {
	GetTransactionByID(id uint64) (string, error)
	GetTransactionsByIDs(ids []uint64) ([]string, error)
}

type repository interface {
	Reader
	SaveTransaction(id uint64, data string)
}

// GetTransactionReader returns transactionReader interface
func GetTransactionReader() Reader {
	return repo
}

// InitService initialises the Transaction Service
func InitService(router *mux.Router) {
	cs := newTransactionService(repo)
	cs.AttachRoute(router)
}
