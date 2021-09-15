package main

import (
	p "callerexp/panichandler"
	"fmt"
)

func createPanic() {
	var s *string
	fmt.Println(*s)
}

func main() {
	defer p.RecoverPanic()
	createPanic()
	// p.Iterator(3, 0)
}
