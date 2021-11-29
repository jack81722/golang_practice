package main

import (
	"bufio"
	"firestoreexp/cache"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"

	"gopkg.in/yaml.v2"
)

type StoreConfig struct {
	ProjectId string
	JsonPath  string
}

func main() {
	data, err := ioutil.ReadFile("proj.yaml")
	if err != nil {
		log.Fatal(err.Error())
	}
	cfg := StoreConfig{}
	err = yaml.Unmarshal([]byte(data), &cfg)
	if err != nil {
		log.Fatal(err.Error())
	}

	c, err := cache.NewFireStore(cfg.ProjectId, cfg.JsonPath, "Cache", "Data")
	if err != nil {
		log.Fatal(err)
	}
	defer c.Close()
	// c.Reset()
	log.Println("start example")
	// c.Set("1", "data", 1)
	// c.Set("2", "exp_data", 0)
	// c.SetMarshal("date", time.Now(), 0)
	c.SetInt64("num", 123, 0)
	log.Println(c.Get("1"))
	log.Println(c.Get("2"))
	var t time.Time
	err = c.GetUnmarshal("date", &t)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(t)
	log.Println(c.GetInt64("num"))
	log.Println("1 exists? ", c.Exist("1"))
	log.Println("2 exists? ", c.Exist("2"))
	log.Println("3 exists? ", c.Exist("3"))

	// c.Set("3", "test_snapshot", 5)
	go func() {
		err := c.OnSnapshot("3", onSnapshot)
		if err != nil {
			return
		}
	}()
	reader := bufio.NewReader(os.Stdin)
	_, _ = reader.ReadString('\n')
	var ln string
	fmt.Scanf("%s", &ln)
}

func onSnapshot(data map[string]interface{}) {
	for k, v := range data {
		log.Println(k, v)
	}
}
