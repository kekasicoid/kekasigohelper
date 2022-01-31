package kekasigohelper

import "net/http"

type Response struct {
	StatusCode int         `json:"status_code"`
	Status     string      `json:"status_desc"`
	Msg        string      `json:"message"`
	Data       interface{} `json:"data"`
	Errors     interface{} `json:"errors"`
}

func (r *Response) MappingResponseSuccess(message string, data interface{}) {
	r.StatusCode = http.StatusOK
	r.Status = http.StatusText(r.StatusCode)
	r.Msg = message
	r.Data = data
	r.Errors = nil
}

func (r *Response) MappingResponseError(statusCode int, message string, error interface{}) {
	r.StatusCode = statusCode
	r.Status = http.StatusText(r.StatusCode)
	r.Msg = message
	r.Data = nil
	r.Errors = error
}
