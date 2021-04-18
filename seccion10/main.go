package main

import "fmt"

type persona struct {
	nombre   string
	apellido string
	saborFav []string
}

func main() {
	p1 := persona{
		nombre:   "Carlos",
		apellido: "Mora",
		saborFav: []string{
			"chocolate",
			"arequipe",
			"chicle",
		},
	}
	fmt.Println(p1.nombre)
	fmt.Println(p1.apellido)
	for i, v := range p1.saborFav {
		fmt.Println("\t", i, v)
	}
}
