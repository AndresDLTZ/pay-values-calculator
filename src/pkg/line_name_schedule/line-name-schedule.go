package line_name_schedule

import "strings"

// Line type provides a structure that takes care of getting Name and Schedule that constitute a line
type Line struct {
	Name 	 string
	Schedule []string
}

// SplitNameAndSchedule will retrieve a line and return a Line type
func SplitNameAndSchedule(line *string) Line {
	splitLine := strings.Split(*line, "=")
	name, schedule := splitLine[0], strings.Split(splitLine[1], ",")
	return Line{Name: name, Schedule: schedule}
}