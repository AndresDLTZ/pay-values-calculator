package tools

import (
	"strconv"
	"testing"
	"time"
)

// array-tools testing
type ElementInArrayResult struct {
	array    []string
	key 	 string
	expected bool
}
var elementInArrayResults = []ElementInArrayResult{
	{ []string{"MO", "TU", "WE"}, "MO", true },
	{ []string{"MO", "TU", "WE"}, "TU", true },
	{ []string{"MO", "TU", "WE"}, "WE", true },
	{ []string{"MO", "TU", "WE"}, "FR", false },
	{ []string{"MO", "TU", "WE"}, "MM", false },
	{ []string{"MO", "TU", "WE"}, "", false },
}
func TestValidateIsElementInArray(t *testing.T) {
	for index, test := range elementInArrayResults {
		result := IsElementInArray(test.array, test.key)
		if result != test.expected {
			t.Errorf("Expected result <<%t>> not given \n\t(%d: %q, %q)", test.expected, index, test.array, test.key)
		}
	}
}

// time-tools testing
func TestParseStringToTime(t *testing.T) {
	hourStr := "12:45"
	result := ParseStringToTime(&hourStr)
	resultStr := strconv.Itoa(result.Hour()) + ":" + strconv.Itoa(result.Minute())
	if resultStr != hourStr {
		t.Errorf("Expected result <<%s>> time not given \n\t(%v)", hourStr, resultStr)
	}
}

type DiffHourTimeResult struct {
	startTime time.Time
	endTime   time.Time
	expected  float64
}
var now = time.Now()
var diffHourTimeResuls = []DiffHourTimeResult{
	{now, now, 0 },
	{now, now.Add(1 * time.Hour), 1 },
	{now, now.Add(2 * time.Hour), 2 },
	{now, now.Add(3 * time.Hour), 3 },
	{now, now.Add(4 * time.Hour), 4 },
}
func TestDiffHourTime(t *testing.T) {
	for index, test := range diffHourTimeResuls {
		result := DiffHourTime(&test.startTime, &test.endTime)
		if result != test.expected {
			t.Errorf("Expected result <<%v>> not given \n\t(%d: %q, %q)", test.expected, index, test.startTime, test.endTime)
		}
	}
}