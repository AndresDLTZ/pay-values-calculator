package validator

import "errors"

var errNameScheduleRate = errors.New("nombre y horarios trabajados no se corresponde con el patrón")

// ValidateNameSchedule will retrieve a line and validate pattern (name and schedules worked)
func ValidateNameSchedule(nsr *string) error {
	dayRangeHourMinutes := "(MO|TU|WE|TH|FR|SA|SU)(([0-1][0-9]|[0-2][0-3]):[0-5][0-9]-([0-1][0-9]|[0-2][0-3]):[0-5][0-9])"
	r := getRegexpCompile("([A-Z])+=(" + dayRangeHourMinutes + ")(," + dayRangeHourMinutes + ")*$")
	if r.MatchString(*nsr) != true { return errNameScheduleRate }
	return nil
}
