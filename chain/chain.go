package chain

import (
	"time"

	"github.com/kgk-dev/blockchain/block"
)

type Chain struct {
	Blocks []*block.Block
	Length int
}

func (c *Chain) AddBlock(block *block.Block) {
	c.Blocks = append(c.Blocks, block)
	c.Length++
}

func (c *Chain) FirstBlock() *block.Block {
	return c.Blocks[0]
}

func (c *Chain) EndBlock() *block.Block {
	return c.Blocks[c.Length-1]
}

func (c *Chain) IsValidChain() bool {
	for firstBlockIndex := 1; firstBlockIndex < c.Length; firstBlockIndex += 1 {
		currentBlock := c.Blocks[firstBlockIndex]
		previousBlock := c.Blocks[firstBlockIndex-1]
		hash := currentBlock.CalculateHash()
		if hash != currentBlock.Hash ||
			currentBlock.PreviousHash != previousBlock.Hash {
			return false
		}
	}
	return true
}

func (c *Chain) GenesisBlock() *block.Block {
	newBlock := &block.Block{
		TimeStamp:    time.Now(),
		PreviousHash: "000000",
	}
	newBlock.Hash = newBlock.CalculateHash()
	return newBlock
}

func Create() *Chain {
	return &Chain{}
}
