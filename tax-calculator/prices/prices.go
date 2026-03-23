package prices

import (
	"errors"
	"fmt"

	"example.com/tax-calc/conversion"
	// "example.com/tax-calc/filemanager"
	"example.com/tax-calc/iomanager"
)

type TaxIncludedPricesJob struct {
	IOManager         iomanager.IOManager `json:"-"`
	TaxRate           float64             `json:"tax_rate"`
	InputPrices       []float64           `json:"input_prices"`
	TaxInlcudedPrices map[string]string   `json:"tax_included_prices"`
}

func NewTaxInclucdedPrices(io iomanager.IOManager, taxRates float64) *TaxIncludedPricesJob {
	return &TaxIncludedPricesJob{
		IOManager:   io,
		InputPrices: []float64{},
		TaxRate:     taxRates,
	}
}

func (job *TaxIncludedPricesJob) Process() error {

	err := job.LoadData("prices.txt")
	if err != nil {
		return errors.New(err.Error())
	}

	result := make(map[string]string)
	for _, price := range job.InputPrices {
		result[fmt.Sprintf("%.2f", price)] = fmt.Sprintf("%.2f", price*(1+job.TaxRate))
	}

	job.TaxInlcudedPrices = result

	err = job.IOManager.WriteResults(job)
	if err != nil {
		return errors.New(err.Error())
	}
	return nil
}

func (job *TaxIncludedPricesJob) LoadData(fileName string) error {
	lines, err := job.IOManager.ReadLines()

	if err != nil {
		// fmt.Println(err)
		return errors.New("Error Reading from the file")
	}
	prices, err := conversion.StringsToFloats(lines)
	if err != nil {
		// fmt.Println(err)
		return errors.New("Error Converting to float")
	}

	job.InputPrices = prices
	return nil
}
