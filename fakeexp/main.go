package main

import (
	"fakeexp/fake"
	"fmt"
)

func main() {
	f := fake.NewFake()
	f.FakeReturn(f.Exp, func() string {
		return "HelloWorld"
	})
	fmt.Println(f.Exp())

	f.FakeReturn(f.Exp, func() string {
		return "Bye"
	})
	fmt.Println(f.Exp())
}
