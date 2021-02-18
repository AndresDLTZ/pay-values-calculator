package validator

import "regexp"

// getRegexpCompile will retrieve a regular expression pattern and return Regexp type
func getRegexpCompile(pattern string) *regexp.Regexp {
	r, err := regexp.Compile(pattern)
	if err != nil {
		panic(err)
	}
	return r
}
