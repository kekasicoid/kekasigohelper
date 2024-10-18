package response

type Response struct {
	Code   string      `json:"response_code"`
	Refnum string      `json:"response_refnum"`
	ID     string      `json:"response_id"`
	Desc   string      `json:"response_desc"`
	Data   interface{} `json:"response_data"`
}

// Initialization of Response
func New(id string) *Response {
	return &Response{
		ID:   id,
		Code: "XX",
		Desc: "General Error",
		Data: new(struct{}),
	}
}
