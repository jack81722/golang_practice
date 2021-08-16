package main

import (
	"context"
	"log"
	"math/rand"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

const (
	workerSize = 1
)

func main() {
	consumer := &Consumer{
		ingestChan: make(chan int),
		jobsChan:   make(chan int, workerSize),
	}
	producer := &Producer{
		produce: consumer.assign,
	}

	go producer.Start()

	ctx, cancel := context.WithCancel(context.Background())
	go consumer.Start(ctx)
	wg := &sync.WaitGroup{}
	wg.Add(workerSize)

	for i := 0; i < workerSize; i++ {
		go consumer.work(i, wg)
	}

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutdown ...")
	cancel()
	wg.Wait()
	log.Println("End")
}

type Producer struct {
	produce func(event int)
}

func (p *Producer) Start() {
	eventIndex := 0
	for {
		log.Printf("Producer request(%d)\n", eventIndex)
		p.produce(eventIndex)
		log.Printf("Producer request(%d) succeed.\n", eventIndex)
		eventIndex++
		time.Sleep(time.Duration(100))
	}
}

type Consumer struct {
	ingestChan chan int
	jobsChan   chan int
}

func (c *Consumer) work(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	for eventIndex := range c.jobsChan {
		log.Printf("worker(%d) handle event(%d).\n", id, eventIndex)
		time.Sleep(time.Millisecond * time.Duration(rand.Intn(5000)+1000))
		log.Printf("worker(%d) job(%d) done.\n", id, eventIndex)
	}
}

func (c *Consumer) assign(event int) {
	c.ingestChan <- event
}

func (c *Consumer) Start(ctx context.Context) {
	// p.Produce -> ingestChan -> jobsChan
	// the event will be queued in channel
	// max queue is (workerSize + 1)
	for {
		select {
		case job := <-c.ingestChan:
			// accept the job and queue in jobs channel
			c.jobsChan <- job
		case <-ctx.Done():
			// ctx will done while cacel
			// close jobs channel guarantee the job won't handle
			close(c.jobsChan)
			log.Println("Close channel")
			return
		}
	}
}
