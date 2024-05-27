package checkout

type Multibuy struct {
	Count   int
	Price   int
}

type Tally struct {
	Barcode string
	Count   int
}

func (t *Tally) SumItem() int {
	var sum int
	if m, ok := multibuys[t.Barcode]; ok {
		multibuycost := (t.Count/m.Count) * m.Price
		itemcost := (t.Count%m.Count) * prices[t.Barcode]
		sum = multibuycost + itemcost
	} else {
		sum = t.Count * prices[t.Barcode]
	}
	return sum
}
