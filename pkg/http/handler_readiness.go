package http

import (
	"checkr.com/idempotent-requests/pkg/captures"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ReadinessHandler struct {
	capturesRepo captures.Repository
}

func NewReadinessHandler(capturesRepo captures.Repository) Handler {
	return ReadinessHandler{
		capturesRepo: capturesRepo,
	}
}

func (h ReadinessHandler) Handle(c *gin.Context) {

	ready := h.capturesRepo.Ready(c)

	if ready {
		c.Status(http.StatusOK)
	} else {
		c.Status(http.StatusServiceUnavailable)
	}

	return
}
