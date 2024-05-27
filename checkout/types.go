package checkout

import "checkoutdemo/price"

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
