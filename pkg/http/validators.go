package http

import (
	"checkr.com/idempotent-requests/pkg/views"
	"encoding/base64"
	"github.com/gin-gonic/gin"
	"net/http"
)

func validIdempotencyKey(c *gin.Context, idempotencyKey string) bool {
	if _, err := base64.URLEncoding.DecodeString(idempotencyKey); err == nil {
		return true
	}
	c.JSON(http.StatusUnprocessableEntity, views.MalformedIdempotencyKey)
	return false
}
