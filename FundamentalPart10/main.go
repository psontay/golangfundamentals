package main

import (
	"FundamentalPart10/practice"
	"fmt"
	"sync"
)

func main() {
	box := practice.DonationBox{
		Balance: 0,
	}
	var wg sync.WaitGroup

	for i := range 100 {
		wg.Add(1)
		go box.Donate(i, &wg)
	}
	wg.Wait()
	fmt.Println("Box its full..., total:", box.Balance)
}
