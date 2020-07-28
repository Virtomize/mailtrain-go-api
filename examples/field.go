package main

import (
	"log"

	"github.com/virtomize/mailtrain-go-api"
)

func main() {

	api, err := gomailtrain.NewAPI("https://mailtrain.example.com", "token")
	if err != nil {
		log.Fatal(err)
	}

	//gomailtrain.SetDebug(true)

	f := gomailtrain.Field{
		ListID: 8,
		Name:   "test",
		Type:   "text",
	}

	err = api.AddField(f)
	if err != nil {
		log.Fatal(err)
	}
}
