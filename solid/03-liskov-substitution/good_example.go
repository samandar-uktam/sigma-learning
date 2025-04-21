package solid

// This is a good example because we are using an interface to make the Area() function in common for any shape object.

/*
	type Shape interface {
		Area() float64
	}

	type Rectangle struct {
		Height float64
		Width  float64
	}

	func (r Rectangle) Area() float64 {
		return r.Height * r.Width
	}

	type Square struct {
		Side float64
	}

	func (r Square) Area() float64 {
		return r.Side * r.Side
	}

	func PrintArea(s Shape) {
		fmt.Println("Area: ", s.Area())
	}
*/

// usage example
/*
	func main() {
		r := Rectangle{
			Height: 2.0,
			Width:  4.0,
		}
		s := Square{
			Side: 3.0,
		}
		PrintArea(s)
		PrintArea(r)
	}
*/
