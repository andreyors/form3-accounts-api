package test

import (
	"bytes"
	"io/ioutil"
	"net/http"
)

type RoundTripFunc func(req *http.Request) *http.Response

func (f RoundTripFunc) RoundTrip(req *http.Request) (*http.Response, error) {
	return f(req), nil
}

func MockClient(code int, body []byte) *http.Client {
	return &http.Client{
		Transport: RoundTripFunc(func(request *http.Request) *http.Response {
			return &http.Response{
				StatusCode: code,
				Body:       ioutil.NopCloser(bytes.NewReader(body)),
				Header:     make(http.Header),
			}
		}),
	}
}
