package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"

	"gopkg.in/yaml.v2"
)

type T struct {
	A string
	B int
}

const (
	USER_NAME = "USER_NAME"
	USER_PASS = "USER_PASS"
)

func main() {
	t := T{}
	data, err := ioutil.ReadFile("env.yaml")
	if err != nil {
		log.Fatal(err.Error())
		return
	}

	err = yaml.Unmarshal([]byte(data), &t)
	if err != nil {
		log.Fatal(err.Error())
		return
	}
	fmt.Printf("T: %s, %d\n", t.A, t.B)
	os.Setenv(USER_NAME, t.A)
	os.Setenv(USER_PASS, strconv.Itoa(t.B))
	fmt.Printf("%s:%v, %s:%v\n", USER_NAME, os.Getenv(USER_NAME), USER_PASS, os.Getenv(USER_PASS))

	values := os.Args[1:]
	for _, v := range values {
		fmt.Println(v)
	}
}
