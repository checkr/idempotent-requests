package http

import (
	"encoding/base64"
)

func validIdempotencyKey(idempotencyKey string) bool {
	if _, err := base64.RawURLEncoding.DecodeString(idempotencyKey); err == nil {
		return true
	}
	return false
}
