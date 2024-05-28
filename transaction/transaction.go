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
	Barcodes     []string
	RunningTotal price.Price
}

func New(id int64) Transaction {
	cache.Put(fmt.Sprint(id), 123)
	return Transaction{ID: id, Barcodes: []string{}, RunningTotal: price.NewFromInt(0)}
}

func (t *Transaction) AddItem(barcode string) {
	t.Barcodes = append(t.Barcodes, barcode)
	t.RunningTotal = t.SumItems()
}

func (t *Transaction) SumItems() price.Price {
	var itemcounts = map[string]*itemTally{}
	var sum price.Price

	for _, b := range t.Barcodes {
		if _, ok := itemcounts[b]; ok {
			itemcounts[b].count += 1
		} else {
			itemcounts[b] = &itemTally{1, price.NewFromInt(0)}
		}
	}
	for _, t := range itemcounts {
		sum = sum.Add(t.runningTotal)
	}
	return sum
}

func countItem(itemCount, multibuyCount int) (totalMultibuys, totalIndividual int) {
	totalMultibuys = itemCount / multibuyCount
	totalIndividual = itemCount % multibuyCount
	return
}
