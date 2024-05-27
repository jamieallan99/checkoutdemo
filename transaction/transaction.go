package transaction

import (
	"checkoutdemo/price"
)

type Transaction struct {
	Barcodes []string
	RunningTotal price.Price
}