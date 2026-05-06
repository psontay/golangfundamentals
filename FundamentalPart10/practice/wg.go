package practice

import (
	"fmt"
	"sync"
	"time"
)

func DownFile(id int, fileName string, wg *sync.WaitGroup) { // must use wg by pointer
	defer wg.Done()
	fmt.Printf("ID: %d Downloading...: %s\n", id, fileName)
	time.Sleep(2 * time.Second)
	fmt.Printf("ID: %d Downloaded: %s\n", id, fileName)
}
