# go-tableprinter
Print a formatted table from GoLang interfaces

## Features
* Handles structs / maps / slices / interfaces
* Colums are alphabetically ordered
* Requires no modification to existing data structures

## Usage
* Make a *Printer (`tp := tableprinter.New()`)
* Make a *Printer with a custom output (`tp := rablePrinter.New().WithOutput(os.Stderr)`)
* Print a table from an interface (`err := tp.Print(interface)`)
* Marshal a table from an interface if you'd rather do something else with the output (`tableBytes, err := tp.Marshal(interface)`)

### Todo
* ~~Handle generic interfaces (default)~~
* ~~Handle pointers~~
* ~~Handle structs~~
* ~~Handle maps~~
* ~~Handle slices of the above~~
