package rest

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/opentracing/opentracing-go"
)

// HttpResponse add layer on top of http.Response
type HttpResponse struct {
	RespBody []byte
	CURL     string
	Raw      ResponseRaw
}

type ResponseDecoder func(data []byte, v interface{}) error

// To unmarshal body to struct using ResponseDecoder
func (h *HttpResponse) To(ctx context.Context, decoder ResponseDecoder, out interface{}) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "HttpResponse.To")
	defer func() {
		ctx.Done()
		span.Finish()
	}()

	if h == nil {
		return fmt.Errorf("h is nil")
	}

	if h.RespBody == nil {
		return fmt.Errorf("response body is nil")
	}

	return decoder(h.RespBody, out)
}

// ToJson unmarshal body to struct using JSON
func (h *HttpResponse) ToJson(ctx context.Context, out interface{}) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "HttpResponse.ToJson")
	defer func() {
		ctx.Done()
		span.Finish()
	}()

	return h.To(ctx, json.Unmarshal, out)
}
