package utils

import (
	"bytes"
	"checkr.com/idempotent-requests/pkg/views"
	"encoding/json"
	"io"
	"net/http"
)

func NewAllocateCaptureRequest(url string, payload interface{}) *http.Request {
	return newRequest(http.MethodPut, url, reqBody(payload))
}

func NewRecordCaptureRequest(url string, payload interface{}) *http.Request {
	return newRequest(http.MethodPost, url, reqBody(payload))
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

func CaptureAllocatePayload(idempotencyKey string) interface{} {
	return views.CaptureAllocation{
		IdempotencyKey: idempotencyKey,
	}
}

func CaptureRecordPayload(idempotencyKey string) interface{} {
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