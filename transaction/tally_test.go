package transaction

import (
	"checkoutdemo/price"
	"testing"
)

func TestSumItem(t *testing.T) {
	testTable := []struct {
		name  string
		tally Tally
		total price.Price
	}{
		{
			"Simple Count", 
			Tally{"A", 1}, 
			price.NewFromInt(10),
		},
	}
	for _, tr := range testTable {
		p := price.NewFromInt(10)
		prices = map[string]*price.Price{
			"A": &p,
		}
		t.Run(tr.name, func(t *testing.T) {
			total := tr.tally.SumItem()
			if !tr.total.Equal(total) {
				t.Errorf("Incorrect total expected: %s, got: %s", tr.total.String(), total.String())
			}
		})
	}
}
