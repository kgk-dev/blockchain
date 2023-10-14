package transaction

import (
	"testing"
)

func TestTransaction(t *testing.T) {
	transaction := Create("A", "B", 100)
	expected := &Transaction{"A", "B", 100}
	if *expected != *transaction {
		t.Fatalf("Exptected: %v, but got %v\n", expected, transaction)
	}
}
