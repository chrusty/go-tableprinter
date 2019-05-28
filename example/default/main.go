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
	fmt.Printf("\n\n")
}
