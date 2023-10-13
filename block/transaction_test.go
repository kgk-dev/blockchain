package block

import (
	"fmt"
	"testing"
)

func TestTransaction(t *testing.T) {
	newTransaction := NewTransaction("A", "B", 100)
	if fmt.Sprintf("%v", newTransaction) != "A B 100" {
		t.Fatal("Fail")
	}
}
