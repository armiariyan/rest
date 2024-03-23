package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/armiariyan/rest"
	"github.com/opentracing/opentracing-go"
)

func main() {
	ctx := context.Background()
	tag := "[myservice]"
	correlationID := "abc"

	requestLogger := &restLogger{
		tag: tag,
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
}

func (r restLogger) BeforeRequest(ctx context.Context, data rest.HookData) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "restLogger.BeforeRequest")
	defer func() {
		span.Finish()
		ctx.Done()
	}()

	msg := fmt.Sprintf("%s - BeforeRequest", r.tag)
	d, _ := json.Marshal(data)
	fmt.Println(msg, string(d))
}

func (r restLogger) AfterRequest(ctx context.Context, data rest.HookData) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "restLogger.AfterRequest")
	defer func() {
		span.Finish()
		ctx.Done()
	}()

	msg := fmt.Sprintf("%s - AfterRequest", r.tag)
	d, _ := json.Marshal(data)
	fmt.Println(msg, string(d))
}
