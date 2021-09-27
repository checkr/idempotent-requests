package tracing

import (
	"github.com/gin-gonic/gin"
	"github.com/opentracing/opentracing-go"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Tracer is a wrapping interface around opentracing.Tracer with extra methods to abstract the tracer implementation
type Tracer interface {
	// OpentracingTracer provides a tracer instance, which implements opentracing interface
	OpentracingTracer() opentracing.Tracer
	// Stop shuts the tracer down
	Stop()
	// MongoDB injects an out-of-the-box instrumentation for MongoDB Client
	MongoDB(*options.ClientOptions) *options.ClientOptions
	// Gin injects an out-of-the-box instrumentation for Gin server
	Gin(*gin.Engine) gin.IRoutes
}
