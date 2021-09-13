package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"

	"cloud.google.com/go/storage"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
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
	ctx := context.Background()
	client, err := storage.NewClient(ctx, option.WithCredentialsFile(cfg.JsonPath))
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()
	fmt.Println("Buckets:")
	it := client.Buckets(ctx, cfg.ProjectId)
	for {
		battrs, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(battrs.Name)
	}
}
