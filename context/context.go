package main

import (
	"context"
	"fmt"
	"time"
)

func enrichContext(ctx context.Context) context.Context {
	return context.WithValue(ctx, "request_id", "12345")
}

func doSomething(ctx context.Context) {
	rID := ctx.Value("request_id")
	for {
		select {
		case <-ctx.Done():
			fmt.Println("time out")
			return
		default:
			fmt.Println("doing something")
		}
		time.Sleep(500 * time.Millisecond)
	}
	fmt.Println(rID)
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	ctx = enrichContext(ctx)
	go doSomething(ctx)
	select {
	case <-ctx.Done():
		fmt.Println("Timeout")
	}
	time.Sleep(2 * time.Second)
}
