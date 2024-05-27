package checkout

var prices = map[string]int {
	"A": 50,
	"B": 30,
	"C": 20,
}

func SumItems(barcodes []string) int {
	var sum int
	for _,b := range barcodes {
		sum += prices[b]
	}
	return sum
}