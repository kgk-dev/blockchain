package main

import (
	"fmt"

	"github.com/kgk-dev/blockchain/block"
	"github.com/kgk-dev/blockchain/chain"
)

func main() {
	BlockChain := chain.Create(3)
	BlockChain.AddBlock(*block.NewTransaction("sender A", "sender B", 100))
	fmt.Println(*BlockChain)
	fmt.Println("Valid", BlockChain.IsChainValid())
}
