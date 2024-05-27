package checkout

var prices = map[string]int{
	"A": 50,
	"B": 30,
	"C": 20,
	"D": 15,
}

var multibuys = map[string]Multibuy{
	"A": {
		Count:   3,
		Price:   130,
	},
	"B": {
		Count:   2,
		Price:   45,
	},
}

func SumItems(barcodes []string) int {
	var itemcount = map[string]*Tally{}
	var sum int

	for _, b := range barcodes {
		if _, ok := itemcount[b]; ok {
			itemcount[b].Count += 1
		} else {
			itemcount[b] = &Tally{b, 1}
		}
	}
	for _, t := range itemcount {
		sum += t.SumItem()
	} 
	return sum
}

