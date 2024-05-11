package goease

import "strings"

// SplitString splits a string into an array of substrings based on a delimiter.
func SplitString(input, delimiter string) []string {
	return strings.Split(input, delimiter)
}

func ToLowerCase(text string) string {
	return strings.ToLower(text)
}
