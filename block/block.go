package block

import (
	sha256 "crypto/sha256"
	"encoding/json"
	"fmt"
	"strings"
)

type Block struct {
	TimeStamp    string
	Data         string
	PreviousHash string
	Hash         string
	Nonce        int
}

func CreateBlock(timeStamp, data, previousHash string) *Block {
	newBlock := &Block{
		TimeStamp:    timeStamp,
		Data:         data,
		PreviousHash: previousHash,
	}
	return newBlock
}

func (b *Block) CalculateHash() string {
	jsonData, _ := json.Marshal(fmt.Sprintf("%v%v%v%v", b.TimeStamp, b.Data, b.PreviousHash, b.Nonce))
	return fmt.Sprintf("%X", sha256.Sum256(jsonData))
}

func (b *Block) MineBlock(difficulty int) bool {
	prefix := make([]string, difficulty+1)
	b.Hash = b.CalculateHash()
	for b.Hash[:difficulty] != strings.Join(prefix, "0") {
		fmt.Println("hash in mine", b.Hash)
		b.Nonce++
		b.Hash = b.CalculateHash()
	}
	return true
}

func (b *Block) String() string {
	return fmt.Sprintf("Timestamp: %v\nData: %v\nPreviousHash: %v\nHash: %v\nNonce: %v\n", b.TimeStamp, b.Data, b.PreviousHash, b.Hash, b.Nonce)
}
