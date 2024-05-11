package goease

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"log"
	"reflect"
)

// JSONB represents a JSONB type typically used to store JSON data in databases.
//
// JSONB is a custom type defined as a map[string]interface{}, allowing for flexible representation of JSON data.
//
// Methods:
//   - Value(): (JSONB method) Converts the JSONB value into a driver.Value for database storage.
//   - Scan(value interface{}): (JSONB method) Populates the JSONB value from a database driver.Value.
//
// Usage Example:
//
//	type MyData struct {
//	    Name string
//	    Age  int
//	}
//
//	jsonData := JSONB{"data": MyData{Name: "John", Age: 30}}
//
//	// Conversion to driver.Value for database storage
//	dbValue, err := jsonData.Value()
//
//	// Populating JSONB from database driver.Value
//	var retrievedData JSONB
//	err := retrievedData.Scan(dbValue)
//
// Note:
//   - JSONB is typically used to represent JSON data in databases that support JSONB storage.
//   - The Value() method is used to convert JSONB into a database-friendly format for storage.
//   - The Scan() method is used to populate JSONB from a database driver.Value.
type JSONB map[string]interface{}

// Value converts the JSONB value into a driver.Value for database storage.
//
// This method converts the JSONB value into a string representation before returning it as a driver.Value. It's typically used for storing JSONB data in databases that support JSONB storage.
//
// Returns:
//   - driver.Value: A driver.Value representing the JSONB value for database storage.
//   - error: An error if there's any issue during the conversion process.
//
// Note:
//   - This method internally uses the encoding/json package to marshal the JSONB value into a string.
//   - Any errors during the conversion process will be returned as an error.
func (j JSONB) Value() (driver.Value, error) {
	defer func() {
		if r := recover(); r != nil {
			log.Println("--------------JSONB Value-----------------")
			err := fmt.Errorf("panic occurred: %v", r)
			log.Println(err)
			log.Println("-------------------------------")
		}
	}()
	valueString, err := json.Marshal(j)
	return string(valueString), err
}

// Scan populates the JSONB value from a database driver.Value.
//
// This method takes a database driver.Value and populates the JSONB value with the corresponding data. It's typically used when retrieving JSONB data from databases that support JSONB storage.
//
// Parameters:
//   - value: interface{} - The database driver.Value to be scanned into the JSONB value.
//
// Returns:
//   - error: An error if there's any issue during the scanning process.
//
// Note:
//   - This method expects the database driver.Value to be a byte slice representing JSON data.
//   - Any errors during the scanning process will be returned as an error.
func (j *JSONB) Scan(value interface{}) error {
	defer func() {
		if r := recover(); r != nil {
			log.Println("--------------JSONB Scan-----------------")
			err := fmt.Errorf("panic occurred: %v", r)
			log.Println(err)
			log.Println("-------------------------------")
		}
	}()

	if data, ok := value.([]byte); ok {
		if err := json.Unmarshal(data, j); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("unexpected type for JSONB: %T", value)
	}

	return nil
}

// ConvertToJSONB converts two input data structures into JSONB types.
//
// This function takes two input interfaces representing data structures and converts them into JSONB types, which are custom types typically used to represent JSON data in databases that support JSONB storage.
//
// Parameters:
//   - oldData: interface{} - The old data structure to be converted into a JSONB type.
//   - newData: interface{} - The new data structure to be converted into a JSONB type.
//
// Returns:
//   - JSONB: A JSONB representation of the old data structure.
//   - JSONB: A JSONB representation of the new data structure.
//   - error: An error if there's any issue during the conversion process.
//
// Usage Example:
// Suppose you have two structs named 'OldData' and 'NewData' like this:
//
//	type OldData struct {
//	    Name string
//	    Age  int
//	}
//
//	type NewData struct {
//	    Name string
//	    Age  int
//	    City string
//	}
//
// You can convert instances of 'OldData' and 'NewData' into JSONB types like this:
// oldData := OldData{Name: "John", Age: 30}
// newData := NewData{Name: "Jane", Age: 25, City: "New York"}
// oldJSONB, newJSONB, err := ConvertToJSONB(oldData, newData)
//
//	if err != nil {
//	    fmt.Println("Error:", err)
//	    return
//	}
//
// Note:
//   - This function internally uses encoding/json package to marshal and unmarshal the data structures into JSONB types.
//   - Any errors during the conversion process will be returned as an error.
func ConvertToJSONB(oldData, newData interface{}) (JSONB, JSONB, error) {
	oldDataJSON, err := json.Marshal(oldData)
	if err != nil {
		return nil, nil, err
	}

	newDataJSON, err := json.Marshal(newData)
	if err != nil {
		return nil, nil, err
	}

	var oldDataMap map[string]interface{}
	if err := json.Unmarshal(oldDataJSON, &oldDataMap); err != nil {
		return nil, nil, err
	}

	var newDataMap map[string]interface{}
	if err := json.Unmarshal(newDataJSON, &newDataMap); err != nil {
		return nil, nil, err
	}

	return JSONB(oldDataMap), JSONB(newDataMap), nil
}

// StructToMap converts a struct into a map[string]interface{}.
//
// This function takes an input interface{} representing a struct and converts it into a map where the keys are the field names of the struct, and the values are the corresponding field values. It is particularly useful when you need to work with data in a more dynamic or generic way.
//
// Parameters:
//   - data: interface{} - The input value that should be a struct. It accepts an interface to be more flexible in what types can be passed.
//
// Returns:
//   - map[string]interface{}: A map representation of the struct where field names are the keys and field values are the values.
//   - error: An error if the input is not a struct.
//
// Usage Example:
// Suppose you have a struct named 'Person' like this:
//
//	type Person struct {
//	    Name  string
//	    Age   int
//	    Email string `json:"email"`
//	}
//
// You can convert an instance of 'Person' into a map using StructToMap like this:
// person := Person{Name: "John", Age: 30, Email: "john@example.com"}
// personMap, err := StructToMap(person)
//
//	if err != nil {
//	    fmt.Println("Error:", err)
//	    return
//	}
//
// The 'personMap' will contain:
//
//	map[string]interface{}{
//	    "Name":  "John",
//	    "Age":   30,
//	    "email": "john@example.com",
//	}
//
// Note:
//   - StructToMap supports struct tags to customize the keys in the resulting map. If a JSON tag is available for a field, it will be used as the key. Otherwise, the field name will be used.
//   - Only exported (public) fields of the struct can be converted, as unexported (private) fields cannot be accessed.
//   - It's important to ensure that the input 'data' is indeed a struct, as non-struct types will result in an error.
func StructToMap(data interface{}) (map[string]interface{}, error) {
	result := make(map[string]interface{})

	value := reflect.ValueOf(data)
	if value.Kind() == reflect.Ptr {
		value = value.Elem()
	}

	if value.Kind() != reflect.Struct {
		return nil, fmt.Errorf("not a struct")
	}

	typ := value.Type()
	for i := 0; i < value.NumField(); i++ {
		field := typ.Field(i)
		fieldValue := value.Field(i).Interface()

		// Use JSON tag if available, otherwise use field name
		jsonTag := field.Tag.Get("json")
		if jsonTag == "" {
			jsonTag = field.Name
		}
		result[jsonTag] = fieldValue
	}

	return result, nil
}
