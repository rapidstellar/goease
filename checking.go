package goease

import (
	"strconv"
	"strings"
	"time"
)

// Check if Int is in Slice
// Example usage:
// intSlice := []int{10, 20, 30, 40}
// intContains := IntContains(intSlice, 20)
// fmt.Println("Slice contains 20:", intContains)
func IntContains(slice []int, element int) bool {
	for _, item := range slice {
		if item == element {
			return true
		}
	}
	return false
}

// Trim String Spaces
// Example usage:
// trimmed := TrimSpaces("  hello world  ")
// fmt.Println("Trimmed string:", trimmed)
func TrimSpaces(str string) string {
	return strings.TrimSpace(str)
}

// Convert String to Boolean
// Example usage:
// b, err := StringToBool("true")
//
//	if err != nil {
//	    fmt.Println("Error converting string to bool:", err)
//	} else {
//
//	    fmt.Println("Converted boolean:", b)
//	}
func StringToBool(str string) (bool, error) {
	return strconv.ParseBool(str)
}

// Format Unix Time to String
// Example usage:
// formattedTime := FormatUnixTime(1609459200, "2006-01-02 15:04:05")
// fmt.Println("Formatted time:", formattedTime)
func FormatUnixTime(unixTime int64, layout string) string {
	return time.Unix(unixTime, 0).Format(layout)
}
