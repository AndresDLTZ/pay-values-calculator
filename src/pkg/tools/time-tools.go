package tools

import (
	"github.com/andresdltz/pay-values-calculator/src/pkg/payment_rate"
	"strings"
	"time"
)

// GetDurationOfValidHourRange will retrieve a hour range and payment rate,
// check what schedule it is and return to the hour difference
func GetDurationOfValidHourRange(hrRange string, sr *payment_rate.ScheduleRate) float64 {
	hrArr := strings.Split(hrRange, "-")
	hr1 := ParseStringToTime(&hrArr[0])
	hr2 := ParseStringToTime(&hrArr[1])

	if !(hr1.Before(ParseStringToTime(&sr.StartTime)) || hr2.After(ParseStringToTime(&sr.EndTime))) {
		return diffHourTime(&hr1, &hr2)
	}
	return -1
}

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

// diffHourTime calculates the difference between two hours, return difference in hours
func diffHourTime(startTime, endTime *time.Time) float64{
	return endTime.Sub(*startTime).Hours()
}
