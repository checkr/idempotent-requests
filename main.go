package main

import (
	"checkr.com/idempotent-requests/pkg/captures/captures_mongo"
	"checkr.com/idempotent-requests/pkg/http"
	"checkr.com/idempotent-requests/pkg/mongodb"
	"checkr.com/idempotent-requests/pkg/tracing"
	"context"
	"github.com/opentracing/opentracing-go"
	"go.uber.org/zap"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	logger, _ := zap.NewProduction()
	zap.ReplaceGlobals(logger)
	defer logger.Sync()

	var tracer tracing.Tracer
	if tracing.GetConfig().TracingEnabled {
		tracer = tracing.NewTracer()
		opentracing.SetGlobalTracer(tracer.OpentracingTracer())
		defer tracer.Stop()
	}

	mongo := mongodb.NewClient(tracer)
	capturesRepo := captures_mongo.NewRepository(mongo)

	httpServer := http.NewServer(tracer, capturesRepo)
	go httpServer.Start()

	waitForShutdownSignal()

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	mongo.Shutdown(ctx)
	httpServer.Stop(ctx)
}

func waitForShutdownSignal() {
	signalChan := make(chan os.Signal, 1)
	signal.Notify(
		signalChan,
		syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT,
	)
	<-signalChan
}
