package oop

import "fmt"

// As there are no classes in golang, we do not have direct inheritance feature.
// But, we can embed the `struct` types by composition.

type Animal struct{}

func (a Animal) Eat() {
	fmt.Println("Eating...")
}

type Dog struct {
	Animal // embedded
}

func (d Dog) Bark() {
	fmt.Println("Woof!")
}

// The Dog object inherits all the methods and features of Animal struct.
