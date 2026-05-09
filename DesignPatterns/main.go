package main

import (
	"DesignPatterns/practice"
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	for i := 0; i < 500; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			d := practice.GetDatabase()
			fmt.Printf("Routine %d get manager %s\n", id, d.Config.Host)
		}(i)
	}
	wg.Wait()
}
