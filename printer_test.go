package tableprinter_test

import (
	"bytes"
	"testing"

	"github.com/chrusty/go-tableprinter"
	"github.com/stretchr/testify/assert"
)

type testCase struct {
	expectedOutput string
	inputValue     interface{}
}

type nestedCruft struct {
	cruftiness float32
	name       string
}

type complexStructure struct {
	name     string
	crufts   []*nestedCruft
	weight   int
	crufty   *bool
	shitNuts map[string]interface{}
}

var (
	basicTests = map[string]testCase{
		"Basic string": {
			inputValue:     "cruft",
			expectedOutput: "  VALUE  \n+-------+\n  cruft  \n",
		},
		"Basic int": {
			inputValue:     5,
			expectedOutput: "  VALUE  \n+-------+\n      5  \n",
		},
		"Basic bool": {
			inputValue:     true,
			expectedOutput: "  VALUE  \n+-------+\n  true   \n",
		},
		"Basic pointer": {
			inputValue:     new(bool),
			expectedOutput: "  VALUE  \n+-------+\n  false  \n",
		},
	}

	mapTests = map[string]testCase{
		"Basic map": {
			inputValue: map[string]interface{}{
				"name":   "prawn_map",
				"age":    7654,
				"crufty": true,
			},
			expectedOutput: "  AGE  | CRUFTY |   NAME     \n+------+--------+-----------+\n  7654 | true   | prawn_map  \n",
		},
		"Basic map with pointers": {
			inputValue: map[string]interface{}{
				"name":   "prawn_map",
				"age":    new(int),
				"crufty": true,
			},
			expectedOutput: "  AGE | CRUFTY |   NAME     \n+-----+--------+-----------+\n    0 | true   | prawn_map  \n",
		},
	}

	sliceTests = map[string]testCase{
		"Slice of strings": {
			inputValue:     []string{"this", "is", "quite", "crufty"},
			expectedOutput: "  VALUE   \n+--------+\n  this    \n  is      \n  quite   \n  crufty  \n",
		},
		"Slice of structs": {
			inputValue: []struct {
				name   string
				age    int
				crufty bool
			}{
				{
					"prawn_struct_1",
					1000,
					true,
				},
				{
					"prawn_struct_2",
					2000,
					false,
				},
				{
					"prawn_struct_3",
					3000,
					true,
				},
			},
			expectedOutput: "  AGE  | CRUFTY |      NAME       \n+------+--------+----------------+\n  1000 | true   | prawn_struct_1  \n  2000 | false  | prawn_struct_2  \n  3000 | true   | prawn_struct_3  \n",
		},
		"Slice of pointers": {
			inputValue: []*struct {
				name   string
				age    int
				crufty bool
			}{
				{
					"prawn_struct_ptr_1",
					1000,
					true,
				},
				{
					"prawn_struct_ptr_2",
					2000,
					false,
				},
			},
			expectedOutput: "  AGE  | CRUFTY |        NAME         \n+------+--------+--------------------+\n  1000 | true   | prawn_struct_ptr_1  \n  2000 | false  | prawn_struct_ptr_2  \n",
		},
	}

	structTests = map[string]testCase{
		"Struct": {
			inputValue: struct {
				name   string
				age    int
				crufty bool
			}{
				"prawn_struct",
				8888,
				true,
			},
			expectedOutput: "  AGE  | CRUFTY |     NAME      \n+------+--------+--------------+\n  8888 | true   | prawn_struct  \n",
		},
		"Struct pointer": {
			inputValue: &struct {
				name   string
				age    int
				crufty bool
			}{
				"prawn_struct_pointer",
				9999,
				false,
			},
			expectedOutput: "  AGE  | CRUFTY |         NAME          \n+------+--------+----------------------+\n  9999 | false  | prawn_struct_pointer  \n",
		},
		"Struct with pointers": {
			inputValue: struct {
				name   string
				age    int
				crufty *bool
			}{
				"prawn_struct_pointer",
				9999,
				new(bool),
			},
			expectedOutput: "  AGE  | CRUFTY |         NAME          \n+------+--------+----------------------+\n  9999 | false  | prawn_struct_pointer  \n",
		},
	}
)

func TestTablePrinter(t *testing.T) {

	// Create a buffer for the output (so we can check what gets printed):
	outputBuffer := bytes.NewBufferString("")

	// Make a new TablePrinter:
	tablePrinter := tableprinter.New().WithOutput(outputBuffer)
	assert.NotNil(t, tablePrinter)

	// Test various types of input:
	testPrint(t, tablePrinter, outputBuffer, basicTests)
	testPrint(t, tablePrinter, outputBuffer, mapTests)
	testPrint(t, tablePrinter, outputBuffer, sliceTests)
	testPrint(t, tablePrinter, outputBuffer, structTests)
	testComplexStructure(t, tablePrinter, outputBuffer)
}

func testPrint(t *testing.T, tp *tableprinter.Printer, outputBuffer *bytes.Buffer, testCases map[string]testCase) {
	for name, tc := range testCases {

		// Reset the buffer:
		outputBuffer.Reset()

		// Run the test with its own name:
		t.Run(name, func(t *testing.T) {

			// Print the value:
			if err := tp.Print(tc.inputValue); err != nil {
				assert.NoError(t, err)
				return
			}

			// Compare the output:
			assert.Equal(t, tc.expectedOutput, outputBuffer.String())
		})
	}
}

func testComplexStructure(t *testing.T, tp *tableprinter.Printer, outputBuffer *bytes.Buffer) {

	testStructure := complexStructure{
		name:   "Complex cruft",
		crufty: new(bool),
		weight: 99,
		crufts: []*nestedCruft{
			{
				cruftiness: 33.3,
				name:       "cruft1",
			},
		},
		shitNuts: map[string]interface{}{
			"crufty":     true,
			"arse_balls": 55,
			"anus": nestedCruft{
				cruftiness: 66.6,
				name:       "cruft2",
			},
		},
	}

	tableprinter.Print(testStructure)
}
