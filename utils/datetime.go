package goease

import "time"

// ParseRFC3339Date parses a date string in RFC3339 format.
//
// Parameters:
//   - dateStr: string - The date string to parse.
//
// Returns:
//   - time.Time: The parsed time if successful, otherwise a zero time.
func ParseRFC3339Date(dateStr string) time.Time {
	if dateStr == "" {
		return time.Time{} // Return a zero time for empty string
	}

	parsedTime, err := time.Parse(time.RFC3339, dateStr)
	if err != nil {
		// Handle the error here, depending on your application logic
		// For example, log the error or return a default value
		return time.Time{} // Return a zero time for invalid input
	}

	return parsedTime
}

// ParseCustomDate parses a date string in a custom format.
//
// Parameters:
//   - dateStr: string - The date string to parse.
//   - layout: string - The layout to use for parsing.
//
// Returns:
//   - time.Time: The parsed time if successful, otherwise a zero time.
func ParseCustomDate(dateStr, layout string) time.Time {
	if dateStr == "" {
		return time.Time{} // Return a zero time for empty string
	}

	parsedTime, err := time.Parse(layout, dateStr)
	if err != nil {
		// Handle the error here, depending on your application logic
		// For example, log the error or return a default value
		return time.Time{} // Return a zero time for invalid input
	}

	return parsedTime
}

// ParseISO8601Date parses a date string in ISO8601 format.
//
// Parameters:
//   - dateStr: string - The date string to parse.
//
// Returns:
//   - time.Time: The parsed time if successful, otherwise a zero time.
func ParseISO8601Date(dateStr string) time.Time {
	return ParseCustomDate(dateStr, "2006-01-02T15:04:05Z07:00")
}
