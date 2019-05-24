package tableprinter

import "fmt"

var (
	ErrAssertion = fmt.Errorf("Unable to assert value")
	ErrNoData    = fmt.Errorf("No data to render")
)
