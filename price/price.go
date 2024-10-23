package price

import (
	"fmt"

	"demo.com/pricetax/conversion"
	"demo.com/pricetax/filemanager"
)

type TaxIncludedPriceJob struct {
	IOManager        filemanager.Filemanager `json:"-"`
	TaxRate          float64                 `json:"tax_rate"`
	InputPrices      []float64               `json:"input_prices"`
	TaxIncludedPrice map[string]string       `json:"tax_included_price"`
}

func NewTaxIncludedPriceJob(fm filemanager.Filemanager, TaxRate float64) *TaxIncludedPriceJob {
	return &TaxIncludedPriceJob{
		IOManager:   fm,
		InputPrices: []float64{200, 300, 400},
		TaxRate:     TaxRate,
	}
}

func (job *TaxIncludedPriceJob) LoadData() {
	lines, err := job.IOManager.ReadFile()
	if err != nil {
		fmt.Println(err)
		return
	}
	prices, err := conversion.StringToFloats(lines)
	if err != nil {
		fmt.Println(err)

		return
	}

	job.InputPrices = prices

}

func (job *TaxIncludedPriceJob) Process() {
	job.LoadData()
	result := make(map[string]string)
	for _, price := range job.InputPrices {
		taxAddedPrice := price * (1 + job.TaxRate)
		result[fmt.Sprintf("%.2f", price)] = fmt.Sprintf("%.2f", taxAddedPrice)
	}
	job.TaxIncludedPrice = result
	fmt.Println(result)
	job.IOManager.WriteJson(job)
}
