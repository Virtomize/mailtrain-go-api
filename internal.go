package gomailtrain

import (
	"crypto/tls"
	"errors"
	"log"
	"net/http"
	"net/url"
	"reflect"
	"strconv"
	"strings"
)

// NewAPI constructor
func NewAPI(uri string, token string) (*API, error) {

	if len(uri) == 0 || len(token) == 0 {
		return nil, errors.New("url or token not set")
	}

	u, err := url.ParseRequestURI(uri)
	if err != nil {
		return nil, err
	}

	a := new(API)
	a.endPoint = u
	a.token = token

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: false},
	}

	a.client = &http.Client{Transport: tr}

	return a, nil
}

// NewAPIWithClient create a new API instance using an existing HTTP client
func NewAPIWithClient(uri string, client *http.Client) (*API, error) {
	if len(uri) == 0 {
		return nil, errors.New("url not set")
	}

	u, err := url.ParseRequestURI(uri)
	if err != nil {
		return nil, err
	}

	a := new(API)
	a.endPoint = u
	a.client = client

	return a, nil
}

// VerifyTLS to enable disable certificate checks
func (a *API) VerifyTLS(set bool) {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: !set},
	}
	a.client = &http.Client{Transport: tr}
}

// DebugFlag is the global debugging variable
var DebugFlag = false

// SetDebug enables debug output
func SetDebug(state bool) {
	DebugFlag = state
}

// Debug outputs debug messages
func Debug(msg interface{}) {
	if DebugFlag {
		log.Printf("%+v\n", msg)
	}
}

func createData(d interface{}) url.Values {
	data := url.Values{}

	typeOf := reflect.TypeOf(d)
	valueOf := reflect.ValueOf(d)

	for i := 0; i < valueOf.NumField(); i++ {
		pname := strings.Split(typeOf.Field(i).Tag.Get("json"), ",")[0]
		if pname == "-" {
			continue
		}
		switch typeOf.Field(i).Type.Kind() {
		case reflect.Bool:
			if isInList(brokenValues, pname) {
				if valueOf.Field(i).Bool() {
					data.Set(pname, "yes")
				}
			} else {
				b := "0"
				if valueOf.Field(i).Bool() {
					b = "1"
				}
				data.Set(pname, b)
			}
		case reflect.Int:
			data.Set(pname, strconv.FormatInt(valueOf.Field(i).Int(), 10))
		case reflect.String:
			data.Set(pname, valueOf.Field(i).String())
		}
	}

	return data
}

func isInList(list []string, e string) bool {
	for _, v := range list {
		if v == e {
			return true
		}
	}
	return false
}
