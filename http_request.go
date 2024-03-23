package rest

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/log"
	"moul.io/http2curl/v2"
)

// DefaultHttpRequester will simplify http request specific to this package's need
type DefaultHttpRequester struct {
	client HttpClient
	hook   []Hook
}

// Validates that current implementation is implement HttpRequester interface.
var _ HttpRequester = &DefaultHttpRequester{}

// DefaultClient will do http request using selected client.
// By using this, you can log http
func DefaultClient(client HttpClient, opts ...Option) (*DefaultHttpRequester, error) {
	if client == nil {
		panic("cannot use nil http client")
	}

	defaultClient := &DefaultHttpRequester{
		client: client,
		hook:   make([]Hook, 0),
	}

	for _, o := range opts {
		if err := o(defaultClient); err != nil {
			return nil, err
		}
	}

	return defaultClient, nil
}

func (r DefaultHttpRequester) beforeHook(ctx context.Context, data HookData) {
	for _, hook := range r.hook {
		if hook == nil {
			continue
		}

		hook.BeforeRequest(ctx, data)
	}
}

func (r DefaultHttpRequester) afterHook(ctx context.Context, data HookData) {
	for _, hook := range r.hook {
		if hook == nil {
			continue
		}

		hook.AfterRequest(ctx, data)
	}
}

func (r DefaultHttpRequester) Get(
	ctx context.Context,
	correlationID,
	path string,
	header http.Header,
) (ret HttpResponse, err error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "Get")
	defer func() {
		span.Finish()
		ctx.Done()
	}()

	ret, err = r.call(ctx, http.MethodGet, correlationID, path, header, nil)
	return
}

func (r DefaultHttpRequester) Post(
	ctx context.Context,
	correlationID,
	path string,
	requestHeader http.Header,
	requestBody []byte,
) (ret HttpResponse, err error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "Post")
	defer func() {
		span.Finish()
		ctx.Done()
	}()

	ret, err = r.call(ctx, http.MethodPost, correlationID, path, requestHeader, requestBody)
	return
}

func (r DefaultHttpRequester) Put(
	ctx context.Context,
	correlationID,
	path string,
	requestHeader http.Header,
	requestBody []byte,
) (ret HttpResponse, err error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "Put")
	defer func() {
		span.Finish()
		ctx.Done()
	}()

	ret, err = r.call(ctx, http.MethodPut, correlationID, path, requestHeader, requestBody)
	return
}

func (r *DefaultHttpRequester) Patch(ctx context.Context, correlationID, path string, requestHeader http.Header, requestBody []byte) (ret HttpResponse, err error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "Patch")
	defer func() {
		span.Finish()
		ctx.Done()
	}()

	ret, err = r.call(ctx, http.MethodPatch, correlationID, path, requestHeader, requestBody)
	return
}

func (r *DefaultHttpRequester) Delete(ctx context.Context, correlationID, path string, requestHeader http.Header, requestBody []byte) (ret HttpResponse, err error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "Delete")
	defer func() {
		span.Finish()
		ctx.Done()
	}()

	ret, err = r.call(ctx, http.MethodDelete, correlationID, path, requestHeader, requestBody)
	return
}

