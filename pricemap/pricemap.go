package pricemap

import "checkoutdemo/price"

type PriceMap map[Barcode]PriceData

type Barcode string

type PriceData struct {
	Barcode Barcode
	Price price.Price
	Multibuy *Multibuy
}

type Multibuy struct {
	Count int
	Price price.Price
}