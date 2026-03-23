package main

import (
	// "fmt"
	"fmt"
	"os"

	"path/filepath"

	// "example.com/tax-calc/cmdmanager"
	"example.com/tax-calc/filemanager"
	"example.com/tax-calc/prices"
)

func main() {

	taxRates := []float64{0, 0.07, 0.1, 0.15}

	outDir := "results/"
	file := "prices.txt"
	os.MkdirAll("results", 0755)
	for _, tax := range taxRates {

		fileName := filepath.Join(outDir, fmt.Sprintf("result_%.0f.json", 1000*tax))

		fm := filemanager.New(file, fileName)

		// cmd := cmdmanager.New()
		taxedPrice := prices.NewTaxInclucdedPrices(fm, tax)

		err := taxedPrice.Process()
		if err != nil {
			fmt.Println(err)
			return
		}
	}

	// fmt.Println(result)

}
