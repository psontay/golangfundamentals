package practice

import (
	"context"
	"fmt"
	"time"
)

func WheyDoer(ctx context.Context) {
	fmt.Println("Still doing whey shake..")
	select {
	case <-time.After(5 * time.Second):
		fmt.Println("Done whey shake")
	case <-ctx.Done():
		fmt.Println("Cancel whey shake")
	}
}

func QueryDatabase(ctx context.Context) {
	fmt.Println("Still querying..")
	select {
	case <-time.After(5 * time.Second):
		fmt.Println("Done querying")
	case <-ctx.Done():
		fmt.Println("Cancel querying")
	}
}
