package main

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"
)

func main() {
	f()
}

func f() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	size := 5
	wg := &sync.WaitGroup{}
	wg.Add(size)
	// instatiate goroutine
	for i := 0; i < size; i++ {
		go func(index int) {
			defer wg.Done()
			time.Sleep(0 * time.Second)
			fmt.Printf("i = %d\n", index)
		}(i)
	}
	// goroutine to wait waitgroup
	ch := make(chan int, 1)
	go func() {
		wg.Wait()
		ch <- 0
	}()

	select {
	case <-ctx.Done():
		log.Println("timeout")
	case <-ch:
		log.Println("Finished")
	}
}
