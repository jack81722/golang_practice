package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	var q Queue
	q = &MyQueue{
		make([]interface{}, 0),
		0,
	}
	for i := 0; i < 5; i++ {
		q.Push(i)
	}
	for q.Len() > 0 {
		//fmt.Printf("len = %d, val = %d\n", q.Len(), q.Pop())
		fmt.Printf("len = %d, val = %d\n", q.Len(), q.PopTimeout(time.Second))
	}
}

type Queue interface {
	Push(obj interface{})
	Pop() interface{}
	Len() int
	PopTimeout(t time.Duration) interface{}
}

type MyQueue struct {
	array  []interface{}
	arrLen int
}

func (q *MyQueue) Push(obj interface{}) {
	q.array = append(q.array, obj)
	q.arrLen++
}

func (q *MyQueue) Pop() interface{} {
	result := q.array[0]
	q.array = q.array[1:]
	q.arrLen--
	return result
}

func (q *MyQueue) Len() int {
	return q.arrLen
}

func (q *MyQueue) PopTimeout(t time.Duration) interface{} {
	ctx, cancel := context.WithTimeout(context.Background(), t)
	defer cancel()
	result := make(chan interface{}, 1)

	go func() {
		time.Sleep(0 * time.Second)
		result <- q.Pop()
	}()

	select {
	case <-ctx.Done():
		panic("timeout")
	case v := <-result:
		return v
	}
}
