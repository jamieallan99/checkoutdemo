package database

import (
	"checkoutdemo/price"
	"checkoutdemo/pricemap"
)

// Mock implementation of a DB load
func LoadPrices() pricemap.PriceMap {
	return pricemap.PriceMap{
		"A": pricemap.PriceData{
			Barcode: "A",
			Price: price.NewFromInt(50),
			Multibuy: &pricemap.Multibuy{
				Count: 3,
				Price: price.NewFromInt(130),
			},
		},
		"B": pricemap.PriceData{
			Barcode: "D",
			Price: price.NewFromInt(30),
			Multibuy: &pricemap.Multibuy{
				Count: 2,
				Price: price.NewFromInt(45),
			},
		},
		"C": pricemap.PriceData{
			Barcode: "C",
			Price: price.NewFromInt(20),
		},
		"D": pricemap.PriceData{
			Barcode: "D",
			Price: price.NewFromInt(15),
		},
	}
}