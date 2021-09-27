package http

import (
	"checkr.com/idempotent-requests/pkg/captures/captures_mock"
	"checkr.com/idempotent-requests/pkg/tracing"
	"checkr.com/idempotent-requests/pkg/views"
	"encoding/base64"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

const (
	PayloadIdempotencyKey  = "idempotency_key"
	PayloadResponse        = "response"
	PayloadResponseStatus  = "response_status"
	PayloadResponseBody    = "response_body"
	PayloadResponseHeaders = "response_headers"
)

func TestCaptures_Record(t *testing.T) {
	capturesRepo := captures_mock.NewRepositoryImpl()
	tracer := tracing.MockTracer{}

	router := NewRouter(tracer, capturesRepo)

	tests := []struct {
		name               string
		req                *http.Request
		reqPayload         interface{}
		expResCode         int
		expResBody         string
		allocationRequired bool
	}{
		{
			name:               "Captures_Record",
			reqPayload:         requestPayload(base64.URLEncoding.EncodeToString([]byte("everything-is-fine"))),
			expResCode:         http.StatusOK,
			allocationRequired: true,
		},
		{
			name:               "Captures_Record_Missing_Allocation",
			reqPayload:         requestPayload(base64.URLEncoding.EncodeToString([]byte("missing-allocation"))),
			expResCode:         http.StatusForbidden,
			allocationRequired: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			if tt.allocationRequired {
				allocate(t, router, tt.reqPayload)
			}

			tt.req = NewRecordRequest(tt.reqPayload)

			respRecorder := httptest.NewRecorder()
			router.ServeHTTP(respRecorder, tt.req)
			assert.Equal(t, tt.expResCode, respRecorder.Code)

			if tt.expResBody != "" {
				assert.Equal(t, tt.expResBody, respRecorder.Body.String())
			}
		})
	}

}

func requestPayload(idempotencyKey string) interface{} {
	headers := []views.ResponseHeader{
		{
			Key:   "Set-Cookie",
			Value: "secret",
		},
	}

	capture := &views.Capture{
		ResponseStatus:  http.StatusOK,
		ResponseBody:    "{\"resource_id\": \"123\"}",
		ResponseHeaders: headers,
	}
	record := views.CaptureRecord{
		IdempotencyKey: idempotencyKey,
		Response:       capture,
	}
	return record
}

func allocate(t *testing.T, router *gin.Engine, payload interface{}) {
	assert.IsType(t, views.CaptureRecord{}, payload)
	idempotencyKey := payload.(views.CaptureRecord).IdempotencyKey
	req := NewAllocationRequest(views.CaptureAllocation{IdempotencyKey: idempotencyKey})

	respRecorder := httptest.NewRecorder()
	router.ServeHTTP(respRecorder, req)
	assert.Equal(t, respRecorder.Code, http.StatusAccepted)
}

func NewRecordRequest(payload interface{}) *http.Request {
	return newRequest(http.MethodPost, ApiV2+Captures, reqBody(payload))
}
