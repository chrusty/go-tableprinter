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
	if err := printer.Print(examples); err != nil {
		panic(err)
	}

	// Use the Marshal method to get bytes:
	tableBytes, _ := printer.Marshal(examples)
	fmt.Printf("\nThis table is %dB\n", len(tableBytes))
}
