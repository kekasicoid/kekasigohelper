package kekasigohelper

import (
	"fmt"
	"strconv"
)

func AnyToString(value interface{}) string {
	switch v := value.(type) {
	case string:
		return v
	case int:
		return strconv.Itoa(v)
	case float64:
		return strconv.FormatFloat(v, 'f', -1, 64)
	// Add cases for other types as needed
	default:
		return fmt.Sprintf("%v", v) // Fallback to fmt for other types
	}
}

// FloatToString converts a float64 to a string with the specified precision
func FloatToString(value float64, precision int) string {
	return fmt.Sprintf("%.*f", precision, value)
}
