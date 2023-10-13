package main

import (
	"fmt"

	"github.com/kgk-dev/blockchain/chain"
)

func main() {
	BlockChain := chain.Create(3)
	BlockChain.AddBlock("I'm programmer")
	BlockChain.AddBlock("I'm blockchain developer")
	BlockChain.AddBlock("I'm software engineer")
	fmt.Println(*BlockChain)
	fmt.Println("Valid", BlockChain.IsChainValid())
}
