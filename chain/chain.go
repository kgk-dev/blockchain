package chain

import (
	"time"

	"github.com/kgk-dev/blockchain/block"
	"github.com/kgk-dev/blockchain/utlis"
)

type BlockChain struct {
	length int
	Chain  []*block.Block
}

func (bc *BlockChain) GenesisBlock() *block.Block {
	if genesisBlock, err := utlis.CreateBlock(time.Now().String(), "Genesis Block", "0"); err == nil {
		return genesisBlock
	}
	return nil
}

func Create() *BlockChain {
	bc := &BlockChain{}
	bc.Chain = append(bc.Chain, bc.GenesisBlock())
	bc.length++
	return bc
}

func (bc *BlockChain) FirstBlock() *block.Block {
	return bc.Chain[0]
}

func (bc *BlockChain) LatestBlock() *block.Block {
	chainLength := len(bc.Chain)
	if chainLength == 1 {
		return bc.FirstBlock()
	}
	return bc.Chain[chainLength-1]
}

func (bc *BlockChain) AddBlock(data string) *BlockChain {
	block, err := utlis.CreateBlock(
		time.Now().String(),
		data,
		bc.LatestBlock().Hash)
	if err == nil {
		bc.Chain = append(bc.Chain, block)
		bc.length++
		return bc
	}
	return bc
}

func (bc *BlockChain) IsChainValid() bool {
	for index := 1; index < bc.length; index += 1 {
		currentBlock := bc.Chain[index]
		previousBlock := bc.Chain[index-1]
		if hash, err := utlis.CalculateHash(
			currentBlock.TimeStamp, currentBlock.Data, currentBlock.PreviousHash,
		); err == nil && hash != currentBlock.Hash {
			return false
		}
		if currentBlock.PreviousHash != previousBlock.Hash {
			return false
		}
	}
	return true
}
