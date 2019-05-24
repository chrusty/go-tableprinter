package tableprinter

import (
	"io"
	"os"
)

// Printer takes care of marshaling interfaces to text tables:
type Printer struct {
	borders bool
	output  io.Writer
}

// New returns a new Printer, configured with default values:
func New() *Printer {
	return &Printer{
		output: os.Stdout,
	}
}

// WithBorders causes the printer to add borders to tables:
func (p *Printer) WithBorders() *Printer {
	p.borders = true
	return p
}

// WithOutput adds an output to the printer:
func (p *Printer) WithOutput(output io.Writer) *Printer {
	p.output = output
	return p
}
