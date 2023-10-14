package main

import (
	"fmt"
	"time"

	"github.com/kgk-dev/blockchain/block"
	"github.com/kgk-dev/blockchain/chain"
	"github.com/kgk-dev/blockchain/transaction"
	"github.com/kgk-dev/blockchain/transactions"
)

const NumberOfTransaction = 1
const Difficulty = 4

func createBlock(
	chain *chain.Chain,
	transactions *transactions.TransactionList) (*block.Block, bool) {
	mineBlock := block.Build(
		time.Now(),
		transactions,
		chain.EndBlock().Hash,
	)
	return mineBlock(Difficulty)
}

func main() {
	//First, create chain
	chain := chain.Create()
	chain.AddBlock(chain.GenesisBlock())

	//List of transactions
	transactions := transactions.Create(2)
	transaction1 := transaction.Create("A", "B", 10)
	transaction2 := transaction.Create("A", "C", 20)
	transactions.Add(transaction1)
	transactions.Add(transaction2)

	//create new Block and then, If Block shows POW, and Block is added to Chain
	if newBlock, ok := createBlock(chain, transactions); ok {
		chain.AddBlock(newBlock)
	}

	fmt.Println("Chain: ", chain)
}
