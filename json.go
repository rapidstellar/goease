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

// ReadJSONB reads JSON data into the target interface.
//
// This function unmarshals the JSON data contained in the 'jsonData' byte slice into the provided 'target' interface{}. The 'target' must be a pointer to the type into which the JSON data will be unmarshaled. If the unmarshaling process encounters an error, it returns that error. Otherwise, it returns nil.
//
// Parameters:
//   - jsonData: []byte - The JSON data to be unmarshaled.
//   - target: interface{} - A pointer to the type into which the JSON data will be unmarshaled.
//
// Returns:
//   - error: An error if the unmarshaling process fails. Otherwise, returns nil.
//
// Example:
// Assuming you have JSON data in a byte slice named 'jsonData' and a struct named 'User', you can use ReadJSONB as follows:
//
//	var user User
//	err := ReadJSONB(jsonData, &user)
//	if err != nil {
//	    fmt.Println("Error:", err)
//	    return
//	}
//
// This will unmarshal the JSON data in 'jsonData' into the 'user' struct.
func ReadJSONB(jsonData []byte, target interface{}) error {
	err := json.Unmarshal(jsonData, target)
	if err != nil {
		return err
	}
	return nil
}

// NewJSONB creates a new JSONB instance from the provided data.
//
// This function marshals the input 'data' into JSON format and then unmarshals it into a map[string]interface{}. It returns the created JSONB instance and any error encountered during the process.
//
// Parameters:
//   - data: interface{} - The data to be converted into JSONB. It can be any data type.
//
// Returns:
//   - JSONB: The created JSONB instance, which is essentially a map[string]interface{}.
//   - error: An error if the marshaling or unmarshaling process fails.
//
// Example:
// Suppose you have some data in a struct named 'Person'. You can create a JSONB instance from this data as follows:
//
//	person := Person{Name: "John", Age: 30}
//	jsonb, err := NewJSONB(person)
//	if err != nil {
//	    fmt.Println("Error:", err)
//	    return
//	}
//
// This will convert the 'person' struct into a JSONB instance.
func NewJSONB(data interface{}) (JSONB, error) {
	dataJSON, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	var dataMap map[string]interface{}
	if err := json.Unmarshal(dataJSON, &dataMap); err != nil {
		return nil, err
	}

	return JSONB(dataMap), nil
}

// MarshalJSONB marshals a JSONB instance into JSON format.
//
// This function takes a JSONB instance as input and marshals it into JSON format. It returns the JSON representation of the input data and any error encountered during the marshaling process.
//
// Parameters:
//   - data: JSONB - The JSONB instance to marshal into JSON format.
//
// Returns:
//   - []byte: The JSON representation of the input JSONB instance.
//   - error: An error if the marshaling process fails.
//
// Example:
// Suppose you have a JSONB instance named 'jsonData' representing some data. You can marshal it into JSON format like this:
//
//	jsonData := JSONB{"name": "John", "age": 30}
//	jsonBytes, err := MarshalJSONB(jsonData)
//	if err != nil {
//	    fmt.Println("Error:", err)
//	    return
//	}
//
// This will marshal the 'jsonData' into a byte slice containing the JSON representation.
func MarshalJSONB(data JSONB) ([]byte, error) {
	return json.Marshal(data)
}

// MarshalJSONBA marshals the provided JSONBA slice into JSON format.
//
// This function takes a JSONBA slice, which is essentially a slice of map[string]interface{}, and marshals it into JSON format using the encoding/json package.
//
// Parameters:
//   - data: JSONBA - The JSONBA slice to be marshaled.
//
// Returns:
//   - []byte: The JSON representation of the provided JSONBA slice.
//   - error: An error if the marshaling process fails.
func MarshalJSONBA(data JSONBA) ([]byte, error) {
	return json.Marshal(data)
}

