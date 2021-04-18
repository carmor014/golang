package main

import (
	"fmt"
)

func main() {
	switch {
	case 2 == 4:
		fmt.Println("Esta no se debe imprimir2")
	case 3 == 3:
		fmt.Println("Imprime")
		fallthrough
	default:
		fmt.Println("Este es default")
	}

	n := "Pera"
	switch n {
	case "Manzana", "Pera", "Mango":
		fmt.Println("Varias frutas")
	case "M":
		fmt.Println("m")
	case "Q":
		fmt.Println("Esta es la q")
	default:
		fmt.Println("Esta es default")
	}
}
