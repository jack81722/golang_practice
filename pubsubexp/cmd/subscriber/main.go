package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"pubsubexp/prepare"
	"syscall"
)

func main() {
	quitAppSignal := make(chan os.Signal, 1)
	p := prepare.Prepare{}
	err := p.DoAll()
	if err != nil {
		log.Fatal(err.Error())
	}
	defer p.Close()
	fmt.Print("start")
	signal.Notify(quitAppSignal, syscall.SIGINT, os.Interrupt, syscall.SIGTERM, syscall.SIGQUIT)
	<-quitAppSignal
}
