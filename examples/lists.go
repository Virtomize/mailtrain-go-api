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

	// read all subscribed lists for an email
	lists, err := api.GetListsByEmail("mail@example.com")
	if err != nil {
		log.Fatal(err)
	}

	for _, list := range lists.Data {
		fmt.Printf("%+v\n", list)
	}

	// read all lists of a namespace
	lists, err = api.GetListsByNamespace(1)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%+v\n", lists)

	// create a new list
	newList := gomailtrain.ListCreate{
		NamespaceID:        1,
		UnSubscribtionMode: gomailtrain.UnSubscribeOneStep,
		Name:               "mylist",
		Description:        "my list description",
		Homepage:           "https://example.com",
		ContactEmail:       "me@example.com",
		FieldWizard:        gomailtrain.FieldWizardFirstLastName,
		PublicSubscribe:    true,
		SendConfiguration:  1,
	}

	data, err := api.CreateList(newList)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%+v\n", data)

	// use the cid not the id to delete a list
	err = api.DeleteList(list[0].CID)
	if err != nil {
		log.Fatal(err)
	}
}
