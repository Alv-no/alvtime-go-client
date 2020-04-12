package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
)

// AlvtimeClient client struct
type AlvtimeClient struct {
	domain string
	httpClient
}

type httpClient interface {
	Do(req *http.Request) (*http.Response, error)
}

type requestError struct {
	Code  string `json:"code"`
	Error string `json:"error"`
}

// New is a helper function to create a *AlvtimeClient based on a domain string
func New(domain string) (*AlvtimeClient, error) {
	c := &AlvtimeClient{
		domain:     domain,
		httpClient: &http.Client{},
	}

	return c, nil
}

func (c *AlvtimeClient) newRequest(
    method string,
    baseURL *url.URL,
    body interface{},
) (
    *http.Request,
    error,
) {
    uri := baseURL.String()

    bytesBuffer, err := createBytesBuffer(body)
    if err != nil {
        return nil, err
    }

    req, err := http.NewRequest(method, uri, bytesBuffer)
    if err != nil {
        return nil, err
    }

    return req, nil
}

func (c *AlvtimeClient) do(req *http.Request) ([]byte, error) {
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	var byteArr []byte
	byteArr, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		var e requestError
		json.Unmarshal(byteArr, &e)
		err = errors.New(
			"\n\tRequest StatusCode: " + strconv.Itoa(resp.StatusCode) +
				", \n\tStatus: " + resp.Status,
		)
		return nil, err
	}

	return byteArr, nil
}

func createBytesBuffer(payload interface{}) (*bytes.Buffer, error) {
	b, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}
	return bytes.NewBuffer(b), nil
}
