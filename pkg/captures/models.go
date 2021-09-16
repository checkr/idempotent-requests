package captures

const (
	StatusCompleted       = "completed"
	StatusInFlightCapture = "capture_is_in_flight"
	StatusAllocated       = "allocated"
)

type Allocation struct {
	IdempotencyKey string   `bson:"idempotency_key,omitempty"`
	Capture        *Capture `bson:"capture,omitempty"`
	Status         string   `bson:"status,omitempty"`
}

type Capture struct {
	ResponseStatus  int32            `bson:"response_status,omitempty"`
	ResponseBody    string           `bson:"response_body,omitempty"`
	ResponseHeaders []ResponseHeader `bson:"response_headers,omitempty"`
}

type ResponseHeader struct {
	Key   string `bson:"key,omitempty"`
	Value string `bson:"value,omitempty"`
}

var AllocationIndex = map[string]int8{
	"idempotency_key": 1,
}
