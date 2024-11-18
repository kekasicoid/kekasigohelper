package kekasigohelper

import (
	b64 "encoding/base64"
	"encoding/json"
	"math/rand"
	"net/http"
	"strings"
	"time"
	"unsafe"
)

// START Generate Random String
// https://stackoverflow.com/questions/22892120/how-to-generate-a-random-string-of-a-fixed-length-in-go

var src = rand.NewSource(time.Now().UnixNano())

const letterBytes = "ABCDEFGHIJKLMNOPQRSTUVWXYZ123456789"
const (
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
)

func RandStringBytesMaskImprSrcUnsafe(n int) string {
	b := make([]byte, n)
	// A src.Int63() generates 63 random bits, enough for letterIdxMax characters!
	for i, cache, remain := n-1, src.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}
	return *(*string)(unsafe.Pointer(&b))
}

// END Generate Random String

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

func Encrypt64(stringToEncrypt string) (encryptedString string) {
	res := b64.StdEncoding.EncodeToString([]byte(stringToEncrypt))
	return res
}

func Decrypt64(encryptedString string, keyString string) (decryptedString string, err error) {
	res, err := b64.StdEncoding.DecodeString(encryptedString)
	if err != nil {
		return "", ErrGeneralMessage
	}
	return string(res), nil
}

func ObjectToString(kks interface{}) string {
	res, _ := json.Marshal(kks)
	return string(res)
}

func HasDuplicateIntArray(nums []int) bool {
	seen := make(map[int]bool)

	for _, num := range nums {
		if seen[num] {
			// Duplicate found
			return true
		}
		seen[num] = true
	}

	// No duplicates found
	return false
}
