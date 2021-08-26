package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	// client, err := DB.Connect()
	// if err != nil{ // TODO }
	// defer client.Close()

	// TODO : something
	// if err != nil {
	//    log.Println(err.Error())
	//	  return
	//}

	// return

	defer World(0)
	defer World(1)
	defer World(2)
}

func World(num int) {
	fmt.Printf("World %v\n", num)
}

func Test() {
	ch := make(chan string, 3)
	defer close(ch)

	fmt.Println("Start")
	go func() {
		strs := []string{
			"Hello",
			"World",
			"!",
		}
		for _, str := range strs {
			ch <- str
			fmt.Print(str + str + "\n")
		}
	}()

	fmt.Println("Waiting ...")
	time.Sleep(5 * time.Second)
	result := make([]string, 3)
	for i := 0; i < 3; i++ {
		result[i] = <-ch
		fmt.Println(result[i])
	}
}

func WGTest() {
	size := 3
	wg := sync.WaitGroup{}
	wg.Add(size)

	for i := 0; i < size; i++ {
		go func(num int) {
			time.Sleep(3 * time.Second)
			fmt.Println(num)
			wg.Done()
		}(i)
	}

	wg.Wait()
	fmt.Println("Finish")
}

func CtxTest() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	size := 3
	wg := sync.WaitGroup{}
	wg.Add(size)

	for i := 0; i < size; i++ {
		go func(num int) {
			fmt.Printf("Start %v\n", num)
			time.Sleep(time.Duration(num) * time.Second)
			fmt.Printf("End %v\n", num)
			wg.Done()
		}(i)
	}

	ch := make(chan int, 1)
	defer close(ch)
	go func() {
		wg.Wait()
		ch <- 0
	}()

	select {
	case <-ctx.Done():
		fmt.Println("Timeout")
	case <-ch:
		fmt.Println("Finish")
	}
}
