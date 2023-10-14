package transaction

import (
	"fmt"
)

type Transaction struct {
	SenderAddress   string
	ReceiverAddress string
	Amount          float64
}

func (t *Transaction) String() string {
	return fmt.Sprintf(
		"Sender: %v,\nReceiver: %v,\nAmount: %v,\n",
		t.SenderAddress, t.ReceiverAddress, t.Amount,
	)
}

func Create(sender, receiver string, amount float64) *Transaction {
	return &Transaction{
		SenderAddress:   sender,
		ReceiverAddress: receiver,
		Amount:          amount,
	}
}
