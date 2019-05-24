package tableprinter

import (
	"fmt"
	"reflect"
)

// Marshal turns an interface into a text table:
func (p *Printer) Marshal(value interface{}) ([]byte, error) {
	kindOfValue := reflect.TypeOf(value).Kind()

	// Take a different approach depending on the type of data that was provided:
	switch kindOfValue {

	// Maps get turned into a single-row table:
	case reflect.Map:
		table, err := p.marshalMapValue(value)
		if err != nil {
			return nil, err
		}
		return table.bytes(p.borders)

	// For pointers we just recurse on their interface:
	case reflect.Ptr:
		return p.Marshal(reflect.ValueOf(value).Elem().Interface())

	// Slices get turned into a multi-row table:
	case reflect.Slice:
		return nil, fmt.Errorf("Unsupported type (%s)", kindOfValue)

	// Structs get turned into a single-row table:
	case reflect.Struct:
		table, err := p.marshalStructValue(value)
		if err != nil {
			return nil, err
		}
		return table.bytes(p.borders)

	// The default is a one-row one-column table:
	default:
		table, err := p.marshalBasicValue(value)
		if err != nil {
			return nil, err
		}
		return table.bytes(p.borders)
	}
}

// marshalBasicValue turns an interface into a single column in a single row:
func (p *Printer) marshalBasicValue(value interface{}) (*table, error) {
	var table = new(table)
	var row = make(tableRow)

	// Just add the one value:
	table.addHeader(defaultFieldName)
	row[defaultFieldName] = fmt.Sprintf("%v", value)
	table.addRow(row)

	return table, nil
}

// marshalMapValue turns a map into a single-row table:
func (p *Printer) marshalMapValue(value interface{}) (*table, error) {
	var table = new(table)
	var row = make(tableRow)

	// Turn the interface into a map[string]interface{}:
	assertedMap, ok := value.(map[string]interface{})
	if !ok {
		return nil, ErrAssertion
	}

	// Add the map fields to the table:
	for fieldName, fieldValue := range assertedMap {
		table.addHeader(fieldName)

		// Handle pointers differently:
		if reflect.TypeOf(fieldValue).Kind() == reflect.Ptr {
			row[fieldName] = fmt.Sprintf("%v", reflect.ValueOf(fieldValue).Elem())
		} else {
			row[fieldName] = fmt.Sprintf("%v", fieldValue)
		}
	}

	// Add the row to the table:
	table.addRow(row)

	return table, nil
}

// marshalStructValue turns a struct into a single-row table:
func (p *Printer) marshalStructValue(value interface{}) (*table, error) {
	var table = new(table)
	var row = make(tableRow)

	// Reflect the value to gain access to its elements:
	reflectedType := reflect.TypeOf(value)
	reflectedValue := reflect.ValueOf(value)

	// Add the struct fields to the table:
	for i := 0; i < reflectedType.NumField(); i++ {
		fieldName := reflectedType.Field(i).Name
		table.addHeader(fieldName)

		// Handle pointers differently:
		if reflectedType.Field(i).Type.Kind() == reflect.Ptr {
			fieldValue := reflectedValue.Field(i).Elem()
			row[fieldName] = fmt.Sprintf("%v", fieldValue)
		} else {
			fieldValue := reflectedValue.Field(i)
			row[fieldName] = fmt.Sprintf("%v", fieldValue)
		}
	}

	// Add the row to the table:
	table.addRow(row)

	return table, nil
}
