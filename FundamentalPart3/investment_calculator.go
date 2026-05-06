package main

import (
	"fmt"
	"math"
)

const inflationRate = 2.5

func InvestmentCalculator() {
	var investmentAmount float64
	var years float64
	expectedReturnRate := 5.5
	fmt.Print("Enter investment amount: ")
	fmt.Scan(&investmentAmount)
	fmt.Print("Enter years: ")
	fmt.Scan(&years)
	fmt.Println("Enter expected return: ")
	fmt.Scan(&expectedReturnRate)

	futureValue, futureRealValue := calculateFutureValues(investmentAmount, years, expectedReturnRate)

	fmt.Printf("Future Real Value = %.2f\n", futureRealValue)
	fmt.Printf("Future Value = %.2f\n", futureValue)
}

func calculateFutureValues(years float64, investmentAmount float64, expectedReturnRate float64) (fv float64, rfv float64) {
	fv = investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	rfv = fv / math.Pow(1+inflationRate/100, years)
	return fv, rfv
}
