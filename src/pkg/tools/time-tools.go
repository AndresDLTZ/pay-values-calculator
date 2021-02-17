package tools

import "time"

// diffHourTime calculates the difference between two hours, return difference in hours
func diffHourTime(startTime, endTime *time.Time) float64{
	return endTime.Sub(*startTime).Hours()
}
