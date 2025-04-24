package adapter

import "fmt"

// This is the source object where our adapter calls
type Printer interface {
	Print(message string)
}

type LegacyPrinter struct{}

func (l *LegacyPrinter) PrintText(text string) {
	fmt.Println("Legacy printer:", text)
}

// This is our adapter where it calls the legacy adapter and
// plays a key role as a bridge between legacy printer object and our app
type PrinterAdapter struct {
	OldPrinter *LegacyPrinter
}

func (pa *PrinterAdapter) Print(message string) {
	// Adapt the interface
	pa.OldPrinter.PrintText(message)
}

func main() {
	legacyPrinter := &LegacyPrinter{}
	adapter := &PrinterAdapter{OldPrinter: legacyPrinter}

	PrintSomething(adapter)
}

func PrintSomething(p Printer) {
	p.Print("Hello, Adapter Pattern!")
}
