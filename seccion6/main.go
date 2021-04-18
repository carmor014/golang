package main

import "fmt"

func main() {

	//for init; condition; post {}
	for i := 0; i <= 10; i++ {
		fmt.Printf("El ciclo externo %d\n", i)
		for j := 0; j < 3; j++ {
			fmt.Printf("\t\tEl ciclo interno %d\n", j)
		}
	}

	//for con una sola condicion es infinito
	k := 11
	for k < 10 {
		fmt.Println(k)
	}

	//como salir de un fr infinito
	m := 5
	for {
		if m > 4 {
			break
		}
		fmt.Println(m)
	}
}