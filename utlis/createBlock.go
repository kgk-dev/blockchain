package utlis

import (
	"github.com/kgk-dev/blockchain/block"
)

func CreateBlock(timeStamp, data, previousHash string) (*block.Block, error) {
	hash, err := CalculateHash(timeStamp, data, previousHash)
	if err == nil {
		return &block.Block{
			TimeStamp:    timeStamp,
			Data:         data,
			PreviousHash: previousHash,
			Hash:         hash,
		}, nil
	}
	return nil, err
}
