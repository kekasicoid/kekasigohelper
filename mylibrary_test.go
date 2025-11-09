package kekasigohelper

import (
	"testing"
	"time"
)

func TestGetTimeAddMinutes(t *testing.T) {
	cases := []struct {
		name      string
		addMinute int
	}{
		{"zero", 0},
		{"positive_five", 5},
		{"negative_five", -5},
		{"large_minutes", 1500}, // more than a day
	}

	for _, tc := range cases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			// avoid parallel tests since they rely on current time
			// capture time before and after call to tolerate minute-boundary races
			nowBefore := time.Now()
			got := GetTimeAddMinutes(tc.addMinute)
			nowAfter := time.Now()

			exp1 := nowBefore.Add(time.Duration(tc.addMinute) * time.Minute).Format("15:04")
			exp2 := nowAfter.Add(time.Duration(tc.addMinute) * time.Minute).Format("15:04")

			if got != exp1 && got != exp2 {
				t.Fatalf("GetTimeAddMinutes(%d) = %q; want %q or %q", tc.addMinute, got, exp1, exp2)
			}
		})
	}
}