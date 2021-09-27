package tracing

import (
	"github.com/opentracing/opentracing-go"
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer"
)

type DatadogTracer struct {
	t opentracing.Tracer
}

func (dt DatadogTracer) OpentracingTracer() opentracing.Tracer {
	return dt.t
}

func (DatadogTracer) Stop() {
	tracer.Stop()
}
