package practice

import (
	"fmt"
	"time"
)

type Order struct {
	ID    int
	Price float64
}

func Worker(id int, jobs chan Order, results chan Order) {
	for job := range jobs {
		fmt.Printf("Worker %d received job %d\n", id, job.ID)
		time.Sleep(1 * time.Second)
		job.Price = job.Price + 10.0
		results <- job
	}
}

func SanXuat(chanOut chan int) {
	for i := 1; i <= 10; i++ {
		chanOut <- i
	}
	close(chanOut)
}

func LocChan(chanIn chan int, chanOut chan int) {
	for num := range chanIn {
		if num%2 == 0 {
			chanOut <- num
		}
	}
	close(chanOut)
}
