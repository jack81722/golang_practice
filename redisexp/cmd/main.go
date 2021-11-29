package main

import (
	"log"
	"redisexp/cache"
	"time"
)

func main() {
	store := cache.NewRedisStore()
	store.Set("exp_key", "exp_value", 3)
	value, err := store.Get("exp_key")
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println(value)
	}
	time.Sleep(4 * time.Second)
	value, err = store.Get("exp_key")
	if err != nil {
		log.Fatal(err)
	}
	log.Println(value)
}
