package main

import (
	"fmt"

	"demo.com/pricetax/filemanager"
	"demo.com/pricetax/price"
)

func main() {
	taxes := []float64{0.1, 2.2, 3.5, 8.3, 15.2}
	for _, tax := range taxes {
		fm := filemanager.New("price.txt", fmt.Sprintf("result%0.f.json", tax))
		taxIncluding := price.NewTaxIncludedPriceJob(fm, tax)
		taxIncluding.Process()
	}

}
