package validator

import (
	"os"
	"testing"
)

type NameScheduleResult struct {
	nameSchedules string
	expected      error
}
var nameScheduleResults = []NameScheduleResult {
	{"RENE=MO10:00-12:00,TU10:00-12:00,TH01:00-03:00,SA14:00-18:00,SU20:00-21:00", nil },
	{"ASTRID=MO10:00-12:00,TH12:00-14:00,SU20:00-21:00", nil },
	{"DAVI=MO10:00-12:00,TH12:00-14:00,SU20:00-21:00,", errNameScheduleRate },
	{"MONICA=MO10:00-12:00,TH12:00-14:00,SU20:00-21:00,", errNameScheduleRate },
	{"JOSE=MO10:00-12:00,TH12:00-14:00,SU20:00-2100", errNameScheduleRate },
	{"NAOMI=MO10:00-12:00,TH12:00-14:00,SU20:00-21:0,", errNameScheduleRate },
	{"JAQUELINE=MO10:00-12:00,TH12:00-14:00,SU17:00-89:00", errNameScheduleRate },
	{"MIA=MO10:00-12:00,TH12:00-14:00,SU20:00-221:00", errNameScheduleRate },
	{"MARIO=SU10:00-12:00", nil },
}
// TestValidateNameSchedule
func TestValidateNameSchedule(t *testing.T) {
	for index, test := range nameScheduleResults {
		err := ValidateNameSchedule(&test.nameSchedules)
		if err != test.expected {
			t.Errorf("Expected result not given \n\t(%d: %q)", index, test.nameSchedules)
		}
	}
}

// TestValidateFileExtension
func TestValidateFileExtension(t *testing.T) {
	file, _ := os.Open("file-to-test.txt")
	fileExtensionValidator := NewFileExtensionValidator(file, "txt")
	err := fileExtensionValidator.ValidateFileExtension()
	if err != nil {
		t.Errorf("Expected result not given (files/file-to-test.txt not txt)")
	}
}