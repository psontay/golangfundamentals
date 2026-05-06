package main

import (
	"fmt"
)

func getUserInput(prompt string) string {
	fmt.Print(prompt)
	var data string
	fmt.Scan(&data)
	return data
}
