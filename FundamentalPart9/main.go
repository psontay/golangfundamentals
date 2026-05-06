package main

import (
	"FundamentalPart9/prices"
	"encoding/json"
	"fmt"
	"os"
)


func main() {
	data, err := ReadFile("data.json")
	if err != nil {
		fmt.Println(err)
	}
	for _, taxRate := range
}
