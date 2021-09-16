package views

type CaptureRecord struct {
	IdempotencyKey string   `json:"idempotency_key"`
	Response       *Capture `json:"response"`
}
