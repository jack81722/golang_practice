package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	go call(ctx, "A")
	go call(ctx, "B")
	go call(ctx, "C")

	time.Sleep(time.Second * 6)
}

func call(c context.Context, name string) {
	ctx, cancel := context.WithCancel(c)
	defer cancel()
	for {
		select {
		case <-ctx.Done():
			// wait
			fmt.Printf("%s end\n", name)
			return
		default:
			time.Sleep(time.Second)
			fmt.Printf("waiting %s ...\n", name)
		}
	}
}
