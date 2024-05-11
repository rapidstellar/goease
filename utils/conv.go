package goease

import (
	"reflect"
	"strconv"
	"strings"
)

// IsSlice checks if the given value is a slice.
//
// This function takes an input value and checks whether it is a slice or not. It utilizes reflection to determine the kind of the value.
//
// Parameters:
//   - v: interface{} - The value to be checked.
//
// Returns:
//   - bool: true if the value is a slice, false otherwise.
//
// Usage Example:
//
//	var arr []int
//	result := IsSlice(arr) // result will be true
//
//	var str string
//	result := IsSlice(str) // result will be false
//
// Note:
//   - This function is useful for checking whether a value is a slice before performing operations specific to slices.
//   - It uses reflection to determine the kind of the value, which may have a performance overhead.
func IsSlice(v interface{}) bool {
	rv := reflect.ValueOf(v)
	return rv.Kind() == reflect.Slice
}

// String to Int Conversion
// Example usage:
// num, err := StringToInt("123")
//
//	if err != nil {
//	    fmt.Println("Error converting string to int:", err)
//	} else {
//
//	    fmt.Println("Converted number:", num)
//	}
func StringToInt(str string) (int, error) {
	return strconv.Atoi(str)
}

// String to Float Conversion
// Example usage:
// f, err := StringToFloat("123.45")
//
//	if err != nil {
//	    fmt.Println("Error converting string to float:", err)
//	} else {
//
//	    fmt.Println("Converted float:", f)
//	}
func StringToFloat(str string) (float64, error) {
	return strconv.ParseFloat(str, 64)
}

// Int to String Conversion
// Example usage:
// str := IntToString(123)
// fmt.Println("Converted string:", str)
func IntToString(num int) string {
	return strconv.Itoa(num)
}

// Helper function to check if a string is in a slice of strings
// Check if String is in Slice
// Example usage:
// slice := []string{"apple", "banana", "cherry"}
// contains := StringContains(slice, "banana")
// fmt.Println("Slice contains 'banana':", contains)
func StringContains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

// Join Int Slice to String
// Example usage:
// ints := []int{1, 2, 3, 4}
// joined := JoinInts(ints, ", ")
// fmt.Println("Joined string:", joined)
func JoinInts(ints []int, sep string) string {
	var strSlice []string
	for _, num := range ints {
		strSlice = append(strSlice, strconv.Itoa(num))
	}
	return strings.Join(strSlice, sep)
}

// Float to String Conversion
// Example usage:
// floatStr := FloatToString(123.456)
// fmt.Println("Converted string:", floatStr)
func FloatToString(num float64) string {
	return strconv.FormatFloat(num, 'f', -1, 64)
}