package tableprinter

import "io"

var defaultTablePrinter *Printer

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
