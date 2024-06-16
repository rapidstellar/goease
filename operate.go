package goease

import (
	"encoding/base64"
	"fmt"
	"strings"
)

// SplitString splits a string into an array of substrings based on a delimiter.
func SplitString(input, delimiter string) []string {
	return strings.Split(input, delimiter)
}

func ToLowerCase(text string) string {
	return strings.ToLower(text)
}

// DecodeBase64 decodes a base64 string into binary data.
//
// This function takes a base64 encoded string as input and decodes it into its binary representation. It returns the decoded binary data and any error encountered during the decoding process.
//
// Parameters:
//   - base64Str: string - The base64 encoded string to decode.
//
// Returns:
//   - []byte: The decoded binary data.
//   - error: An error if the decoding process fails.
//
// Example:
// You can use DecodeBase64 to decode a base64 string like this:
//
//	base64Str := "SGVsbG8gV29ybGQ="
//	data, err := DecodeBase64(base64Str)
//	if err != nil {
//	    fmt.Println("Error:", err)
//	    return
//	}
//
// This will decode the base64 string into binary data.
func DecodeBase64(base64Str string) ([]byte, error) {
	// Decode the base64 string into binary data
	data, err := base64.StdEncoding.DecodeString(base64Str)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// ExtractImageTypeFromBase64 extracts the image type from a base64 encoded data URI.
//
// This function takes a data URI string as input, which should be in the format "data:image/type;base64,...", and extracts the image type from it. It returns the extracted image type and any error encountered during the extraction process.
//
// Parameters:
//   - dataURI: string - The data URI string from which to extract the image type.
//
// Returns:
//   - string: The extracted image type (e.g., "jpeg", "png").
//   - error: An error if the data URI format is invalid or if the extraction process fails.
//
// Example:
// Suppose you have a data URI string representing a JPEG image like this:
//
//	dataURI := "data:image/jpeg;base64,/9j/4AAQSkZJRgABAQEAYABgAAD/4QA6RXhpZgAATU0AKgAAAAgAAQA"
//	imageType, err := ExtractImageTypeFromBase64(dataURI)
//	if err != nil {
//	    fmt.Println("Error:", err)
//	    return
//	}
//
// This will extract the image type "jpeg" from the data URI.
func ExtractImageTypeFromBase64(dataURI string) (string, error) {
	// Check if the data URI starts with "data:image/". If not, return an error.
	if !strings.HasPrefix(dataURI, "data:image/") {
		return "", fmt.Errorf("invalid data URI format")
	}

	// Find the end of the image type declaration (e.g., "data:image/jpeg;base64,")
	endIndex := strings.Index(dataURI, ";base64,")
	if endIndex == -1 {
		return "", fmt.Errorf("invalid data URI format")
	}

	// Extract and return the image type.
	imageType := dataURI[len("data:image/"):endIndex]
	return imageType, nil
}
