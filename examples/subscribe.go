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
	subscription := gomailtrain.Subscription{
		CID:                 "Pv3POS1Wm",
		Email:               "test@foo.bar",
		FirstName:           "foo",
		LastName:            "bar",
		Timezone:            "UTC",
		ForceSubscribe:      true,
		RequireConfirmation: false,
	}

	sub, err := api.Subscribe(subscription)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(sub)

	unsub, err := api.Unsubscribe("Pv3POS1Wm", "test@foo.bar")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(unsub)

	delsub, err := api.DeleteSubscription("Pv3POS1Wm", "test@foo.bar")

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(delsub)

}
