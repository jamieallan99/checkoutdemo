package transaction

import (
	"checkoutdemo/price"
)

var pmap = map[string]int64{
	"A": 50,
	"B": 30,
	"C": 20,
	"D": 15,
}

var prices = map[string]*price.Price{}

var multibuys = map[string]Multibuy{
	"A": {
		Count:   3,
		Price:   price.NewFromInt(130),
	},
	"B": {
		Count:   2,
		Price:   price.NewFromInt(45),
	},
}

type Transaction struct {
	Barcodes []string
	RunningTotal price.Price
}

func (t *Transaction) AddItem(barcode string) {
	t.Barcodes = append(t.Barcodes, barcode)
	t.RunningTotal = t.SumItems()
}

func (t *Transaction) SumItems() price.Price {
	LoadPrices()
	var itemcount = map[string]*Tally{}
	var sum price.Price

	for _, b := range t.Barcodes {
		if _, ok := itemcount[b]; ok {
			itemcount[b].Count += 1
		} else {
			itemcount[b] = &Tally{b, 1}
		}
	}
	for _, t := range itemcount {
		sum = sum.Add(t.SumItem())
	} 
	return sum
}

type Multibuy struct {
	Count int64
	Price price.Price
}

type Tally struct {
	Barcode string
	Count   int64
}

func (t *Tally) SumItem() price.Price {
	var sum price.Price
	if m, ok := multibuys[t.Barcode]; ok {
		multibuycost := m.Price.Mul(price.NewFromInt(t.Count/m.Count))
		itemcost := prices[t.Barcode].Mul(price.NewFromInt(t.Count%m.Count))
		sum = multibuycost.Add(itemcost)
	} else {
		sum = prices[t.Barcode].Mul(price.NewFromInt(t.Count)) 
	}
	return sum
}

func LoadPrices() {
	for b, p:= range pmap {
		pprice := price.NewFromInt(p)
		prices[b] = &pprice
	}
}