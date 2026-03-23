package main

import (
	// "fmt"
	"os"
	// "path/filepath"

	"example.com/tax-calc/cmdmanager"
	// "example.com/tax-calc/filemanager"
	"example.com/tax-calc/prices"
)

func main() {

	taxRates := []float64{0, 0.07, 0.1, 0.15}

	// outDir := "results/"
	os.MkdirAll("results", 0755)
	for _, tax := range taxRates {

		// fileName := filepath.Join(outDir, fmt.Sprintf("result_%.0f.json", 1000*tax))

		// fm := filemanager.New("prices.txt", fileName)

		cmd := cmdmanager.New()
		taxedPrice := prices.NewTaxInclucdedPrices(cmd, tax)

		taxedPrice.Process()
	}

	// fmt.Println(result)

}
