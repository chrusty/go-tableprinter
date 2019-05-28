package tableprinter_test

import (
	"bytes"
	"fmt"
	"testing"
	"time"

	"github.com/chrusty/go-tableprinter"
	"github.com/stretchr/testify/assert"
)

type testCase struct {
	expectedOutput string
	inputValue     interface{}
}

type nestedCruft struct {
	Cruftiness float32
	Name       string
}

func (n *nestedCruft) String() string {
	return fmt.Sprintf("%s: (Cruftiness: %v)", n.Name, n.Cruftiness)
}

type complexStructure struct {
	Name         string
	Started      time.Time
	Cruft        nestedCruft
	NestedCruft  *nestedCruft
	Crufts       []*nestedCruft
	Weight       int
	Crufty       *bool
	CruftMap     map[string]interface{}
	privateField string
}

var (
	testTime, _ = time.Parse(time.RFC3339, "2019-05-29T12:19:20Z")

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
		"Basic time": {
			inputValue:     testTime.UTC(),
			expectedOutput: "              VALUE              \n+-------------------------------+\n  2019-05-29 12:19:20 +0000 UTC  \n",
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
				Name   string
				Age    int
				Crufty bool
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
				Name   string
				Age    int
				Crufty bool
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
				Name   string
				Age    int
				Crufty bool
			}{
				"prawn_struct",
				8888,
				true,
			},
			expectedOutput: "  AGE  | CRUFTY |     NAME      \n+------+--------+--------------+\n  8888 | true   | prawn_struct  \n",
		},
		"Struct pointer": {
			inputValue: &struct {
				Name   string
				Age    int
				Crufty bool
			}{
				"prawn_struct_pointer",
				9999,
				false,
			},
			expectedOutput: "  AGE  | CRUFTY |         NAME          \n+------+--------+----------------------+\n  9999 | false  | prawn_struct_pointer  \n",
		},
		"Struct with pointers": {
			inputValue: struct {
				Name   string
				Age    int
				Crufty *bool
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
	tablePrinter := tableprinter.New().WithBorders(false).WithOutput(outputBuffer).WithSortedHeaders(true)
	assert.NotNil(t, tablePrinter)

	// Test various types of input:
	testPrint(t, tablePrinter, outputBuffer, basicTests)
	testPrint(t, tablePrinter, outputBuffer, mapTests)
	testPrint(t, tablePrinter, outputBuffer, sliceTests)
	testPrint(t, tablePrinter, outputBuffer, structTests)
	testComplexStructure(t, tablePrinter, outputBuffer)

	// Test some calls to the default printer:
	testDefaultPrinter(t)
}

func testPrint(t *testing.T, tp *tableprinter.Printer, outputBuffer *bytes.Buffer, testCases map[string]testCase) {
	for name, tc := range testCases {

		// Run the test with its own name:
		t.Run(name, func(t *testing.T) {

			// Reset the buffer:
			outputBuffer.Reset()

			// Print the value:
			if err := tp.Print(tc.inputValue); err != nil {
				assert.NoError(t, err)
				return
			}

			// Compare the output:
			assert.Equal(t, tc.expectedOutput, outputBuffer.String())

			// Marshal as bytes:
			tableBytes, err := tp.Marshal(tc.inputValue)
			if err != nil {
				assert.NoError(t, err)
				return
			}
			assert.Equal(t, len(outputBuffer.Bytes()), len(tableBytes))
		})
	}
}

func testComplexStructure(t *testing.T, tp *tableprinter.Printer, outputBuffer *bytes.Buffer) {
	t.Run("Complex structure", func(t *testing.T) {

		testStructure := complexStructure{
			Name:    "Complex cruft",
			Crufty:  new(bool),
			Started: testTime.UTC(),
			Weight:  99,
			Cruft: nestedCruft{
				Cruftiness: 99.99,
				Name:       "cruft5",
			},
			Crufts: []*nestedCruft{
				{
					Cruftiness: 33.3,
					Name:       "cruft1",
				},
			},
			CruftMap: map[string]interface{}{
				"cruft_bool": true,
				"cruft_int":  55,
				"cruft_struct": nestedCruft{
					Cruftiness: 66.6,
					Name:       "cruft2",
				},
			},
		}

		// Reset the buffer:
		outputBuffer.Reset()

		// Print the value:
		err := tp.Print(testStructure)
		assert.NoError(t, err)

		// Compare the output:
		assert.Equal(t, "             CRUFT            |                                 CRUFTMAP                                  |            CRUFTS            | CRUFTY |     NAME      | NESTEDCRUFT |            STARTED            | WEIGHT | PRIVATEFIELD  \n+-----------------------------+---------------------------------------------------------------------------+------------------------------+--------+---------------+-------------+-------------------------------+--------+--------------+\n  cruft5: (Cruftiness: 99.99) | map[cruft_bool:true cruft_int:55 cruft_struct:cruft2: (Cruftiness: 66.6)] | [cruft1: (Cruftiness: 33.3)] | false  | Complex cruft | <nil>       | 2019-05-29 12:19:20 +0000 UTC |     99 | <unexported>  \n", outputBuffer.String())
	})
}

func testDefaultPrinter(t *testing.T) {
	outputBuffer := bytes.NewBufferString("")

	t.Run("Print nil to default printer", func(t *testing.T) {
		err := tableprinter.Print(nil)
		assert.Error(t, err)
	})

	t.Run("Print to default printer", func(t *testing.T) {
		tableprinter.SetOutput(outputBuffer)
		tableprinter.SetBorder(true)
		err := tableprinter.Print("cruft")
		assert.NoError(t, err)
		assert.Equal(t, "+-------+\n| VALUE |\n+-------+\n| cruft |\n+-------+\n", outputBuffer.String())
	})

	t.Run("Marshal with default printer", func(t *testing.T) {
		tableprinter.SetBorder(true)
		marshaledBytes, err := tableprinter.Marshal("more cruft")
		assert.NoError(t, err)
		assert.Equal(t, "+------------+\n|   VALUE    |\n+------------+\n| more cruft |\n+------------+\n", string(marshaledBytes))
	})
}
