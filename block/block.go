package block

import (
	sha256 "crypto/sha256"
	"fmt"
	"strings"
	"time"

	"github.com/kgk-dev/blockchain/transactions"
)

type MineFn func(difficult int) (*Block, bool)

type Block struct {
	TimeStamp    time.Time
	Transactions *transactions.TransactionList
	PreviousHash string
	Hash         string
	Nonce        int
}

func (b *Block) CalculateHash() string {
	return fmt.Sprintf("%x", (sha256.Sum256([]byte(fmt.Sprintf(
		"%v %v %v", b.TimeStamp, b.PreviousHash, b.Nonce,
	)))))
}

func (b *Block) mineBlock(difficulty int) string {
	workload := strings.Join(make([]string, difficulty+1), "0")
	hash := b.CalculateHash()
	for hash[:difficulty] != workload {
		b.Nonce++
		hash = b.CalculateHash()
	}
	return hash
}

func (b *Block) String() string {
	return fmt.Sprintf(
		"Time: %v,\nHash: %v,\nPreviousHash: %v,\nTransactions: %v,\nNonce: %v,\n",
		b.TimeStamp, b.Hash, b.PreviousHash, b.Transactions, b.Nonce,
	)
}

func Build(
	timeStamp time.Time,
	transaction *transactions.TransactionList,
	previousHash string,
) MineFn {
	newBlock := &Block{
		TimeStamp:    timeStamp,
		Transactions: transaction,
		PreviousHash: previousHash,
	}
	return func(difficult int) (*Block, bool) {
		hash := newBlock.mineBlock(difficult)
		if hash != "" {
			newBlock.Hash = hash
			return newBlock, true
		}
		return nil, false
	}
}
