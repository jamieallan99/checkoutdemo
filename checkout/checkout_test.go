package checkout

import "testing"

func TestSumItems(t *testing.T) {
	testtable := []struct {
		name     string
		barcodes []string
		expected int
	}{
		{"Basic Sum", []string{"A", "B", "C"}, 100},
		{"Multibuy Sum", []string{"A", "A", "A"}, 130},
	}
	for _, tr := range testtable {
		t.Run(tr.name, func(t *testing.T) {
			result := SumItems(tr.barcodes)
			if result != tr.expected {
				t.Errorf("Incorrect Sum expected: %d, got: %d", tr.expected, result)
			}
		})
	}
}