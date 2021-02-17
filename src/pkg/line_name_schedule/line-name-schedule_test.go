package line_name_schedule

import (
	"strings"
	"testing"
)

type LineNameScheduleResult struct {
	lineStr    string
	expected Line
}
var lineNameScheduleResults = []LineNameScheduleResult{
	{ "RENE=MO10:00-12:00", Line{ "RENE", []string{"MO10:00-12:00"} } },
	{ "ASTRID=MO10:00-12:00,TH12:00-14:00,SU20:00-21:00",
		Line{ "ASTRID", []string{"MO10:00-12:00", "TH12:00-14:00", "SU20:00-21:00"} } },
	{ "DAVID=MO10:00-12:00,TU10:00-12:00,TH01:00-03:00,SU20:00-21:00",
		Line{ "DAVID", []string{"MO10:00-12:00", "TU10:00-12:00", "TH01:00-03:00", "SU20:00-21:00"} } },
	{ "MONICA=TH12:00-14:00,SU20:00-21:00",
		Line{ "MONICA", []string{"TH12:00-14:00", "SU20:00-21:00"} }},
	{ "JAQUE=MO10:00-12:00,TU10:00-12:00,SA14:00-18:00,SU20:00-21:00",
		Line{ "JAQUE", []string{"MO10:00-12:00", "TU10:00-12:00", "SA14:00-18:00", "SU20:00-21:00"} }},
	{ "ANNA=MO10:00-12:00,TH12:00-14:00",
		Line{ "ANNA", []string{"MO10:00-12:00", "TH12:00-14:00"} }},
}

func TestSplitNameAndSchedule(t *testing.T) {
	for index, test := range lineNameScheduleResults {
		result := SplitNameAndSchedule(&test.lineStr)
		if result.Name != test.expected.Name ||
			strings.Join(result.Schedule, "") != strings.Join(test.expected.Schedule, "") {
			t.Errorf("Expected result <<%q>> not given \n\t(%d. %q)", test.expected, index, test.lineStr)
		}
	}
}