package http

import (
	"checkr.com/idempotent-requests/pkg/adapters"
	"checkr.com/idempotent-requests/pkg/captures"
	"checkr.com/idempotent-requests/pkg/views"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
)

type AllocateCaptureHandler struct {
	capturesRepo captures.Repository
}

func NewAllocateCaptureHandler(capturesRepo captures.Repository) Handler {
	return AllocateCaptureHandler{
		capturesRepo: capturesRepo,
	}
}

func (h AllocateCaptureHandler) Handle(c *gin.Context) {

	var attemptedAllocation views.CaptureAllocation
	if err := c.BindJSON(&attemptedAllocation); err != nil {
		c.JSON(http.StatusBadRequest, views.MalformedPayload)
		return
	}

	if !validIdempotencyKey(attemptedAllocation.IdempotencyKey) {
		c.JSON(http.StatusUnprocessableEntity, views.MalformedIdempotencyKey)
		return
	}

	allocation, err := h.capturesRepo.Allocate(c.Request.Context(), attemptedAllocation.IdempotencyKey)

	if err != nil {
		zap.S().Errorw("failed to allocate a capture", "err", err.Error())
		c.Status(http.StatusInternalServerError)
		return
	}

	switch allocation.Status {
	case captures.StatusCompleted:
		c.JSON(http.StatusOK, adapters.AllocationToCaptureRecord(allocation))
		return
	case captures.StatusInFlightCapture:
		c.Status(http.StatusConflict)
		return
	case captures.StatusAllocated:
		c.Status(http.StatusAccepted)
	default:
		zap.S().Errorw("unknown allocation status", "status", allocation.Status)
		c.Status(http.StatusInternalServerError)
		return
	}

}
