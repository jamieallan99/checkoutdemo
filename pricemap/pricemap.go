package pricemap

import "checkoutdemo/price"

type PriceMap map[string]PriceData

type PriceData struct {
	Barcode string
	Price price.Price
	Multibuy *Multibuy
}

type Multibuy struct {
	Count int64
	Price price.Price
}