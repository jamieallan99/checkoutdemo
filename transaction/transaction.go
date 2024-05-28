package transaction

import (
	"fmt"

	"checkoutdemo/cache"
	"checkoutdemo/price"
	"checkoutdemo/pricemap"
)

type itemTally struct {
	count        int
	runningTotal price.Price
}

func (i *itemTally) CalculateItemCost(pd pricemap.PriceData) {
	if pd.Multibuy != nil {
		mbtotal, indtotal := countItem(i.count, pd.Multibuy.Count)
		mbprice := pd.Multibuy.Price.Mul(price.NewFromInt(int64(mbtotal)))
		indprice := pd.Price.Mul(price.NewFromInt(int64(indtotal)))
		i.runningTotal = mbprice.Add(indprice)
	} else {
		i.runningTotal = pd.Price.Mul(price.NewFromInt(int64(i.count)))
	}
}

type Transaction struct {
	ID           int64
	barcodes     []string
	itemTallies  map[string]*itemTally
	RunningTotal price.Price
}

func New(id int64, pm pricemap.PriceMap) Transaction {
	cache.Put(fmt.Sprint(id), pm)
	return Transaction{ID: id, barcodes: []string{}, itemTallies: make(map[string]*itemTally), RunningTotal: price.NewFromInt(0)}
}

func (t *Transaction) AddItem(barcode string) {
	data, err := cache.Get(fmt.Sprint(t.ID))
	if err != nil {
		fmt.Printf("Cache miss for transaction.ID: %d", t.ID)
		return
	}
	pm, ok := data.(pricemap.PriceMap)
	if ! ok {
		fmt.Printf("Data not in correct format, expected pricemap.PriceMap")
		return
	}
	t.barcodes = append(t.barcodes, barcode)
	if _, ok := t.itemTallies[barcode]; ok {
		t.itemTallies[barcode].count += 1
	} else {
		t.itemTallies[barcode] = &itemTally{1, price.NewFromInt(0)}
	}
	t.itemTallies[barcode].CalculateItemCost(pm[barcode])
	t.RunningTotal = t.SumItems()
}

func (t *Transaction) SumItems() price.Price {
	var sum price.Price
	for _, it := range t.itemTallies {
		sum = sum.Add(it.runningTotal)
	}
	return sum
}

func countItem(itemCount, multibuyCount int) (totalMultibuys, totalIndividual int) {
	totalMultibuys = itemCount / multibuyCount
	totalIndividual = itemCount % multibuyCount
	return
}
