package main

import (
	"FundamentalPart10/practice"
	"fmt"
)

func main() {
	var orders []practice.Order
	for i := 1; i <= 5; i++ {
		orders = append(orders, practice.Order{
			ID:    i,
			Price: float64(10 + i),
		})
	}

	jobs := make(chan practice.Order, len(orders))
	results := make(chan practice.Order, len(orders))

	for i := 1; i < 4; i++ {
		go practice.Worker(i, jobs, results)
	}
	for _, o := range orders {
		jobs <- o
	}
	close(jobs)

	for i := 1; i <= len(orders); i++ {
		ans := <-results
		fmt.Printf("%d\t%f\n", ans.ID, ans.Price)
	}
	close(results)
}
