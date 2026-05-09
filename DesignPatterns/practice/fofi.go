package practice

import (
	"fmt"
	"math/rand/v2"
	"time"
)

func FakeCallingEcommerce(name string, ans chan string) {
	time.Sleep(time.Duration(rand.IntN(3)+1) * time.Second)
	ans <- fmt.Sprintf("Found product in %s with price: %d", name, rand.IntN(5)+20)
}

func ReadZaloMessage(chanOut chan string) {
	for i := range 3 {
		time.Sleep(1 * time.Second)
		chanOut <- fmt.Sprintf("Read Zalo Message from %d", i)
	}
}

func ReadFaceBookMessage(chanOut chan string) {
	for i := range 3 {
		time.Sleep(1 * time.Second)
		chanOut <- fmt.Sprintf("Read FaceBook Message from %d", i)
	}
}
