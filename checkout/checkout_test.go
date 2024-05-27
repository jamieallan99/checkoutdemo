package checkout

import (
	"testing"

	"github.com/shopspring/decimal"
)

func TestSumItems(t *testing.T) {
	testtable := []struct {
		name     string
		barcodes []string
		expected Price
	}{
		{"Basic Sum", []string{"A", "B", "C"}, Price{decimal.NewFromInt(100)}},
		{"Multibuy Sum", []string{"A", "A", "A"}, Price{decimal.NewFromInt(130)}},
	}
	for _, tr := range testtable {
		t.Run(tr.name, func(t *testing.T) {
			result := SumItems(tr.barcodes)
			if !tr.expected.Amount.Equal(result.Amount) {
				t.Errorf("Incorrect Sum expected: %d, got: %d", tr.expected, result)
			}
		})
	}
}