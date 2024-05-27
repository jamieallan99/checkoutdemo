package checkout

import (
	"testing"

	"checkoutdemo/price"
)

func TestSumItems(t *testing.T) {
	testtable := []struct {
		name     string
		barcodes []string
		expected price.Price
	}{
		{"Basic Sum", []string{"A", "B", "C"}, price.NewFromInt(100)},
		{"Multibuy Sum", []string{"A", "A", "A"}, price.NewFromInt(130)},
		{"Complex Multibuy Sum", []string{"A", "A", "A", "B", "B", "B", "C", "A", "A", "B", "A"}, price.NewFromInt(370)},
	}
	for _, tr := range testtable {
		t.Run(tr.name, func(t *testing.T) {
			result := SumItems(tr.barcodes)
			if !tr.expected.Equal(result) {
				t.Errorf("Incorrect Sum expected: %s, got: %s", tr.expected.String(), result.String())
			}
		})
	}
}
