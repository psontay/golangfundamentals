package main

import (
	"FundamentalPart10/practice"
	"fmt"
	"time"
)

func main() {
	ketQua := make(chan string, 2)
	go practice.CallServerA(ketQua)
	go practice.CallServerB(ketQua)
	select {
	case <-ketQua:
		res := <-ketQua
		fmt.Println("ketQua", res)
	case <-time.After(time.Second * 2):
		fmt.Println("ketQua timeout")
	}
}
