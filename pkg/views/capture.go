package views

type Capture struct {
	ResponseStatus  int32            `json:"response_status"`
	ResponseBody    string           `json:"response_body"`
	ResponseHeaders []ResponseHeader `json:"response_headers"`
}
