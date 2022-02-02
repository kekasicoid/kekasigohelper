package kekasigohelper

import (
	"net/http"
	"strings"
)

func InArray(str string, list []string) bool {
	str = strings.ToLower(str)
	for _, v := range list {
		if strings.ToLower(v) == str {
			return true
		}
	}
	return false
}
func GetStatusCode(err error) int {
	if err == nil {
		return http.StatusOK
	}
	switch err {
	case ErrInternalServerError:
		return http.StatusInternalServerError
	case ErrForbidden:
		return http.StatusForbidden
	case ErrNotFound:
		return http.StatusNotFound
	case ErrUnAuthorize:
		return http.StatusUnauthorized
	case ErrBadParamInput:
		return http.StatusBadRequest
	default:
		return http.StatusInternalServerError
	}
}
