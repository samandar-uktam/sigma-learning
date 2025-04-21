package solid

import "fmt"

// This is a good example because each interface is specific. Any object can implement it.
// e.g: Old school printer can only print, so it uses only printer interface.
// e.g: Photo copier can scan and print, so it implements Printer and scanner Interface methods based on its needs

type Printer interface {
	Print()
}

type Scanner interface {
	Scan()
}

type Faxer interface {
	Fax()
}

type OldPrinter struct{}

func (p OldPrinter) Print() {
	fmt.Println("Printing...")
}

type PhotoCopier struct{}

func (p PhotoCopier) Print() {
	fmt.Println("Printing...")
}

func (p PhotoCopier) Scan() {
	fmt.Println("Scanning...")
}
