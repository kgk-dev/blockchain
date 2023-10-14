package block

import (
	"testing"
	"time"

	"github.com/kgk-dev/blockchain/transaction"
)

func TestBlock(t *testing.T) {
	newTransaction := transaction.Create("A", "B", 100)
	mine := Build(time.Now(), [NumOfTransactions]*transaction.Transaction{newTransaction}, "0")
	_, ok := mine(5)
	if !ok {
		t.Fatalf("Block cannot create")
	}
}
