package gomailtrain

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

// Lists type
type Lists struct {
	Data []List `json:"data"`
}

// List type
type List struct {
	ID   int    `json:"id"`
	CID  string `json:"cid"`
	Name string `json:"name"`
}

// ListCreate type
type ListCreate struct {
	NamespaceID             int    `json:"namespace"`
	UnSubscribtionMode      int    `json:"unsubscription_mode"`
	Name                    string `json:"name,omitempty"`
	Description             string `json:"description,omitempty"`
	Homepage                string `json:"homepage,omitempty"`
	ContactEmail            string `json:"contact_email,omitempty"`
	FieldWizard             string `json:"fieldwizard,omitempty"`
	ToName                  string `json:"to_name,omitempty"`
	ListUnSubscribeDisabled bool   `json:"listunsubscribe_disabled,omitempty"`
	PublicSubscribe         bool   `json:"public_subscribe,omitempty"`
	SendConfiguration       int    `json:"send_configuration,omitempty"`
}

// CreateListResponse type
type CreateListResponse struct {
	Data CreateListData `json:"data"`
}

// CreateListData type
type CreateListData struct {
	ID string `json:"id"`
}

// GetListsByEmail returns all subscribed lists for an email address
func (a *API) GetListsByEmail(email string) (*Lists, error) {
	ep, err := url.ParseRequestURI(a.endPoint.String() + "/api/lists/" + email)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodGet, ep.String(), nil)
	if err != nil {
		return nil, err
	}

	res, err := a.Request(req)
	if err != nil {
		return nil, err
	}

	var lists Lists
	err = json.Unmarshal(res, &lists)
	if err != nil {
		return nil, err
	}

	return &lists, nil
}

// GetListsByNamespace returns all lists of a namespaceID
func (a *API) GetListsByNamespace(namespaceID int) (*Lists, error) {
	ep, err := url.ParseRequestURI(a.endPoint.String() + "/api/lists-by-namespace/" + strconv.Itoa(namespaceID))
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodGet, ep.String(), nil)
	if err != nil {
		return nil, err
	}

	res, err := a.Request(req)
	if err != nil {
		return nil, err
	}

	var lists Lists
	err = json.Unmarshal(res, &lists)
	if err != nil {
		return nil, err
	}

	return &lists, nil
}

// CreateList creates a new list using a ListCreate object
func (a *API) CreateList(list ListCreate) (*CreateListResponse, error) {
	// root namespace has id 1
	if list.NamespaceID == 0 {
		return nil, fmt.Errorf("no namespace set")
	}

	// no need to check UnSubscribtionMode since 0 is a valid option
	ep, err := url.ParseRequestURI(a.endPoint.String() + "/api/list")
	if err != nil {
		return nil, err
	}

	// use the default configured system setting if nothing is set
	if list.SendConfiguration == 0 {
		list.SendConfiguration = 1
	}

	// we add all data by reflecting the list
	data := createData(list)

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

	var resp CreateListResponse
	err = json.Unmarshal(res, &resp)
	if err != nil {
		return nil, err
	}

	return &resp, nil
}

// DeleteList deletes a list by its cid
func (a *API) DeleteList(listCID string) error {
	ep, err := url.ParseRequestURI(a.endPoint.String() + "/api/list/" + listCID)
	if err != nil {
		return err
	}

	req, err := http.NewRequest(http.MethodDelete, ep.String(), nil)
	if err != nil {
		return err
	}

	res, err := a.Request(req)
	if err != nil {
		return err
	}

	fmt.Printf("%+v\n", res)

	return nil

}
