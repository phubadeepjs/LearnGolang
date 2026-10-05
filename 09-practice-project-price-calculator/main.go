package main

import (
	"fmt"

	"github.com/phubadeepjs/practise/filemanager"
	"github.com/phubadeepjs/practise/prices"
)

func main() {
	taxRates := []float64{0, 0.07, 0.1, 0.15}
	doneChans := make([]chan bool, len(taxRates))
	errChans := make([]chan error, len(taxRates))

	for index, taxRates := range taxRates {
		doneChans[index] = make(chan bool)
		errChans[index] = make(chan error)
		fm := filemanager.New("testdata/prices.txt", fmt.Sprintf("testdata/result_%.0f.json", taxRates*100))
		// cmdm := cmdmanager.New()
		priceJob := prices.NewTaxIncludedPriceJob(fm, taxRates)

		go priceJob.Process(doneChans[index], errChans[index])

		// if err != nil {
		// 	fmt.Println("Could not process job.")
		// 	fmt.Println(err)
		// }
	}

	for index := range taxRates {
		select {
		case err := <-errChans[index]:
			if err != nil {
				fmt.Println(err)
			}
		case <-doneChans[index]:
			fmt.Println("Done!")
		}
	}
}
