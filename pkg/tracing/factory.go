package tracing

import (
	"go.uber.org/zap"
)

var tracerConstructors = map[string]func(enabled bool) Tracer{
	Datadog: NewDatadogTracer,
}

func NewTracer() Tracer {
	cfg := GetConfig()
	if _, ok := tracerConstructors[cfg.TracerImplementation]; !ok {
		zap.S().Panicf("not valid opentracing tracer implementation %s", cfg.TracerImplementation)
	}

	return tracerConstructors[cfg.TracerImplementation](cfg.TracingEnabled)
}
