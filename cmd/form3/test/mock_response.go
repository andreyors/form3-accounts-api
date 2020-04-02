package test

import (
	"bytes"
	"io/ioutil"
	"net/http"
)

func MockResponse(name string) *http.Response {
	body := Sample(name)

	return &http.Response{
		Body: ioutil.NopCloser(bytes.NewReader(body)),
	}
}
