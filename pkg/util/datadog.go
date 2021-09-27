package util

import "os"

func DDServiceName() string {
	s := os.Getenv("DD_SERVICE")
	if len(s) == 0 {
		s = "idempotent-requests"
	}
	return s
}
