package tools

import "testing"

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