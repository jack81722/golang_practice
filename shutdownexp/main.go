package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	handler := http.NewServeMux()
	handler.HandleFunc("/", HelloWorld)

	srv := http.Server{
		Addr:    ":8081",
		Handler: handler,
	}

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal(err.Error())
		}
	}()
	fmt.Println("Server start")

	<-sig
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)

	go func() {
		log.Printf("system call:%+v\n", <-sig)
		cancel()
	}()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal(err.Error())
	}
	log.Println("Server shutdonw")
	// err := <-quit
	// if err != nil {
	// 	fmt.Println("error")
	// 	fmt.Println(err.Error())
	// }
}

func HelloWorld(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("helloworld"))
}
