package rest

import (
	"context"
	"net/http"
)

// AuthVirgoHttpRequester is a service for http client request
type HttpRequester interface {
	Get(ctx context.Context, correlationID, path string, header http.Header) (ret HttpResponse, err error)
	Post(ctx context.Context, correlationID, path string, requestHeader http.Header, requestBody []byte) (ret HttpResponse, err error)
	Put(ctx context.Context, correlationID, path string, requestHeader http.Header, requestBody []byte) (ret HttpResponse, err error)
	Patch(ctx context.Context, correlationID, path string, requestHeader http.Header, requestBody []byte) (ret HttpResponse, err error)
	Delete(ctx context.Context, correlationID, path string, requestHeader http.Header, requestBody []byte) (ret HttpResponse, err error)
}

type HttpClient interface {
	Do(req *http.Request) (*http.Response, error)
}
