package transaction

import (
	"testing"
	"time"

	"checkoutdemo/cache"
	"checkoutdemo/price"
	"checkoutdemo/pricemap"
)

var testPrices = pricemap.PriceMap{
	"A": pricemap.PriceData{
		Barcode: "A",
		Price: price.NewFromInt(50),
		Multibuy: &pricemap.Multibuy{
			Count: 3,
			Price: price.NewFromInt(130),
		},
	},
	"B": pricemap.PriceData{
		Barcode: "B",
		Price: price.NewFromInt(30),
		Multibuy: &pricemap.Multibuy{
			Count: 2,
			Price: price.NewFromInt(45),
		},
	},
	"C": pricemap.PriceData{
		Barcode: "C",
		Price: price.NewFromInt(20),
		Multibuy: nil,
	},
}

var (
	basicItemTally = tallyMap {
		"A": {1, price.NewFromInt(50)},
		"B": {1, price.NewFromInt(30)},
		"C": {1, price.NewFromInt(20)},
	}
	multibuyItemTally = tallyMap {
		"A": {3, price.NewFromInt(130)},
	}
	complexItemTally = tallyMap {
		"A": {6, price.NewFromInt(260)},
		"B": {4, price.NewFromInt(90)},
		"C": {1, price.NewFromInt(20)},
	}
)

func TestSumItems(t *testing.T) {
	cache.ManualInit()
	defer cache.KillStore()
	testtable := []struct {
		name     string
		itemTallies tallyMap
		expected price.Price
	}{
		{"Basic Sum", basicItemTally, price.NewFromInt(100)},
		{"Multibuy Sum", multibuyItemTally, price.NewFromInt(130)},
		{"Complex Multibuy Sum", complexItemTally, price.NewFromInt(370)},
	}
	for _, tr := range testtable {
		t.Run(tr.name, func(t *testing.T) {
			var transaction = New(time.Now().Unix(), testPrices)
			transaction.itemTallies = tr.itemTallies
			result := transaction.SumItems()
			if !tr.expected.Equal(result) {
				t.Errorf("Incorrect Sum expected: %s, got: %s", tr.expected.String(), result.String())
			}
		})
	}
}

func TestAddItem(t *testing.T) {
	cache.ManualInit()
	defer cache.KillStore()
	testtable := []struct {
		name     string
		barcodes []pricemap.Barcode
		total    price.Price
	}{
		{"Basic Sum", []pricemap.Barcode{"A", "B", "C"}, price.NewFromInt(100)},
		{"Multibuy Sum", []pricemap.Barcode{"A", "A", "A"}, price.NewFromInt(130)},
		{"Complex Multibuy Sum", []pricemap.Barcode{"A", "A", "A", "B", "B", "B", "C", "A", "A", "B", "A"}, price.NewFromInt(370)},
	}
	for _, tr := range testtable {
		t.Run(tr.name, func(t *testing.T) {
			var transaction = New(time.Now().Unix(), testPrices)
			for _, b := range tr.barcodes {
				transaction.AddItem(b)
			}
			if !tr.total.Equal(transaction.RunningTotal) {
				t.Errorf("Incorrect Sum expected: %s, got: %s", tr.total.String(), transaction.RunningTotal.String())
			}
		})
	}
}

func TestCountItem(t *testing.T) {
	testTable := []struct {
		name            string
		itemcount       int
		multibuyCount   int
		totalMultibuy   int
		totalIndividual int
	}{
		{"Simple Count", 1, 3, 0, 1},
		{"Multibuy Count", 3, 3, 1, 0},
		{"Complex Multibuy Count", 4, 3, 1, 1},
	}
	for _, tr := range testTable {
		t.Run(tr.name, func(t *testing.T) {
			totalMultibuy, totalIndividual := countItem(tr.itemcount, tr.multibuyCount)
			if tr.totalMultibuy != totalMultibuy {
				t.Errorf("Incorrect total of multibuys, expected: %d, got: %d", tr.totalMultibuy, totalMultibuy)
			}
			if tr.totalIndividual != totalIndividual {
				t.Errorf("Incorrect total of individual items, expected: %d, got: %d", tr.totalIndividual, totalIndividual)
			}
		})
	}
}

func TestCalculateItemCost(t *testing.T) {
	testTable := []struct {
		name          string
		priceData     pricemap.PriceData
		tally         itemTally
		expectedTotal price.Price
	}{
		{
			"Simple case",
			pricemap.PriceData{
				Barcode: "",
				Price:   price.NewFromInt(10),
			},
			itemTally{1, price.NewFromInt(0)},
			price.NewFromInt(10),
		},
	}
	for _, tr := range testTable {
		t.Run(tr.name, func(t *testing.T) {
			tr.tally.CalculateItemCost(tr.priceData)
			if !tr.expectedTotal.Equal(tr.tally.runningTotal) {
				t.Errorf("Incorrect total calculated, expected: %s, got: %s", tr.expectedTotal.String(), tr.tally.runningTotal.String())
			}
		})
	}
}
