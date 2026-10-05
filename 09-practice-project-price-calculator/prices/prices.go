package prices

import (
	"fmt"

	"github.com/phubadeepjs/practise/conversion"
	"github.com/phubadeepjs/practise/iomanager"
)

type TaxIncludedPriceJob struct {
	IOManager         iomanager.IOManager `json:"-"`
	TaxRate           float64             `json:"tax_rate"`
	InputPrices       []float64           `json:"input_prices"`
	TaxIncludedPrices map[string]string   `json:"tax_included_prices"`
}

func NewTaxIncludedPriceJob(iom iomanager.IOManager, taxRate float64) *TaxIncludedPriceJob {
	return &TaxIncludedPriceJob{
		IOManager:   iom,
		InputPrices: []float64{10, 20, 30},
		TaxRate:     taxRate,
	}
}

func (t TaxIncludedPriceJob) Process(doneChan chan bool, errChan chan error) {
	err := t.LoadData()

	// errChan <- errors.New("An error!")

	if err != nil {
		// return err
		errChan <- err
		return
	}

	result := make(map[string]string)

	for _, price := range t.InputPrices {
		taxIncludedPrice := price * (1 + t.TaxRate)
		result[fmt.Sprintf("%.2f", price)] = fmt.Sprintf("%.2f", taxIncludedPrice)
	}

	t.TaxIncludedPrices = result
	t.IOManager.WriteResult(t)
	doneChan <- true
}

func (t *TaxIncludedPriceJob) LoadData() error {
	lines, err := t.IOManager.ReadLines()
	if err != nil {
		fmt.Println(err)
		return err
	}

	prices, err := conversion.StringToFloat(lines)

	if err != nil {
		fmt.Println(err)
		return err
	}

	t.InputPrices = prices

	return nil
}
