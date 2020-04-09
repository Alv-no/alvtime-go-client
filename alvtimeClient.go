package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"
)

type AlvtimeClient struct {
	domain string
	httpClient
}

type httpClient interface {
	Do(req *http.Request) (*http.Response, error)
}

type payload struct{}

func (p payload) bytesBuffer() *bytes.Buffer {
    b, _ := json.Marshal(p)
    return bytes.NewBuffer(b)
}

type requestError struct {
	Code  string `json:"code"`
	Error string `json:"error"`
}

func (alvtimeClient *AlvtimeClient) do(req *http.Request) ([]byte, error) {
	resp, err := alvtimeClient.httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	var byteArr []byte
	byteArr, err = ioutil.ReadAll(resp.Body)
	resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		var e requestError
		json.Unmarshal(byteArr, &e)
		err = errors.New(
			"\n\tRequest StatusCode: " + strconv.Itoa(resp.StatusCode) +
				", \n\tCODE: " + e.Code +
				", \n\tMESSAGE: " + e.Error,
		)
		return nil, err
	}

	return byteArr, nil
}
