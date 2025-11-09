package kekasigohelper

import (
	b64 "encoding/base64"
	"encoding/json"
	"math/rand"
	"net/http"
	"os"
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

func OnlyOneRemove(str string, remove string) string {
	count := strings.Count(str, remove)
	if count == 1 {
		return strings.Replace(str, remove, "", 1)
	}
	return str
}

func RemoveDuplicates(input []int) []int {
	uniqueMap := make(map[int]bool)
	var result []int

	for _, value := range input {
		if !uniqueMap[value] {
			uniqueMap[value] = true
			result = append(result, value)
		}
	}

	return result
}

func ReverseSlice[T any](s []T) []T {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}

func ReadJSONFile[T any](filename string) (*T, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var data T
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&data); err != nil {
		return nil, err
	}
	return &data, nil
}

func InArrayGeneric[T comparable](val T, arr []T) bool {
    for _, item := range arr {
        if item == val {
            return true
        }
    }
    return false
}

func GetTimeAddMinutes(addMinute int) string {
    // Get current time
	currentTime := time.Now()
	
	// Add minutes
	newTime := currentTime.Add(time.Duration(addMinute) * time.Minute)
	
	// Format as HH:MM
	return newTime.Format("15:04")
}