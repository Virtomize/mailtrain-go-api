package gomailtrain

import (
	"net/http"
	"net/url"
)

// API is the main api data structure
type API struct {
	endPoint *url.URL
	client   *http.Client
	token    string
	debug    bool
}
