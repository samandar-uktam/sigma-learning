package solid

// This is a bad example because ...

// Let's imagine we have a multi-function machine that can print, scan and fax.
/*
	type Machine interface {
		Print(doc string)
		Scan(doc string)
		Fax(doc string)
	}


	// But, we have an old school printer that can only print.

	type OldPrinter struct{}

	func (o OldPrinter) Print() {
		fmt.Println("Printing...")
	}

	func (o OldPrinter) Scan() {
		panic("scan not supported")
	}

	func (o OldPrinter) Fax() {
		panic("fax not supported")
	}

	// Here we are implementing unnecessary methods for old school printer.
*/
