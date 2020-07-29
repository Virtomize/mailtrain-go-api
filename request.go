package gomailtrain

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

// Request implementation
func (a *API) Request(req *http.Request) ([]byte, error) {
	req.Header.Add("Accept", "application/json, */*")
	a.Auth(req)

	Debug("Request:")
	Debug(req)

	resp, err := a.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	Debug(fmt.Sprintf("Status Code: %d", resp.StatusCode))

	res, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	Debug("Response Body:")
	Debug(string(res))

	switch resp.StatusCode {
	case http.StatusOK, http.StatusCreated, http.StatusPartialContent:
		return res, nil
	case http.StatusNoContent, http.StatusResetContent:
		return nil, nil
	case http.StatusUnauthorized:
		return nil, fmt.Errorf("authentication failed")
	case http.StatusServiceUnavailable:
		return nil, fmt.Errorf("service is not available: %s", resp.Status)
	case http.StatusInternalServerError:
		return nil, fmt.Errorf("internal server error: %s", resp.Status)
	case http.StatusConflict:
		return nil, fmt.Errorf("conflict: %s", resp.Status)
	}

	return nil, fmt.Errorf("unknown response status: %s", resp.Status)

}
