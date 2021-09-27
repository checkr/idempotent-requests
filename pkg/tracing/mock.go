package tracing

import (
	"github.com/gin-gonic/gin"
	"github.com/opentracing/opentracing-go"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MockTracer struct{}

func (MockTracer) OpentracingTracer() opentracing.Tracer {
	return nil
}

func (MockTracer) Stop() {

}

func (MockTracer) MongoDB(options *options.ClientOptions) *options.ClientOptions {
	return options
}

func (MockTracer) Gin(engine *gin.Engine) gin.IRoutes {
	return engine
}
