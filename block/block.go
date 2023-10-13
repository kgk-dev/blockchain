package block

import (
	sha256 "crypto/sha256"
	"encoding/json"
	"fmt"
	"strings"
)

type Block struct {
	TimeStamp    string
	Transaction  Transaction
	PreviousHash string
	Hash         string
	Nonce        int
}

func CreateBlock(timeStamp string, transaction Transaction, previousHash string) *Block {
	newBlock := &Block{
		TimeStamp:    timeStamp,
		Transaction:  transaction,
		PreviousHash: previousHash,
	}
	return newBlock
}

func (b *Block) CalculateHash() string {
	jsonData, _ := json.Marshal(
		fmt.Sprintf("%v %v %v %v", b.TimeStamp, b.Transaction, b.PreviousHash, b.Nonce),
	)
	return fmt.Sprintf("%X", sha256.Sum256(jsonData))
}

func (b *Block) MineBlock(difficulty int) bool {
	prefix := make([]string, difficulty+1)
	b.Hash = b.CalculateHash()
	for b.Hash[:difficulty] != strings.Join(prefix, "0") {
		b.Nonce++
		b.Hash = b.CalculateHash()
	}
	return true
}

func (b *Block) String() string {
	return fmt.Sprintf("Timestamp: %v\nData: %v\nPreviousHash: %v\nHash: %v\nNonce: %v\n", b.TimeStamp, b.Transaction, b.PreviousHash, b.Hash, b.Nonce)
}
