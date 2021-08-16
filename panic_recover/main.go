package main

import "log"

func main() {
	// define recover
	defer func() {
		log.Println(recover())
	}()
	// trigger panic
	work()

	log.Println("end")
}

func work() {
	panic("work")
}
