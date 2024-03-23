# REST

A library to do HTTP Call.

## Features

* [x] Circuit breaker using [github.com/sony/gobreaker](github.com/sony/gobreaker)
* [x] Multiple read Body response
* [x] Hook Before and After request for logging purpose
* [ ] Retry
* [ ] httptrace https://golang.org/pkg/net/http/httptrace/

```go
package main

import (
	"context"
	"log"
	"fmt"
	"net/http"
    
	"github.com/opentracing/opentracing-go"
	"gitlab.com/gobang/logger"
	"gitlab.com/gobang/rest"
)

func main() {
	ctx := context.Background()
	tag := "[myservice]"
	correlationID := "abc"

	linkAjaLogger := logger.Noop()

    requestLogger := &restLogger{
    	tag: tag,
    	log: linkAjaLogger,
	}


	client := &http.Client{}
	httpClient, err := rest.DefaultClient(client, rest.AddHook(requestLogger))
	if err != nil {
		log.Fatal(err)
		return
	}

	resp, err := httpClient.Get(ctx, correlationID, "http://example.com", http.Header{})
	if err != nil {
		log.Fatal(err)
		return
	}

	log.Println(resp.Raw.Body)
}


type restLogger struct {
	tag string
	log logger.Logger
}

func (r restLogger) BeforeRequest(ctx context.Context, data rest.HookData) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "restLogger.BeforeRequest")
	defer func() {
		span.Finish()
		ctx.Done()
	}()

	msg := fmt.Sprintf("%s - BeforeRequest", r.tag)
	r.log.DebugWithMetadata(data.CorrelationID, r.tag, data.Request.Method, data.URL,
		logger.Field{
			Key: "message",
			Val: msg,
		},
		logger.Field{
			Key: "data",
			Val: data,
		},
	)
}

func (r restLogger) AfterRequest(ctx context.Context, data rest.HookData) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "restLogger.AfterRequest")
	defer func() {
		span.Finish()
		ctx.Done()
	}()

	msg := fmt.Sprintf("%s - AfterRequest", r.tag)
	r.log.DebugWithMetadata(data.CorrelationID, r.tag, data.Request.Method, data.URL,
		logger.Field{
			Key: "message",
			Val: msg,
		},
		logger.Field{
			Key: "data",
			Val: data,
		},
	)
}
```
