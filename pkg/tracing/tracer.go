package tracing

import "github.com/opentracing/opentracing-go"

// Tracer is a wrapping interface around opentracing.Tracer, which adds Stop method
type Tracer interface {
	OpentracingTracer() opentracing.Tracer
	Stop()
}
