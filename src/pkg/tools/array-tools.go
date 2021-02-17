package tools

// IsElementInArray search key in array
func IsElementInArray(arr []string, key string) bool {
	if len(arr) == 0 { return false }
	if arr[0] == key { return true }
	return IsElementInArray(arr[1:], key)
}