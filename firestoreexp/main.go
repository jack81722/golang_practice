package main

import (
	"context"
	"io/ioutil"
	"log"

	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
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

	ctx := context.Background()
	conf := &firebase.Config{ProjectID: cfg.ProjectId}
	app, err := firebase.NewApp(ctx, conf, option.WithCredentialsFile(cfg.JsonPath))
	if err != nil {
		log.Fatal(err.Error())
	}
	client, err := app.Firestore(ctx)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer client.Close()

	// ========= do something on firestore ==========

	// _, _, err = client.Collection("users").Add(ctx, map[string]interface{}{
	// 	"first": "Ada",
	// 	"last":  "Lovelace",
	// 	"born":  1815,
	// })
	// if err != nil {
	// 	log.Fatal(err.Error())
	// }
	_, _, err = client.Collection("users").Add(ctx, map[string]interface{}{
		"first":    "Bob",
		"last":     "Lovelace",
		"born":     1820,
		"favorite": "apple",
	})
	if err != nil {
		log.Fatal(err.Error())
	}
}
