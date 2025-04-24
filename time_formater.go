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

func DateDiffMidnight(now time.Time) time.Duration {
	midnight := time.Date(now.Year(), now.Month(), now.Day(), 23, 59, 59, 0, now.Location())
	return time.Until(midnight)
}

func DateDiffDay(startDate string, endDate string) (days int, err error) {
	// Parse the dates
	layout := "2006-01-02"
	start, err := time.Parse(layout, startDate)
	if err != nil {
		return 0, fmt.Errorf("error parsing start date: " + err.Error())
	}
	end, err := time.Parse(layout, endDate)
	if err != nil {
		return 0, fmt.Errorf("error parsing end date: " + err.Error())
	}
	// Calculate the difference in days
	return int(end.Sub(start).Hours() / 24), nil
}
