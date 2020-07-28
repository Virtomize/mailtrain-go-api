package gomailtrain

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

// Field type
type Field struct {
	ListID        int    `json:"-"`
	Name          string `json:"name"`
	Type          string `json:"type"`
	Group         string `json:"group"`
	GroupTemplate string `json:"group_template"`
	Visible       bool   `json:"visible"`
}

// AddField to a list
func (a *API) AddField(f Field) error {
	ep, err := url.ParseRequestURI(a.endPoint.String() + "/api/field/" + strconv.Itoa(f.ListID))
	if err != nil {
		return err
	}

	if !isInList(validFieldTypes, f.Type) {
		return fmt.Errorf("%s is not a valid type", f.Type)
	}

	data := createData(f)

	req, err := http.NewRequest(http.MethodPost, ep.String(), strings.NewReader(data.Encode()))
	if err != nil {
		return err
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))

	_, err = a.Request(req)
	if err != nil {
		return err
	}

	return nil
}
