package main

import (
	// "fmt"

	// "errors"
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

	doneChannels := make([]chan bool, len(taxRates))
	errorChannels := make([]chan error, len(taxRates))

	for i, tax := range taxRates {
		doneChannels[i] = make(chan bool)
		errorChannels[i] = make(chan error)

		fileName := filepath.Join(outDir, fmt.Sprintf("result_%.0f.json", 1000*tax))

		fm := filemanager.New(file, fileName)

		// cmd := cmdmanager.New()
		taxedPrice := prices.NewTaxInclucdedPrices(fm, tax)

		go taxedPrice.Process(doneChannels[i], errorChannels[i])

		// if err != nil {
		// 	fmt.Println(err)
		// 	return
		// }
	}

	for i := range taxRates {
		select {
		case err := <-errorChannels[i]:
			if err != nil {
				fmt.Println(err)
			}
		case <-doneChannels[i]:
			fmt.Println("Done")
		}
	}

	// for _, errChannel := range errorChannels {
	// 	<-errChannel
	// }

	// for _, channel := range doneChannels {
	// 	<-channel
	// }

	// fmt.Println(result)

}
