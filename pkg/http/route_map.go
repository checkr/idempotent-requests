package http

import (
	"checkr.com/idempotent-requests/pkg/captures"
	"net/http"
)

const Captures = "/captures"

func initApiV2Routes(repo captures.Repository) []Route {
	return Routes{
		{
			"AllocateCapture",
			http.MethodPut,
			Captures,
			NewAllocateCaptureHandler(repo),
		},
		{
			"RecordCapture",
			http.MethodPost,
			Captures,
			NewRecordCaptureHandler(repo),
		},
	}
}
func initApiMetaRoutes(repo captures.Repository) []Route {
	return Routes{
		{
			"Ping",
			http.MethodGet,
			"/ping",
			NewPingHandler(),
		},
		{
			"Panic",
			http.MethodGet,
			"/panic",
			NewPanicHandler(),
		},
		{
			"Readiness",
			http.MethodGet,
			"/ready",
			NewReadinessHandler(repo),
		},
	}
}
