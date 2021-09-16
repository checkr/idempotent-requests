package adapters

import (
	"checkr.com/idempotent-requests/pkg/captures"
	"checkr.com/idempotent-requests/pkg/views"
)

func AllocationToCaptureRecord(allocation *captures.Allocation) *views.CaptureRecord {
	respHeaders := make([]views.ResponseHeader, len(allocation.Capture.ResponseHeaders))

	for i, header := range allocation.Capture.ResponseHeaders {
		respHeaders[i] = views.ResponseHeader{
			Key:   header.Key,
			Value: header.Value,
		}
	}

	capture := &views.Capture{
		ResponseStatus:  allocation.Capture.ResponseStatus,
		ResponseBody:    allocation.Capture.ResponseBody,
		ResponseHeaders: respHeaders,
	}

	record := &views.CaptureRecord{
		IdempotencyKey: allocation.IdempotencyKey,
		Response:       capture,
	}

	return record
}

func CaptureRecordToAllocation(capture *views.CaptureRecord) *captures.Allocation {
	respHeaders := make([]captures.ResponseHeader, len(capture.Response.ResponseHeaders))

	for i, header := range capture.Response.ResponseHeaders {
		respHeaders[i] = captures.ResponseHeader{
			Key:   header.Key,
			Value: header.Value,
		}
	}

	response := &captures.Capture{
		ResponseStatus:  capture.Response.ResponseStatus,
		ResponseBody:    capture.Response.ResponseBody,
		ResponseHeaders: respHeaders,
	}

	allocation := &captures.Allocation{
		IdempotencyKey: capture.IdempotencyKey,
		Capture:        response,
		Status:         captures.StatusCompleted,
	}

	return allocation
}
