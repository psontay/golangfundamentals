package main

import (
	"FundamentalPart3/fileops"
	"errors"
	"fmt"
)

const accountBalanceFileName = "balance.txt"

func MiniBankApplication() {
	balance, err := fileops.GetFloatFromFile(accountBalanceFileName)
	if err != nil {
		fmt.Println("Error getting balance")
		fmt.Println("Set balance to default $0.00")
		balance = 0
	}
	for {
		PresentOptions()
		var choice int
		fmt.Scan(&choice)
		if choice == 4 {
			fmt.Println("See u soon")
			break
		}
		switch choice {
		case 1:
			fmt.Printf("Your balance is $%.2f\n", balance)
		case 2:
			var depositAmount float64
			fmt.Printf("Enter the amount to deposit:")
			fmt.Scan(&depositAmount)
			newBalance, err := DepositCalculator(depositAmount, balance)
			if err != nil {
				fmt.Println("Error:", err)
			} else {
				balance = newBalance
				fileops.WriteFloatToFile(balance, accountBalanceFileName)
				fmt.Println("Success!")
			}
		case 3:
			var withdrawAmount float64
			fmt.Printf("Enter the amount to withdraw:")
			fmt.Scan(&withdrawAmount)
			newBalance, err := WithdrawCalculator(withdrawAmount, balance)
			if err != nil {
				fmt.Println("Error:", err)
			} else {
				balance = newBalance
				fileops.WriteFloatToFile(balance, accountBalanceFileName)
				fmt.Println("Success!")
			}
		default:
			fmt.Println("Invalid choice")
		}
	}
}
func DepositCalculator(amount, balance float64) (float64, error) {
	if amount <= 0 {
		return balance, errors.New("invalid amount")
	}
	balance += amount
	return balance, nil
}
func WithdrawCalculator(amount, balance float64) (float64, error) {
	if amount > balance {
		return balance, errors.New("insufficient balance")
	}
	balance -= amount
	return balance, nil
}
