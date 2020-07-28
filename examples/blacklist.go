package main

import (
	"fmt"
	"log"

	"github.com/virtomize/mailtrain-go-api"
)

func main() {

	api, err := gomailtrain.NewAPI("https://mailtrain.example.com", "token")
	if err != nil {
		log.Fatal(err)
	}

	//gomailtrain.SetDebug(true)

	// start, limit, search pattern
	mails, err := api.GetBlacklistMails(0, 10000, "")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%+v\n", mails)

	err = api.AddMailToBlackList("test@example.com")
	if err != nil {
		log.Fatal(err)
	}

	err = api.DeleteMailFromBlackList("test@example.com")
	if err != nil {
		log.Fatal(err)
	}
}
