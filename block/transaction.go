package block

import "fmt"

type Transaction struct {
	senderAddress   string
	receiverAddress string
	amount          int
}

func NewTransaction(
	senderAddress,
	receiverAddress string,
	amount int) *Transaction {
	return &Transaction{senderAddress, receiverAddress, amount}
}

func (t *Transaction) String() string {
	return fmt.Sprintf("%v %v %v", t.senderAddress, t.receiverAddress, t.amount)
}
