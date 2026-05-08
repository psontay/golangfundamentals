package main

import (
	"FundamentalPart10/practice"
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()
	go practice.QueryDatabase(ctx)
	time.Sleep(time.Second * 2)
	fmt.Println("Calling practice with the database!")
	time.Sleep(time.Second * 2)
}
