package views

var MalformedPayload = NewError("Malformed payload")
var MalformedIdempotencyKey = NewError("Malformed idempotency key. Use URL safe base64 encoding")
var CaptureIsCompleted = NewError("Capture has been completed or has not been allocated")

func NewError(msg string) ErrorPayload {
	return ErrorPayload{Error: msg}
}

type ErrorPayload struct {
	Error string `json:"error"`
}
