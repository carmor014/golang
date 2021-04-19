package main

import (
	"fmt"
)

func main() {

	f := func() {
		fmt.Println("Mi primera expresión función")
	}
	f()

	g := func(x int) {
		fmt.Println("El año del descubrimiento de América fue:", x)
	}
	g(1492)
}
