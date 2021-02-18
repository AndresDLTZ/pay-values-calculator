package tools

import (
	"time"
)

// ParseStringToTime will retrieve a string and return a time
func ParseStringToTime(timeStr *string) time.Time {
	// layout is format of string
	layout := "15:04"
	varTime, err := time.Parse(layout, *timeStr)
	if err != nil {
		panic(err)
	}
	if *timeStr == "00:00" {
		varTime = varTime.Add(24 * time.Hour)
	}
	return varTime
}

// DiffHourTime calculates the difference between two hours, return difference in hours
func DiffHourTime(startTime, endTime *time.Time) float64{
	return endTime.Sub(*startTime).Hours()
}
