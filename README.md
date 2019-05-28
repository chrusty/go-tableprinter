# go-tableprinter
Print a formatted table from GoLang interfaces. This can be useful if you're building a CLI, or just prefer a more human-readable interpretation of your data.

## Features
* Handles structs / maps / slices / interfaces
* Colums are alphabetically ordered by default (this can be disabled if you prefer)
* Tables can optionally have borders (disabled by default)
* Requires no modification to existing data structures
* Nil values are listed as `<nil>`

## Limitations
* Currently unable to print unexported struct fields, similar to JSON or YAML (listed as `<unexported>`)

## Usage

You can use this package to print tables either by using the default *Printer, or making your own (with specific config).

### Default

* Print a table from an interface:
```go
package main

import (
	"fmt"

	"github.com/chrusty/go-tableprinter"
)

type ExampleType struct {
	Name           string
	Age            int
	FavouriteWords []string
	Tags           map[string]interface{}
	IsCrufty       bool
}

func main() {

	// Prepare some example data:
	example := ExampleType{
		Name:           "prawn",
		Age:            15248,
		FavouriteWords: []string{"Cruft", "Crufts", "Crufty"},
		Tags: map[string]interface{}{
			"crufty": true,
			"grumpy": true,
		},
		IsCrufty: true,
	}

	// Disable sorting then print using the default printer:
	tableprinter.SetSortedHeaders(false)
	tableprinter.Print(example)
	fmt.Printf("\n\n")

	// Re-enable sorting then print using the default printer:
	tableprinter.SetSortedHeaders(true)
	tableprinter.Print(example)
	fmt.Printf("\n\n")

	// See what happens when we print a map:
	tableprinter.Print(example.Tags)
	fmt.Printf("\n\n")

	// See what happens when we print a slice:
	tableprinter.Print(example.FavouriteWords)
	fmt.Printf("\n\n")

	// Set borders then print using the default printer:
	tableprinter.SetBorder(true)
	tableprinter.Print(example)
}
```
Result:
```
  NAME  |  AGE  |    FAVOURITEWORDS     |             TAGS             | ISCRUFTY
+-------+-------+-----------------------+------------------------------+----------+
  prawn | 15248 | [Cruft Crufts Crufty] | map[crufty:true grumpy:true] | true


   AGE  |    FAVOURITEWORDS     | ISCRUFTY | NAME  |             TAGS
+-------+-----------------------+----------+-------+------------------------------+
  15248 | [Cruft Crufts Crufty] | true     | prawn | map[crufty:true grumpy:true]


  CRUFTY | GRUMPY
+--------+--------+
  true   | true


  VALUE
+--------+
  Cruft
  Crufts
  Crufty


+-------+-----------------------+----------+-------+------------------------------+
|  AGE  |    FAVOURITEWORDS     | ISCRUFTY | NAME  |             TAGS             |
+-------+-----------------------+----------+-------+------------------------------+
| 15248 | [Cruft Crufts Crufty] | true     | prawn | map[crufty:true grumpy:true] |
+-------+-----------------------+----------+-------+------------------------------+
```

### *Printer
* Make a *Printer (`tp := tableprinter.New()`)
* Make a *Printer with borders (`tp := tableprinter.New().WithBorders(true)`)
* Make a *Printer with borders and unsorted columns (`tp := tableprinter.New().WithBorders(true).WithSortedHeaders(false)`)
* Make a *Printer with a custom output (`tp := tablePrinter.New().WithOutput(os.Stderr)`)
* Print a table from an interface (`err := tp.Print(interface)`)
* Marshal a table from an interface if you'd rather do something else (other than printing) with the output (`tableBytes, err := tp.Marshal(interface)`)





## History

### ToDo
* Optional row numbering

### 0.3.0
* Optionally sort columns by headers

### 0.2.0
* Use go-spew to format values
* Provide a default tableformatter
* Make Borders optional

### 0.1.0
* Handle generic interfaces (default)
* Handle pointers
* Handle structs
* Handle maps
* Handle slices of the above
