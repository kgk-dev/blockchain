package transactions

import (
	"encoding/json"
	"fmt"

	"github.com/kgk-dev/blockchain/transaction"
)

type TransactionList struct {
	Transactions []*transaction.Transaction
	Length       int
	Cap          int
}

func (tl *TransactionList) Add(newTransaction *transaction.Transaction) bool {
	if tl.Length < tl.Cap {
		tl.Transactions = append(tl.Transactions, newTransaction)
		return true
	}
	return false
}

func (tl *TransactionList) String() string {
	jsonData, _ := json.Marshal(tl.Transactions)
	return fmt.Sprintf("%s", jsonData)
}

func Create(size int) *TransactionList {
	transList := &TransactionList{}
	transList.Cap = size
	return transList
}
