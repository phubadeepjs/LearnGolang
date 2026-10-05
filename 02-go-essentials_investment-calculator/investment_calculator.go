package main

import (
	"fmt"
	"math"
)

const inflationRate float64 = 2.5

func main() {
	var investmentAmount float64
	expectedReturnRate := 5.5
	var years float64

	investmentAmount = userInput("Investment Amount: ")
	// fmt.Print("Investment Amount: ")
	// fmt.Scan(&investmentAmount)
	expectedReturnRate = userInput("Exoected Reture Rate: ")
	// fmt.Print("Exoected Reture Rate: ")
	// fmt.Scan(&expectedReturnRate)
	years = userInput("Years: ")
	// fmt.Print("Years: ")
	// fmt.Scan(&years)

	futureValue, futureRealValue := calculateFutureValue(investmentAmount, expectedReturnRate, years)
	// futureValue := investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	// futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)

	// formattedFV := fmt.Sprintf("Future Value: %.2f\n", futureValue)
	// formattedRFV := fmt.Sprintf("Future Value (Adjusted for inflation): %.2f\n", futureRealValue)
	// outputs information
	// fmt.Println("Future Value:", futureValue)
	// fmt.Println("Future Value (Adjusted for inflation):", futureRealValue)
	// fmt.Printf(`Future Value: %.2f
	// Future Value (Adjusted for inflation): %.2f`, futureValue, futureRealValue)
	showUserOutput(futureValue, futureRealValue)
}

func userInput(text string) (userInput float64) {
	fmt.Print(text)
	fmt.Scan(&userInput)
	return userInput
}

func calculateFutureValue(investmentAmount, expectedReturnRate, years float64) (fv float64, rfv float64) {
	fv = investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	rfv = fv / math.Pow(1+inflationRate/100, years)
	return fv, rfv
	// return
}

func showUserOutput(futureValue, futureRealValue float64) {
	fv := fmt.Sprintf("Future Value: %.2f\n", futureValue)
	rfv := fmt.Sprintf("Future Value (Adjusted for inflation): %.2f\n", futureRealValue)
	fmt.Print(fv, rfv)
}
