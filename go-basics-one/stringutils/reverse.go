package stringutils

// Reverse function reverse string and return string
func Reverse(str string) (result string) {
	for _, v := range str {
		result = string(v) + result
	}
	return result
}
