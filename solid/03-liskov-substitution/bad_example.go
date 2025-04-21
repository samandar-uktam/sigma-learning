package solid

// This is a bad example because both rectangle and square can not use the Area() method in common.

/*
	type Rectangle struct {
		Width  float64
		Height float64
	}

	func (r Rectangle) Area() float64 {
		return r.Width * r.Height
	}

	type Square struct {
		Side float64
	}

	func (s Square) Area() float64 {
		return s.Side * s.Side
	}
*/

// usage example
/*
	func PrintArea(r Rectangle) {
		fmt.Println("Area:", r.Area())
	}

	func main() {
		r := Rectangle{Width: 10, Height: 5}
		s := Square{Side: 5}

		PrintArea(r)     // works
		PrintArea(s)     // compile-time error: Square is not Rectangle
	}
*/
