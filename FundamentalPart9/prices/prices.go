package prices

import (
	"FundamentalPart9/manager"
	"fmt"
)

type TaxIncludedPriceJob struct {
	TaxRate           float64            `json:"tax_rate"`
	InputPrices       []float64          `json:"input_prices"`
	TaxIncludedPrices map[string]float64 `json:"tax_included_prices"`
	IOManager         manager.IOManager  `json:"-"`
}

func NewTaxIncludedPriceJob(ioManager manager.IOManager, taxRate float64) *TaxIncludedPriceJob {
	return &TaxIncludedPriceJob{
		TaxRate:           taxRate,
		InputPrices:       []float64{},
		TaxIncludedPrices: map[string]float64{},
		IOManager:         ioManager,
	}
}

func (taxIncludedPriceJob *TaxIncludedPriceJob) Process() error {
	err := taxIncludedPriceJob.IOManager.ReadJSON(taxIncludedPriceJob)
	if err != nil {
		return err
	}
	for _, price := range taxIncludedPriceJob.InputPrices {
		taxIncludedPriceJob.TaxIncludedPrices[fmt.Sprintf("%.2f", price)] = price * (1 + taxIncludedPriceJob.TaxRate)
	}
	return taxIncludedPriceJob.IOManager.WriteJSON(taxIncludedPriceJob)
}
