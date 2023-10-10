package block

import "fmt"

type Block struct {
	TimeStamp    string
	Data         string
	PreviousHash string
	Hash         string
}

func (b *Block) String() string {
	return fmt.Sprintf("Timestamp: %v\nData: %v\nPreviousHash: %v\nHash: %v\n", b.TimeStamp, b.Data, b.PreviousHash, b.Hash)
}
