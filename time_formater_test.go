package kekasigohelper

import (
	"testing"
)

// TestDateDiffDay tests the DateDiffDay function for various date inputs.
// It checks the difference in days between two dates and handles invalid date formats.
// The test cases include valid date pairs, same date, and invalid date formats.
func TestDateDiffDay(t *testing.T) {
	tests := []struct {
		startDate string
		endDate   string
		expected  int
		expectErr bool
	}{
		{"2023-01-01", "2023-01-02", 1, false},
		{"2023-01-01", "2023-01-01", 0, false},
		{"2023-01-01", "2023-01-10", 9, false},
		{"2023-01-10", "2023-01-01", -9, false},
		{"invalid-date", "2023-01-01", 0, true},
		{"2023-01-01", "invalid-date", 0, true},
	}

	for _, test := range tests {
		t.Run(test.startDate+"_"+test.endDate, func(t *testing.T) {
			days, err := DateDiffDay(test.startDate, test.endDate)
			if test.expectErr {
				if err == nil {
					t.Errorf("expected an error but got none")
				}
			} else {
				if err != nil {
					t.Errorf("did not expect an error but got: %v", err)
				}
				if days != test.expected {
					t.Errorf("expected %d days, got %d", test.expected, days)
				}
			}
		})
	}
}

// TestDateDiffMonth tests the DateDiffMonth function for various date inputs.
// It checks the difference in months between two dates and handles invalid date formats.
// The test cases include valid date pairs, same date, and invalid date formats.
func TestDateDiffMonth(t *testing.T) {
	tests := []struct {
		startDate string
		endDate   string
		expected  float64
		expectErr bool
	}{
		{"2023-01-01", "2023-02-02", 1.0666666666666667, false},
		{"2022-01-01", "2023-01-01", 12.166666666666666, false},
		{"2023-01-01", "2023-01-01", 0.0, false},
		{"2023-01-01", "2023-04-01", 3.0, false},
		{"2023-04-01", "2023-01-01", -3.0, false},
		{"invalid-date", "2023-01-01", 0.0, true},
		{"2023-01-01", "invalid-date", 0.0, true},
	}

	for _, test := range tests {
		t.Run(test.startDate+"_"+test.endDate, func(t *testing.T) {
			months, err := DateDiffMonth(test.startDate, test.endDate)
			if test.expectErr {
				if err == nil {
					t.Errorf("expected an error but got none")
				}
			} else {
				if err != nil {
					t.Errorf("did not expect an error but got: %v", err)
				}
				if months != test.expected {
					t.Errorf("expected %f months, got %f", test.expected, months)
				}
			}
		})
	}
}
