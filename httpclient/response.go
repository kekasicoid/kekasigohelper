// Package httpclient
// @author Daud Valentino
package httpclient

import (
	"encoding/json"
	"errors"
	"net/http"
	"reflect"
	"time"
)

var (
	errorCannotAddressable = errors.New("destination cannot addressable")
)

// ClientResponse representation of http client response
type Response struct {
	StatusCode int
	Body       []byte
	BodyStr    string
	Latency    time.Duration
	Header     http.Header
}

// DecodeJSON decode response byte to struct
func (cr Response) DecodeJSON(dest interface{}) error {

	v := reflect.ValueOf(dest)
	if v.Kind() != reflect.Ptr {
		return errorCannotAddressable
	}

	return json.Unmarshal(cr.Body, dest)
}

// String cast response byte to string
func (cr Response) String() string {
	return string(cr.Body)
}

// RawByte return raw byte data
func (cr Response) RawByte() []byte {
	return cr.Body
}

// Header return http header response
func (cr Response) GetHeader() http.Header {
	return cr.Header
}

// Header return http status response
func (cr Response) Status() int {
	return cr.StatusCode
}

// Header return http status response
func (cr Response) GetLatency() time.Duration {
	return cr.Latency
}
