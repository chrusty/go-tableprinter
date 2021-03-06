package tableprinter

import (
	"bytes"
	"sort"
	"strings"

	"github.com/olekukonko/tablewriter"
)

const (
	// defaultFieldName is the column header for individual values that have no field name:
	defaultFieldName     = "value"
	nilFieldValue        = "<nil>"
	spewPointerString    = "<*>"
	unexportedFieldValue = "<unexported>"
)

// tableRow is a map of fields which make up a row:
type tableRow map[string]string

// table is an in-memory representation of a table:
type table struct {
	headers      []string
	rows         []tableRow
	maxRowLength int
}

// setField sets a named field with a given value:
func (r tableRow) setField(field, value string) {
	r[field] = strings.ReplaceAll(value, spewPointerString, "")
}

// addHeader adds a header field:
func (t *table) addHeader(header string) {
	t.headers = append(t.headers, header)
}

// addRow appends a new row to our list:
func (t *table) addRow(row tableRow) {
	t.rows = append(t.rows, row)
}

// bytes renders a table as bytes:
func (t *table) bytes(sortedHeaders, borders bool) ([]byte, error) {

	// Make sure we actually have some data:
	if len(t.rows) == 0 {
		return nil, ErrNoData
	}

	// Create a buffer for the output (so we can collect what gets printed):
	tableBuffer := bytes.NewBuffer(nil)

	// Use a tablewriter:
	tw := tablewriter.NewWriter(tableBuffer)

	// Sort the headers:
	if sortedHeaders {
		sort.Strings(t.headers)
	}

	// Add the headers:
	tw.SetHeader(t.headers)

	// Tables without borders:
	tw.SetBorder(borders)

	tw.SetAutoWrapText(false)

	// Append the rows:
	for _, row := range t.rows {
		tw.Append(t.sortRow(row))
	}

	// Render the table:
	tw.Render()

	// Just return whatever was rendered to the buffer:
	return tableBuffer.Bytes(), nil
}

// sortRow returns a row in the corrent order (according to the header):
func (t *table) sortRow(row tableRow) []string {
	var sortedRow []string

	// Add the row fields in the same order as the headers:
	for _, header := range t.headers {
		sortedRow = append(sortedRow, row[header])
	}

	return sortedRow
}
