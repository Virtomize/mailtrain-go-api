/*
Package gomailtrain implementing mailtrain API

Simple example:

	api, err := gomailtrain.NewAPI("https://mailtrain.example.com", "token")
	if err != nil {
		// handle error
	}

	// read all subscribed lists for an email
	lists, err := api.GetListsByEmail("mail@example.com")
	if err != nil {
		// handle error
	}
*/
package gomailtrain