// NewJSONBA creates a new JSONBA instance from the provided data.
//
// This function marshals the input 'data' into JSON format and then unmarshals it into a slice of map[string]interface{}. It returns the created JSONBA instance and any error encountered during the process.
//
// Parameters:
//   - data: interface{} - The data to be converted into JSONBA. It can be any data type.
//
// Returns:
//   - JSONBA: The created JSONBA instance, which is essentially a slice of map[string]interface{}.
//   - error: An error if the marshaling or unmarshaling process fails.
func NewJSONBA(data interface{}) (JSONBA, error) {
	dataJSON, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	var dataMap []map[string]interface{}
	if err := json.Unmarshal(dataJSON, &dataMap); err != nil {
		return nil, err
	}

	return JSONBA(dataMap), nil
}

// UnmarshalJSON unmarshals JSON data into the target interface{}.
//
// This function takes JSON data as input and unmarshals it into the provided target interface{}. If the input data is already a []byte, it directly unmarshals it; otherwise, it marshals the data into []byte first. It returns any error encountered during the unmarshaling process.
//
// Parameters:
//   - data: interface{} - The JSON data to unmarshal. It can be either a []byte or any other data type that can be marshaled into []byte.
//   - target: interface{} - A pointer to the type into which the JSON data will be unmarshaled.
//
// Returns:
//   - error: An error if the unmarshaling process fails.
//
// Example:
// Suppose you have JSON data in a byte slice named 'jsonData' and a struct named 'User', you can use UnmarshalJSON like this:
//
//	var user User
//	err := UnmarshalJSON(jsonData, &user)
//	if err != nil {
//	    fmt.Println("Error:", err)
//	    return
//	}
//
// This will unmarshal the JSON data in 'jsonData' into the 'user' struct.
func UnmarshalJSON(data interface{}, target interface{}) error {
	var jsonData []byte
	var err error

	// Check if data is already a []byte
	if bytesData, ok := data.([]byte); ok {
		jsonData = bytesData
	} else {
		// If data is not a []byte, marshal it to []byte
		jsonData, err = json.Marshal(data)
		if err != nil {
			return err
		}
	}

	// Unmarshal the JSON data into the target interface{}
	return json.Unmarshal(jsonData, target)
}

type JSONBA []map[string]interface{}

// Value converts the JSOBA value into a driver.Value for database storage.
//
// This method converts the JSOBA value into a string representation before returning it as a driver.Value. It's typically used for storing JSONB data in databases that support JSONB storage.
//
// Returns:
//   - driver.Value: A driver.Value representing the JSOBA value for database storage.
//   - error: An error if there's any issue during the conversion process.
//
// Note:
//   - This method internally uses the encoding/json package to marshal the JSOBA value into a string.
//   - Any errors during the conversion process will be returned as an error.
func (j JSONBA) Value() (driver.Value, error) {
	defer func() {
		if r := recover(); r != nil {
			log.Println("--------------JSOBA Value-----------------")
			err := fmt.Errorf("panic occurred: %v", r)
			log.Println(err)
			log.Println("-------------------------------")
		}
	}()
	valueString, err := json.Marshal(j)
	return string(valueString), err
}

// Scan populates the JSOBA value from a database driver.Value.
//
// This method takes a database driver.Value and populates the JSOBA value with the corresponding data. It's typically used when retrieving JSONB data from databases that support JSONB storage.
//
// Parameters:
//   - value: interface{} - The database driver.Value to be scanned into the JSOBA value.
//
// Returns:
//   - error: An error if there's any issue during the scanning process.
//
// Note:
//   - This method expects the database driver.Value to be a byte slice representing JSON data.
//   - Any errors during the scanning process will be returned as an error.
func (j *JSONBA) Scan(value interface{}) error {
	defer func() {
		if r := recover(); r != nil {
			log.Println("--------------JSONBA Scan-----------------")
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
		return fmt.Errorf("unexpected type for JSONBA: %T", value)
	}

	return nil
}
