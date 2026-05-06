package main

import (
	"fmt"
	"os"
)

func ProfitCalculator() {
	var revenue float64
	var expenses float64
	var taxRate float64
	fmt.Print("Enter Revenue: ")
	fmt.Scan(&revenue)
	fmt.Print("Enter Expenses: ")
	fmt.Scan(&expenses)
	fmt.Print("Enter Tax Rate: ")
	fmt.Scan(&taxRate)
	ebt, profit, ratio := calculateFinancials(revenue, expenses, taxRate)
	fmt.Println("Ratio =", ratio)
	fmt.Println("Profit =", profit)
	fmt.Println("EBT =", ebt)
	writeCalculatedToFile(ebt, profit, ratio)
}

func calculateFinancials(revenue float64, expenses float64, taxRate float64) (ebt float64, profit float64, ratio float64) {
	ebt = revenue - expenses
	profit = ebt * (1 - taxRate/100)
	ratio = ebt / profit
	return ebt, profit, ratio
}
func writeCalculatedToFile(ebt, profit, ratio float64) {
	dataText := fmt.Sprintf("EBT: %.1f\nProfit: %.1f\nRatio: %.1f", ebt, profit, ratio)
	os.WriteFile("investmentCalculator.txt", []byte(dataText), 0644)
}
