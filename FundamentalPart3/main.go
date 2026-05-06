package main

import "fmt"

func main() {
	//ProfitCalculator()
	//InvestmentCalculator()
	//MiniBankApplication()
	age := 32
	var agePtr *int = &age
	fmt.Println("agePtr", agePtr)
	fmt.Println("age value", *agePtr)
	getAdultYears(&age)
	getAdultYears(agePtr)
}
