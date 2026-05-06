package main

import (
	"FundamentalPart7/practice"
	"fmt"
)

func main() {
	p1 := &practice.Product{
		Price:    100,
		Quantity: 1,
	}
	mapProduct := map[string]*practice.Product{
		"p1": p1,
	}
	err := practice.UpdateStock(mapProduct, "p2", 10)
	if err != nil {
		fmt.Println(err)
	}
	err = practice.UpdateStock(mapProduct, "p1", 10)
	p2 := &practice.Product{
		Price:    100,
		Quantity: -1000,
	}
	mapProduct["p2"] = p2
	err = practice.UpdateStock(mapProduct, "p2", 10)
	if err != nil {
		fmt.Println(err)
	}
}
