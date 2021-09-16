package http

import (
	"checkr.com/idempotent-requests/pkg/captures"
	"net/http"
)

const Captures = "/captures"

func initRoutes(repo captures.Repository) []Route {
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
