package tracing

import (
	"checkr.com/idempotent-requests/pkg/util"
	"go.uber.org/zap"
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/opentracer"
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer"
)

var tracerConstructors = map[string]func() Tracer{
	Datadog: NewDatadogTracer,
}

func NewTracer() Tracer {
	cfg := GetConfig()
	if _, ok := tracerConstructors[cfg.TracerImplementation]; !ok {
		zap.S().Panicf("not valid opentracing tracer implementation %s", cfg.TracerImplementation)
	}

	return tracerConstructors[cfg.TracerImplementation]()
}

func NewDatadogTracer() Tracer {
	dt := DatadogTracer{
		t: opentracer.New(tracer.WithServiceName(util.DDServiceName())),
	}
	return dt
}
