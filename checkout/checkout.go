package checkout

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

func SumItems(barcodes []string) price.Price {
	LoadPrices()
	var itemcount = map[string]*Tally{}
	var sum price.Price

	for _, b := range barcodes {
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

func LoadPrices() {
	for b, p:= range pmap {
		pprice := price.NewFromInt(p)
		prices[b] = &pprice
	}
}