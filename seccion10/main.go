package main

import "fmt"

type persona struct {
	nombre   string
	apellido string
}

func main() {
	p1 := persona{
		nombre:   "Eduar",
		apellido: "tua",
	}
	p2 := persona{
		nombre:   "Carlos",
		apellido: "Mora",
	}
	fmt.Println(p1, p2)
	fmt.Println(p1.nombre)
	fmt.Println(p2.apellido)
}
