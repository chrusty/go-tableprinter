package tableprinter

import "io"

var defaultTablePrinter *Printer

// init establishes the default printer (which can be used without having to instantiate and maintian a *Printer in-code):
func init() {
	defaultTablePrinter = New()
}

// Print marshals an interface and prints it to the configured output:
func Print(value interface{}) error {
	return defaultTablePrinter.Print(value)
}

// Marshal turns an interface into a text table:
func Marshal(value interface{}) ([]byte, error) {
	return defaultTablePrinter.Marshal(value)
}

// SetBorder configures the default printer with a borders:
func SetBorder(borders bool) {
	defaultTablePrinter.borders = borders
}

// SetOutput configures the default printer with a specified output:
func SetOutput(output io.Writer) {
	defaultTablePrinter.output = output
}

// SetSortedHeaders configures the default printer to sort columns by their headers:
func SetSortedHeaders(sortedHeaders bool) {
	defaultTablePrinter.sortedHeaders = sortedHeaders
}
