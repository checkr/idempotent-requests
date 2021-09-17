package integration

import (
	"checkr.com/idempotent-requests/tests/integation/pkg/utils"
	"encoding/base64"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
	"time"
)

const (
	CapturesV2Path = "/api/v2/captures"
)

func TestCaptures_Allocate(t *testing.T) {

	cfg := GetConfig()

	CapturesV2URL := cfg.Host + CapturesV2Path

	utils.WaitForServerToBeReady(cfg.Host)

	httpClient := &http.Client{
		Timeout: 10 * time.Second,
	}

	idempotencyKey := base64.URLEncoding.EncodeToString([]byte(uuid.New().String()))

	allocationReq := utils.NewAllocateCaptureRequest(CapturesV2URL, utils.CaptureAllocatePayload(idempotencyKey))

	resp, err := httpClient.Do(allocationReq)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusAccepted, resp.StatusCode)

	captureRecordReq := utils.NewRecordCaptureRequest(CapturesV2URL, utils.CaptureRecordPayload(idempotencyKey))

	resp, err = httpClient.Do(captureRecordReq)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)

}
