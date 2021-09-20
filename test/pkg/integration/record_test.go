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

func TestCaptures_Record_Without_Allocation(t *testing.T) {

	cfg := GetConfig()

	CapturesV2URL := cfg.Host + CapturesV2Path

	utils.WaitForServerToBeReady(cfg.Host)

	httpClient := &http.Client{
		Timeout: 10 * time.Second,
	}

	idempotencyKey := base64.URLEncoding.EncodeToString([]byte(uuid.New().String()))

	captureRecordReq := utils.NewRecordCaptureRequest(CapturesV2URL, utils.CaptureRecordPayload(idempotencyKey))

	resp, err := httpClient.Do(captureRecordReq)

	if assert.NoError(t, err) {
		assert.Equal(t, http.StatusForbidden, resp.StatusCode)
	}

}
