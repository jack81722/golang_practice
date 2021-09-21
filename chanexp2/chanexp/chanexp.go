package chanexp

import (
	"log"
	"runtime"
	"time"
)

type Vector []float64

func (v Vector) DoSomthing(i, n int, u Vector, c chan int) {
	for ; i < n; i++ {
		time.Sleep(time.Second * 5)
		log.Printf("[%d] run", i)
		v[i] += u[i] + v[i]
	}
	c <- 1
}

var numCPU = runtime.NumCPU()

func (v Vector) DoAll(u Vector) {
	c := make(chan int, numCPU)
	for i := 0; i < numCPU; i++ {
		go v.DoSomthing(i*len(v)/numCPU, (i+1)*len(v)/numCPU, u, c)
	}
	for i := 0; i < numCPU; i++ {
		<-c
	}
}
