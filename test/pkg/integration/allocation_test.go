package integration

import (
	"checkr.com/idempotent-requests/pkg/views"
	"checkr.com/idempotent-requests/tests/integation/pkg/utils"
	"encoding/base64"
	"encoding/json"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"reflect"
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

	if assert.NoError(t, err) {
		assert.Equal(t, http.StatusAccepted, resp.StatusCode)
	}

	captureRecordReq := utils.NewRecordCaptureRequest(CapturesV2URL, utils.CaptureRecordPayload(idempotencyKey))

	resp, err = httpClient.Do(captureRecordReq)

	if assert.NoError(t, err) {
		assert.Equal(t, http.StatusOK, resp.StatusCode)
	}

	resp, err = httpClient.Do(allocationReq)
	if assert.NoError(t, err) {
		assert.Equal(t, http.StatusOK, resp.StatusCode)
	}

	defer resp.Body.Close()
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	assert.NoError(t, err)

	var capture views.CaptureRecord
	err = json.Unmarshal(bodyBytes, &capture)

	if assert.NoError(t, err) {
		assert.True(t, reflect.DeepEqual(capture.Response, utils.CaptureRecordPayload(idempotencyKey).(views.CaptureRecord).Response))
	}

}

func TestCaptures_Allocate_Duplicate(t *testing.T) {

	cfg := GetConfig()

	CapturesV2URL := cfg.Host + CapturesV2Path

	utils.WaitForServerToBeReady(cfg.Host)

	httpClient := &http.Client{
		Timeout: 10 * time.Second,
	}

	idempotencyKey := base64.URLEncoding.EncodeToString([]byte(uuid.New().String()))

	allocationReq := utils.NewAllocateCaptureRequest(CapturesV2URL, utils.CaptureAllocatePayload(idempotencyKey))

	resp, err := httpClient.Do(allocationReq)

	if assert.NoError(t, err) {
		assert.Equal(t, http.StatusAccepted, resp.StatusCode)
	}

	resp, err = httpClient.Do(allocationReq)

	if assert.NoError(t, err) {
		assert.Equal(t, http.StatusConflict, resp.StatusCode)
	}

}
