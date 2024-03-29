package http

import (
	"checkr.com/idempotent-requests/pkg/captures"
	"checkr.com/idempotent-requests/pkg/tracing"
	"context"
	"github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"time"
)

type Server struct {
	Server *http.Server
}

func NewServer(tracer tracing.Tracer, records captures.Repository) *Server {
	router := NewRouter(tracer, records)

	srv := &http.Server{
		Addr:         ":8080",
		Handler:      router,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
	}
	return &Server{Server: srv}
}

func NewRouter(tracer tracing.Tracer, captures captures.Repository) *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	router := gin.New()

	tracer.Gin(router)
	router.Use(ginzap.Ginzap(zap.S().Desugar(), time.RFC822, false))
	router.Use(ginzap.RecoveryWithZap(zap.S().Desugar(), false))

	registerHandlers(router, captures)

	return router
}

func (s *Server) Start() {
	zap.S().Info("Starting HTTP server...")
	if err := s.Server.ListenAndServe(); err != nil {
		zap.S().Error(err)
	}
}

func (s *Server) Stop(ctx context.Context) {
	zap.S().Info("Stopping HTTP server...")
	if err := s.Server.Shutdown(ctx); err != nil {
		zap.S().Panic(err)
	}
}
