package kekasigohelper

import (
	"testing"
)

func TestUcFirst(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"kekasi", "Kekasi"},
		{"Kekasi", "Kekasi"},
		{"", ""},
		{"123abc", "123abc"},
		{"a", "A"},
	}

	for _, test := range tests {
		if output := UcFirst(test.input); output != test.expected {
			t.Errorf("UcFirst(%q) = %q; want %q", test.input, output, test.expected)
		} else {
			LoggerInfo(output)
		}
	}
}
