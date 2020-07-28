package gomailtrain

import "net/http"

// Auth using token
func (a *API) Auth(req *http.Request) {
	q := req.URL.Query()
	q.Add("access_token", a.token)
	req.URL.RawQuery = q.Encode()
}
