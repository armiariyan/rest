package rest

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
)

// DoFuncHttpErr trigger HTTP call error
var DoFuncHttpErr = func(req *http.Request) (*http.Response, error) {
	body := ioutil.NopCloser(bytes.NewReader([]byte("hello world")))
	return &http.Response{
		StatusCode: http.StatusOK,
		Body:       body,
	}, fmt.Errorf("error")
}

// DoFuncErrParseBodyResp trigger unmarshal error
var DoFuncErrParseBodyResp = func(req *http.Request) (*http.Response, error) {
	body := ioutil.NopCloser(bytes.NewReader([]byte(`{`)))
	return &http.Response{
		StatusCode: http.StatusOK,
		Body:       body,
	}, nil
}
