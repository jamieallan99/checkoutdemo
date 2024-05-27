package price

import "github.com/shopspring/decimal"

// Wrapper on decimal so that implementation can be swapped if needed
type Price struct {
	amount decimal.Decimal
}

// Equal returns whether the numbers represented by p and p2 are equal
func (p *Price) Equal(p2 Price) bool {
	return p.amount.Equal(p2.amount)
}

// Add returns p + p2
func (p *Price) Add(p2 Price) Price {
	return Price{p.amount.Add(p2.amount)}
}

// Mul returns p * p2
func (p *Price) Mul(p2 Price) Price {
	return Price{p.amount.Mul(p2.amount)}
}


// String returns the string representation of the price
// with the fixed point.
//
// Example:
//
//	d := New(-12345, -3)
//	println(d.String())
//
// Output:
//
//	-12.345
func(p *Price) String() string {
	return p.amount.String()
}

// NewFromInt converts an int64 to Price.
//
// Example:
//
//	NewFromInt(123).String() // output: "123"
//	NewFromInt(-10).String() // output: "-10"
func NewFromInt(i int64) Price {
	return Price{decimal.NewFromInt(i)}
}