func (r DefaultHttpRequester) call(
	ctx context.Context,
	method,
	correlationID,
	path string,
	requestHeader http.Header,
	requestBody []byte,
) (ret HttpResponse, err error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "call")
	now := time.Now()
	request := &http.Request{}
	requestRaw := HttpRequest{}

	span.LogFields(
		log.String("method", method),
		log.String("path", path),
		log.Object("header", requestHeader),
	)

	requestHeader.Set(correlationIDKey, correlationID)

	defer func() {
		r.afterHook(ctx, HookData{
			Error:         err,
			URL:           path,
			CURL:          ret.CURL,
			StartTime:     now,
			Request:       requestRaw,
			Response:      ret.Raw,
			CorrelationID: correlationID,
		})

		span.Finish()
		ctx.Done()
	}()

	request.Method = method
	request.URL = &url.URL{}
	request.Header = requestHeader
	request.Body = ioutil.NopCloser(bytes.NewBuffer(requestBody))
	request.ContentLength = int64(len(requestBody))

	ret = HttpResponse{}
	ret.CURL = ""
	ret.Raw = ResponseRaw{}

	requestURL, err := url.Parse(path)
	if err != nil {
		err = fmt.Errorf("fail parse url %s: %s", path, err.Error())

		r.beforeHook(ctx, HookData{
			Error:         err,
			URL:           path,
			CURL:          ret.CURL,
			StartTime:     now,
			Request:       requestRaw,
			Response:      ResponseRaw{},
			CorrelationID: correlationID,
		})

		return ret, err
	}

	request.URL = requestURL
	request = request.WithContext(ctx)
	_ = span.Tracer().Inject(span.Context(), opentracing.HTTPHeaders, opentracing.HTTPHeadersCarrier(request.Header))

	if command, errCurl := http2curl.GetCurlCommand(request); errCurl == nil {
		ret.CURL = command.String()
	}

	var reqBodyInterface interface{}
	if err := json.Unmarshal(requestBody, &reqBodyInterface); err != nil {
		reqBodyInterface = string(requestBody)
	}

	requestRaw = HttpRequest{
		Method:           request.Method,
		URL:              requestURL,
		Proto:            request.Proto,
		ProtoMajor:       request.ProtoMajor,
		ProtoMinor:       request.ProtoMinor,
		Header:           request.Header,
		Body:             reqBodyInterface,
		ContentLength:    request.ContentLength,
		TransferEncoding: request.TransferEncoding,
		Close:            request.Close,
		Host:             request.Host,
		Form:             request.Form,
		PostForm:         request.PostForm,
		MultipartForm:    request.MultipartForm,
		Trailer:          request.Trailer,
		RemoteAddr:       request.RemoteAddr,
		RequestURI:       request.RequestURI,
		TLS:              request.TLS,
	}

	r.beforeHook(ctx, HookData{
		Error:         nil,
		URL:           path,
		CURL:          ret.CURL,
		StartTime:     now,
		Request:       requestRaw,
		Response:      ret.Raw,
		CorrelationID: correlationID,
	})

	span.LogFields(
		log.String("curl", ret.CURL),
	)

	resp, errHttp := r.client.Do(request)
	if resp == nil {
		if errHttp != nil {
			// https://github.com/golang/go/blob/2d77d3330537e11a0d9a233ba5f4facf262e9d8c/src/net/http/client.go#L724
			if strings.Contains(errHttp.Error(), "Client.Timeout") {
				err = ErrHttpTimeout
				return
			}

			err = fmt.Errorf("error response http.Do is nil, err http %s", errHttp.Error())
			return
		}
		err = fmt.Errorf("error response http.Do is nil")
		return
	}

	defer func() {
		if err := resp.Body.Close(); err != nil {
			span.LogFields(
				log.String("error_close", err.Error()),
			)
			return
		}
	}()

	ret.Raw.Status = resp.Status
	ret.Raw.StatusCode = resp.StatusCode
	ret.Raw.Proto = resp.Proto
	ret.Raw.ProtoMajor = resp.ProtoMajor
	ret.Raw.ProtoMinor = resp.ProtoMinor
	ret.Raw.Header = resp.Header
	ret.Raw.Body = nil
	ret.Raw.ContentLength = resp.ContentLength
	ret.Raw.TransferEncoding = resp.TransferEncoding
	ret.Raw.Uncompressed = resp.Uncompressed

	buf := new(bytes.Buffer)
	defer buf.Reset()
	_, err = buf.ReadFrom(resp.Body)
	if err != nil {
		err = fmt.Errorf("error read body response %s", err.Error())
		return
	}

	ret.Raw.Body = buf.String()
	var body interface{}
	if err := json.Unmarshal(buf.Bytes(), &body); err == nil {
		ret.Raw.Body = body
	}

	ret.RespBody = buf.Bytes()

	// Last handle of HTTP error
	if errHttp != nil {
		err = errHttp
	}

	return
}
