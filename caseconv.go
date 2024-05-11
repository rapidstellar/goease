package goease

import "unicode"

// ConvertPascalToSnakeWithExtraKey converts keys in a map from PascalCase to snake_case.
// It also checks for additional key mappings defined in configs.KEY_CONVERT_MAPPING
// and uses those mappings if available.
//
// Parameters:
//   input: A map[string]interface{} representing the input data with keys possibly in PascalCase.
//
// Returns:
//   A map[string]interface{} with keys converted to snake_case. If a key is found in
//   configs.KEY_CONVERT_MAPPING, it will be replaced with the corresponding value. If not,
//   the key will be converted to snake_case.
func ConvertPascalToSnakeWithExtraKey(input map[string]interface{}, extraKeyMappings map[string]string) map[string]interface{} {
	convertedItem := make(map[string]interface{})

	for key, value := range input {
		// Check if the key is in the extra key mappings
		if mappedKey, ok := extraKeyMappings[key]; ok {
			convertedItem[mappedKey] = value
		} else {
			// If not in mappings, convert to snake_case
			snakeKey := convertPascalToSnakeCase(key)
			convertedItem[snakeKey] = value
		}
	}

	return convertedItem
}

// convertPascalToSnakeCase converts a string from PascalCase to snake_case.
//
// Parameters:
//   s: A string in PascalCase.
//
// Returns:
//   A string converted to snake_case.
func convertPascalToSnakeCase(s string) string {
	var result []rune
	for i, r := range s {
		if i > 0 && unicode.IsUpper(r) {
			result = append(result, '_')
		}
		result = append(result, unicode.ToLower(r))
	}
	return string(result)
}
