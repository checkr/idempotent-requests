package http

import (
	"bytes"
	"checkr.com/idempotent-requests/pkg/captures/captures_mock"
	"checkr.com/idempotent-requests/pkg/tracing"
	"encoding/base64"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

const CapturesV2URL = "/api/v2/captures"

func TestCaptures_Allocate(t *testing.T) {
	capturesRepo := captures_mock.NewRepositoryImpl()
	tracer := tracing.MockTracer{}

	router := NewRouter(tracer, capturesRepo)

	tests := []struct {
		name       string
		req        *http.Request
		expResCode int
		expResBody string
	}{
		{name: "Captures_Allocate", req: NewAllocationRequest(map[string]string{"idempotency_key": base64.RawURLEncoding.EncodeToString([]byte("qwe"))}), expResCode: http.StatusAccepted, expResBody: ""},
		{name: "Captures_Allocate_UUID_With_Extra_Char", req: NewAllocationRequest(map[string]string{"idempotency_key": base64.RawURLEncoding.EncodeToString([]byte("F3028015-BB17-479B-B2D6-3A096B7980592"))}), expResCode: http.StatusAccepted, expResBody: ""},
		{name: "Captures_Allocate_PreEncoded", req: NewAllocationRequest(map[string]string{"idempotency_key": "RjMwMjgwMTUtQkIxNy00NzlCLUIyRDYtM0EwOTZCNzk4MDU5MQ"}), expResCode: http.StatusAccepted, expResBody: ""},
		{name: "Captures_Allocate_Malformed_Idempotency_Key", req: NewAllocationRequest(map[string]string{"idempotency_key": "==qwe--"}), expResCode: http.StatusUnprocessableEntity, expResBody: ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			respRecorder := httptest.NewRecorder()
			router.ServeHTTP(respRecorder, tt.req)
			assert.Equal(t, tt.expResCode, respRecorder.Code)

			if tt.expResBody != "" {
				assert.Equal(t, tt.expResBody, respRecorder.Body.String())
			}
		})
	}

}

func NewAllocationRequest(payload interface{}) *http.Request {
	return newRequest(http.MethodPut, CapturesV2URL, reqBody(payload))
}

func newRequest(method, url string, body io.Reader) *http.Request {
	req, err := http.NewRequest(method, url, body)

	if err != nil {
		panic(err.Error())
	}

	return req
}

func reqBody(payload interface{}) *bytes.Buffer {
	putBody, err := json.Marshal(payload)

	if err != nil {
		panic(err.Error())
	}

	return bytes.NewBuffer(putBody)
}
