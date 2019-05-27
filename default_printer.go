package tableprinter

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
