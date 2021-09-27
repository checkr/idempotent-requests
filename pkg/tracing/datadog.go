package tracing

import (
	"github.com/gin-gonic/gin"
	"github.com/opentracing/opentracing-go"
	"go.mongodb.org/mongo-driver/mongo/options"
	gintrace "gopkg.in/DataDog/dd-trace-go.v1/contrib/gin-gonic/gin"
	mongotrace "gopkg.in/DataDog/dd-trace-go.v1/contrib/go.mongodb.org/mongo-driver/mongo"
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/opentracer"
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer"
	"os"
)

type DatadogTracer struct {
	enabled bool
	t       opentracing.Tracer
}

func NewDatadogTracer(enabled bool) Tracer {
	dt := DatadogTracer{
		enabled: enabled,
		t:       opentracer.New(tracer.WithServiceName(DDServiceName())),
	}
	return dt
}

func (dt DatadogTracer) OpentracingTracer() opentracing.Tracer {
	return dt.t
}

func (DatadogTracer) Stop() {
	tracer.Stop()
}

func (dt DatadogTracer) MongoDB(clientOptions *options.ClientOptions) *options.ClientOptions {
	if dt.enabled {
		clientOptions.Monitor = mongotrace.NewMonitor(mongotrace.WithServiceName(DDServiceName()))
	}
	return clientOptions
}

func (dt DatadogTracer) Gin(router *gin.Engine) gin.IRoutes {
	if dt.enabled {
		router.Use(gintrace.Middleware(DDServiceName()))
	}
	return router
}

func DDServiceName() string {
	s := os.Getenv("DD_SERVICE")
	if len(s) == 0 {
		s = "idempotent-requests"
	}
	return s
}
