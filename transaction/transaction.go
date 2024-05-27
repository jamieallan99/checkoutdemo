package transaction

import (
	"checkoutdemo/price"
	"checkoutdemo/tally"
)

type Transaction struct {
	Barcodes []string
	RunningTotal price.Price
}

func (t *Transaction) AddItem(barcode string) {
	t.Barcodes = append(t.Barcodes, barcode)
	t.RunningTotal = t.SumItems()
}

func (t *Transaction) SumItems() price.Price {
	var itemcount = map[string]*tally.Tally{}
	var sum price.Price

	for _, b := range t.Barcodes {
		if _, ok := itemcount[b]; ok {
			itemcount[b].Count += 1
		} else {
			itemcount[b] = &tally.Tally{b, 1}
		}
	}
	for _, t := range itemcount {
		sum = sum.Add(t.SumItem())
	} 
	return sum
}