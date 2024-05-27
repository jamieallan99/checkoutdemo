package tally

import "checkoutdemo/price"

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

type Multibuy struct {
	Count int64
	Price price.Price
}

type Tally struct {
	Barcode string
	Count   int64
}

func (t *Tally) SumItem() price.Price {
	LoadPrices()
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