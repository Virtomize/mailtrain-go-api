package gomailtrain

import (
	"encoding/json"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

// BlacklistResponse type
type BlacklistResponse struct {
	Data BlacklistData `json:"data"`
}

// BlacklistData type
type BlacklistData struct {
	Start  int      `json:"start"`
	Limit  int      `json:"limit"`
	Emails []string `json:"emails"`
}

// GetBlacklistMails returns blacklisted mail addresses
func (a *API) GetBlacklistMails(start int, limit int, search string) (*BlacklistResponse, error) {
	ep, err := url.ParseRequestURI(a.endPoint.String() + "/api/blacklist/get")
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodGet, ep.String(), nil)
	if err != nil {
		return nil, err
	}

	q := req.URL.Query()
	q.Add("start", strconv.Itoa(start))
	q.Add("limit", strconv.Itoa(limit))
	q.Add("search", search)
	req.URL.RawQuery = q.Encode()

	res, err := a.Request(req)
	if err != nil {
		return nil, err
	}

	var mails BlacklistResponse
	err = json.Unmarshal(res, &mails)
	if err != nil {
		return nil, err
	}

	return &mails, nil
}

// AddMailToBlacklist adds a mail to the blacklist
func (a *API) AddMailToBlacklist(mail string) error {
	ep, err := url.ParseRequestURI(a.endPoint.String() + "/api/blacklist/add")
	if err != nil {
		return err
	}

	data := url.Values{}
	data.Set("email", mail)

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

// DeleteMailFromBlacklist adds a mail to the blacklist
func (a *API) DeleteMailFromBlacklist(mail string) error {
	ep, err := url.ParseRequestURI(a.endPoint.String() + "/api/blacklist/delete")
	if err != nil {
		return err
	}

	data := url.Values{}
	data.Set("email", mail)

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
