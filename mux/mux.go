package mux

import (
	"github.com/gorilla/mux"

	"github.com/middleware-labs/golang-apm/tracker"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gorilla/mux/otelmux"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	oteltrace "go.opentelemetry.io/otel/trace"
)

type AttributeOption []attribute.KeyValue

func Middleware(config *tracker.Config) mux.MiddlewareFunc {
	return otelmux.Middleware(config.ServiceName)
}

func CreateTracer(config *tracker.Config) oteltrace.Tracer {
	var tracer = otel.Tracer(config.ServiceName)
	return tracer
}

func WithAttributes(attributes attribute.KeyValue) oteltrace.SpanStartEventOption {
	return oteltrace.WithAttributes(attributes)
}
