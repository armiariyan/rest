package rest

import (
	"errors"
)

// We use Jaeger go client,
// see that in https://github.com/jaegertracing/jaeger-client-go/blob/v2.22.1/tracer.go#L103-L104
// it use getDefaultHeadersConfig() https://github.com/jaegertracing/jaeger-client-go/blob/v2.22.1/header.go#L62
// The value of TraceContextHeaderName is uber-trace-id that will be used in Jaeger to propagate tracing context via HTTP call using
// http header. See https://github.com/jaegertracing/jaeger-client-go/blob/v2.22.1/constants.go#L56-L58
// Deprecated: don't use this anymore since this too specific with Jaeger Client. Use opentracing.Inject instead.
const httpHeaderSpanPropagatorKey = "Uber-Trace-Id"

// Correlation-ID is from gateway service
const correlationIDKey = "Correlation-ID"

var ErrHttpTimeout = errors.New("Client.Timeout exceeded while awaiting headers")
