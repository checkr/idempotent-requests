package main

import (
	"checkr.com/idempotent-requests/pkg/captures/captures_mongo"
	"checkr.com/idempotent-requests/pkg/http"
	mongodb "checkr.com/idempotent-requests/pkg/mongodb"
	"context"
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

	mongo := mongodb.NewClient("mongodb://root:password123@localhost:27017")
	capturesRepo := captures_mongo.NewRepository(mongo)

	httpServer := http.NewServer(capturesRepo)
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
