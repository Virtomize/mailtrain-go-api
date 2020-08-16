package gomailtrain

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

// Subscription type
type Subscription struct {
	CID                 string `json:"-"`
	Email               string `json:"email"`
	Name                string `json:"merge_name,omitempty"`
	FirstName           string `json:"merge_first_name,omitempty"`
	LastName            string `json:"merge_last_name,omitempty"`
	Timezone            string `json:"timezone"`
	ForceSubscribe      bool   `json:"force_subscribe"`
	RequireConfirmation bool   `json:"require_confirmation"`
}

// SubscribeResponse type
type SubscribeResponse struct {
	Data SubscribeData `json:"data"`
}

// SubscribeData type
type SubscribeData struct {
	ID string `json:"id"`
}

// UnsubscribeResponse type
type UnsubscribeResponse struct {
	Data UnsubscribeData `json:"data"`
}

// UnsubscribeData type
type UnsubscribeData struct {
	ID           int  `json:"id"`
	Unsubscribed bool `json:"unsubscribed"`
}

// DeleteSubscriptionResponse type
type DeleteSubscriptionResponse struct {
	Data DeleteSubscriptionData `json:"data"`
}

// DeleteSubscriptionData type
type DeleteSubscriptionData struct {
	ID      int  `json:"id"`
	Deleted bool `json:"deleted"`
}

// Subscribe to a list
func (a *API) Subscribe(s Subscription) (*SubscribeResponse, error) {
	ep, err := url.ParseRequestURI(a.endPoint.String() + "/api/subscribe/" + s.CID)
	if err != nil {
		return nil, err
	}

	// validate timezone
	_, err = time.LoadLocation(s.Timezone)
	if err != nil {
		return nil, err
	}

	data := createData(s)

	req, err := http.NewRequest(http.MethodPost, ep.String(), strings.NewReader(data.Encode()))
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))

	res, err := a.Request(req)
	if err != nil {
		return nil, err
	}

	var resp SubscribeResponse
	err = json.Unmarshal(res, &resp)
	if err != nil {
		return nil, err
	}

	return &resp, nil
}

// Unsubscribe from a list
func (a *API) Unsubscribe(listCID string, email string) (*UnsubscribeResponse, error) {
	ep, err := url.ParseRequestURI(a.endPoint.String() + "/api/unsubscribe/" + listCID)
	if err != nil {
		return nil, err
	}

	if email == "" {
		return nil, fmt.Errorf("email is empty")
	}

	data := url.Values{}
	data.Set("email", email)

	req, err := http.NewRequest(http.MethodPost, ep.String(), strings.NewReader(data.Encode()))
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))

	res, err := a.Request(req)
	if err != nil {
		return nil, err
	}

	var resp UnsubscribeResponse
	err = json.Unmarshal(res, &resp)
	if err != nil {
		return nil, err
	}

	return &resp, nil
}

// DeleteSubscription .
func (a *API) DeleteSubscription(listCID string, email string) (*DeleteSubscriptionResponse, error) {
	ep, err := url.ParseRequestURI(a.endPoint.String() + "/api/delete/" + listCID)
	if err != nil {
		return nil, err
	}

	if email == "" {
		return nil, fmt.Errorf("email is empty")
	}

	data := url.Values{}
	data.Set("email", email)

	req, err := http.NewRequest(http.MethodPost, ep.String(), strings.NewReader(data.Encode()))
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))

	res, err := a.Request(req)
	if err != nil {
		return nil, err
	}

	var resp DeleteSubscriptionResponse
	err = json.Unmarshal(res, &resp)
	if err != nil {
		return nil, err
	}

	return &resp, nil
}
