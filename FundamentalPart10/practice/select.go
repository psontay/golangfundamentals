package practice

import (
	"math/rand/v2"
	"time"
)

func CallDatabase(ans chan string) {
	time.Sleep(time.Second * 5)
	ans <- "get JSON successfully"
}

func CallServerA(ch chan string) {
	sleepTime := rand.Int64N(3)
	time.Sleep(time.Duration(sleepTime) * time.Second)
	ch <- "Server A"
}

func CallServerB(ch chan string) {
	sleepTime := rand.Int64N(3)
	time.Sleep(time.Duration(sleepTime) * time.Second)
	ch <- "Server B"
}
