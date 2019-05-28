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

### [Using the default printer](examples/default)
```go
package main

import (
	"fmt"

	"github.com/chrusty/go-tableprinter"
)

type exampleType struct {
	Name           string
	Age            int
	FavouriteWords []string
	Tags           map[string]interface{}
	IsCrufty       bool
}

func main() {

	// Prepare some example data:
	example := exampleType{
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
Output:
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

### [Using a custom printer](examples/custom)
```go
package main

import (
	"github.com/chrusty/go-tableprinter"
)

type exampleType struct {
	Name           string
	Age            int
	FavouriteWords []string
	Tags           map[string]interface{}
	IsCrufty       bool
}

func main() {

	// Make a custom printer with the default values:
	printer := tableprinter.New().WithBorders(true)

	// Prepare some example data (this time a slice):
	examples := []exampleType{
		{
			Name:           "prawn",
			Age:            15248,
			FavouriteWords: []string{"Cruft", "Crufts", "Crufty"},
			Tags: map[string]interface{}{
				"crufty": false,
				"grumpy": true,
			},
			IsCrufty: false,
		},
		{
			Name:           "CruftLord",
			Age:            99999,
			FavouriteWords: []string{"CruftLord", "CruftMaster", "Darth Crufter"},
			Tags: map[string]interface{}{
				"crufty": true,
				"grumpy": false,
			},
			IsCrufty: true,
		},
	}

	// Use the custom printer to print the examples:
	printer.Print(examples)
}
```
Output:
```
+-------+---------------------------------------+----------+-----------+-------------------------------+
|  AGE  |            FAVOURITEWORDS             | ISCRUFTY |   NAME    |             TAGS              |
+-------+---------------------------------------+----------+-----------+-------------------------------+
| 15248 | [Cruft Crufts Crufty]                 | false    | prawn     | map[crufty:false grumpy:true] |
| 99999 | [CruftLord CruftMaster Darth Crufter] | true     | CruftLord | map[crufty:true grumpy:false] |
+-------+---------------------------------------+----------+-----------+-------------------------------+
```





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
