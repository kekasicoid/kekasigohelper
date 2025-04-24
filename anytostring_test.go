package kekasigohelper

import (
	"testing"
)

func TestFloatToString(t *testing.T) {
	tests := []struct {
		name      string
		value     float64
		precision int
		expected  string
	}{
		{"Zero precision", 123.456, 0, "123"},
		{"Two decimal places", 123.456, 2, "123.46"},
		{"No rounding needed", 123.4, 2, "123.40"},
		{"High precision", 123.456789, 6, "123.456789"},
		{"Negative value", -123.456, 2, "-123.46"},
		{"Zero value", 0.0, 2, "0.00"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := FloatToString(tt.value, tt.precision)
			if result != tt.expected {
				t.Errorf("FloatToString(%f, %d) = %s; want %s", tt.value, tt.precision, result, tt.expected)
			}
		})
	}
}
