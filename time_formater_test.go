package kekasigohelper

import (
	"testing"
)

func TestDateDiffMonth(t *testing.T) {
	startDate := "2023-01-01"
	endDate := "2023-05-05"

	data, err := DateDiffMonth(startDate, endDate)
	if err != nil {
		t.Error(err)
	}
	LoggerInfo(data)
}
