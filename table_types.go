package tableprinter

import (
	"fmt"
	"reflect"
)

func (p *Printer) makeTable(value interface{}) (*table, error) {

	// Take a different approach depending on the type of data that was provided:
	switch reflect.TypeOf(value).Kind() {

	// Maps get turned into a single-row table:
	case reflect.Map:
		return p.tableFromMapValue(value)

	// For pointers we just recurse on their interface:
	case reflect.Ptr:
		return p.makeTable(reflect.ValueOf(value).Elem().Interface())

	// Slices get turned into a multi-row table:
	case reflect.Slice:
		return p.tableFromSliceValue(value)

	// Structs get turned into a single-row table:
	case reflect.Struct:
		return p.tableFromStructValue(value)

	// The default is a one-row one-column table:
	default:
		return p.tableFromBasicValue(value)
	}
}

// tableFromBasicValue turns an interface into a single column in a single row:
func (p *Printer) tableFromBasicValue(value interface{}) (*table, error) {
	var table = new(table)
	var row = make(tableRow)

	// Just add the one value:
	table.addHeader(defaultFieldName)
	row[defaultFieldName] = fmt.Sprintf("%v", value)
	table.addRow(row)

	return table, nil
}

// tableFromMapValue turns a map into a single-row table:
func (p *Printer) tableFromMapValue(value interface{}) (*table, error) {
	var table = new(table)
	var row = make(tableRow)

	// Turn the value into a map[string]interface{}:
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

// tableFromSliceValue turns a slice into a multi-row table:
func (p *Printer) tableFromSliceValue(value interface{}) (*table, error) {
	var table = new(table)

	// Reflect the value to gain access to its elements:
	reflectedValue := reflect.ValueOf(value)

	// Turn each entry into a table (with a row that we can take):
	for i := 0; i < reflectedValue.Len(); i++ {
		tempTable, err := p.makeTable(reflectedValue.Index(i).Interface())
		if err != nil {
			return nil, err
		}

		// Add the new row and headers to our table:
		table.headers = tempTable.headers
		table.addRow(tempTable.rows[0])
	}

	return table, nil
}

// tableFromStructValue turns a struct into a single-row table:
func (p *Printer) tableFromStructValue(value interface{}) (*table, error) {
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
