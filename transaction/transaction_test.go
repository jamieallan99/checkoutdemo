package transaction

import (
	"testing"

	"checkoutdemo/price"
)

func TestAddItem(t *testing.T) {
	testtable := []struct {
		name     string
		barcodes []string
		total price.Price
	}{
		{"Basic Sum", []string{"A", "B", "C"}, price.NewFromInt(100)},
		{"Multibuy Sum", []string{"A", "A", "A"}, price.NewFromInt(130)},
		{"Complex Multibuy Sum", []string{"A", "A", "A","B","B","B","C","A","A","B","A"}, price.NewFromInt(370)},
	}
	for _, tr := range testtable {
		t.Run(tr.name, func(t *testing.T) {
			var transaction Transaction
			for _, b := range tr.barcodes {
				transaction.AddItem(b)
			}
			if !tr.total.Equal(transaction.RunningTotal) {
				t.Errorf("Incorrect Sum expected: %s, got: %s", tr.total.String(), transaction.RunningTotal.String())
			}
		})
	}
}