package transaction

import (
	"errors"
	"fmt"

	"github.com/signature_service/utils"
)

type trepo struct {
	transactions map[uint64]string
}

func newTransactionRepository() *trepo {
	return &trepo{
		transactions: map[uint64]string{},
	}
}

// GetTransactionByID returns transaction provided it's id
func (repo *trepo) GetTransactionByID(id uint64) (string, error) {
	val, ok := repo.transactions[id]
	if ok != true {
		err := errors.New("Invalid Transaction")
		return "", err
	}
	base64Res := utils.EncodeBase64([]byte(val))
	return string(base64Res), nil
}

// GetTransactionsByIDs returns transactions by there id's
func (repo *trepo) GetTransactionsByIDs(ids []uint64) ([]string, error) {
	if !repo.areTransactionsValid(ids) {
		return nil, fmt.Errorf("Some transactions are missing")
	}
	return repo.getTransactionsByIDs(ids), nil
}

func (repo *trepo) getTransactionsByIDs(ids []uint64) []string {
	var res []string
	for _, id := range ids {
		res = append(res, repo.transactions[id])
	}
	return res
}

func (repo *trepo) areTransactionsValid(ids []uint64) bool {
	for _, id := range ids {
		if _, ok := repo.transactions[id]; !ok {
			return false
		}
	}
	return true
}

// SaveTransaction saves the transaction with its id
func (repo *trepo) SaveTransaction(id uint64, data string) {
	repo.transactions[id] = data
}
