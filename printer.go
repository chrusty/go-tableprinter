package tableprinter

import (
	"fmt"
	"io"
	"os"

	"github.com/davecgh/go-spew/spew"
)

// Printer takes care of marshaling interfaces to text tables:
type Printer struct {
	borders       bool
	output        io.Writer
	sortedHeaders bool
	spewConfig    *spew.ConfigState
}

// New returns a new Printer, configured with default values:
func New() *Printer {

	spewConfig := spew.NewDefaultConfig()
	spewConfig.DisableCapacities = true
	spewConfig.DisableMethods = true
	spewConfig.DisablePointerAddresses = true
	spewConfig.DisablePointerMethods = true
	spewConfig.SortKeys = true
	spewConfig.SpewKeys = true

	return &Printer{
		borders:       false,
		output:        os.Stdout,
		sortedHeaders: true,
		spewConfig:    spewConfig,
	}
}

// WithBorders causes the printer to add borders to tables:
func (p *Printer) WithBorders(borders bool) *Printer {
	p.borders = borders
	return p
}

// WithOutput adds an output to the printer:
func (p *Printer) WithOutput(output io.Writer) *Printer {
	p.output = output
	return p
}

// WithSortedHeaders causes the printer to alphabetically sort columns by their headers:
func (p *Printer) WithSortedHeaders(sortedHeaders bool) *Printer {
	p.sortedHeaders = sortedHeaders
	return p
}

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

// Marshal turns an interface into a text table:
func (p *Printer) Marshal(value interface{}) ([]byte, error) {

	// Turn the value into a table:
	table, err := p.makeTable(value)
	if err != nil {
		return nil, err
	}

	return table.bytes(p.sortedHeaders, p.borders)
}
