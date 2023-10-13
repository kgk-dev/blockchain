package chain

import (
	"time"

	"github.com/kgk-dev/blockchain/block"
)

type BlockChain struct {
	length     int
	difficulty int
	Chain      []*block.Block
}

func Create(difficulty int) *BlockChain {
	blockchain := &BlockChain{}
	blockchain.GenesisBlock()
	blockchain.difficulty = difficulty
	return blockchain
}

func (bc *BlockChain) GenesisBlock() {
	genesisBlock := &block.Block{
		TimeStamp:    time.Now().String(),
		Data:         "Genesis Block",
		PreviousHash: "0",
	}
	genesisBlock.Hash = genesisBlock.CalculateHash()
	bc.Chain = append(bc.Chain, genesisBlock)
	bc.length++
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
	block := block.CreateBlock(
		time.Now().String(),
		data,
		bc.LatestBlock().Hash)
	ok := block.MineBlock(bc.difficulty)
	if ok {
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
		hash := currentBlock.CalculateHash()
		if hash != currentBlock.Hash ||
			currentBlock.PreviousHash != previousBlock.Hash {
			return false
		}
	}
	return true
}
