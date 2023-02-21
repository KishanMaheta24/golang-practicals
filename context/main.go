package main

import (
	"context"
	"fmt"
	"time"
)

func printNum(ctx context.Context) context.Context {
	return context.WithValue(ctx, "asdfg", "123d5644")
}

func printVal(ctx context.Context) {
	val := ctx.Value("asdfg")

	for {
		select {
		case <-ctx.Done():
			fmt.Println("finished....")
			return
		default:
			fmt.Println("still going....  ")
			time.Sleep(time.Millisecond * 200)
		}

	}
	fmt.Println(val)
	fmt.Println(ctx.Err())

}

func main() {
	fmt.Println("Context Example")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	ctx = printNum(ctx)
	printVal(ctx)
}
