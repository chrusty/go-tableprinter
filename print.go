package tableprinter

import (
	"fmt"
)

// Print marshals an interface and prints it to the configured output:
func (p *Printer) Print(value interface{}) error {

	// Marshal the value to bytes:
	marshaledBytes, err := p.Marshal(value)
	if err != nil {
		return err
	}

	// Now print the marshaled bytes:
	if _, err := fmt.Fprint(p.output, string(marshaledBytes)); err != nil {
		return err
	}

	return nil
}
