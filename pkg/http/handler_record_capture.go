package http

import (
	"checkr.com/idempotent-requests/pkg/adapters"
	"checkr.com/idempotent-requests/pkg/captures"
	"checkr.com/idempotent-requests/pkg/util"
	"checkr.com/idempotent-requests/pkg/views"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
)

type RecordCaptureHandler struct {
	capturesRepo captures.Repository
}

func NewRecordCaptureHandler(capturesRepo captures.Repository) Handler {
	return RecordCaptureHandler{
		capturesRepo: capturesRepo,
	}
}

func (h RecordCaptureHandler) Handle(c *gin.Context) {

	var capture views.CaptureRecord
	if err := c.BindJSON(&capture); err != nil {
		c.JSON(http.StatusBadRequest, views.MalformedPayload)
		return
	}

	if !validIdempotencyKey(capture.IdempotencyKey) {
		c.JSON(http.StatusUnprocessableEntity, views.MalformedIdempotencyKey)
		return
	}

	allocation := adapters.CaptureRecordToAllocation(&capture)

	if err := h.capturesRepo.Record(c, allocation); err != nil {

		if err == captures.ErrConflictOrMissing {
			c.JSON(http.StatusForbidden, views.CaptureIsCompleted)
			return
		}

		zap.S().Errorw("failed to record a capture", util.KeyErr, err.Error())
		c.Status(http.StatusInternalServerError)
		return
	}

	c.Status(http.StatusOK)

}
