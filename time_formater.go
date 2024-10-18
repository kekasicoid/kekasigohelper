package kekasigohelper

import (
	"fmt"
	"time"
)

func DateDiffMonth(startDate string, endDate string) (month float64, err error) {
	// Parse the dates
	layout := "2006-01-02"
	start, err := time.Parse(layout, startDate)
	if err != nil {
		return 0.0, fmt.Errorf("error parsing start date: " + err.Error())
	}
	end, err := time.Parse(layout, endDate)
	if err != nil {
		return 0.0, fmt.Errorf("error parsing end date: " + err.Error())
	}
	return float64(end.Sub(start).Hours() / 24 / 30), nil
}