package main

import "fmt"

type animal struct {
	nombre string
	patas  int
}

func main() {
	perro := inicializarAnimal("mateo", 2)
	gato := inicializarAnimal("jhon", 3)
	cambioNombre(perro)
	cambioPatas(gato)
	fmt.Println(perro)
	fmt.Println(gato)
}

func inicializarAnimal(nombre string, patas int) *animal { //retorna puntero
	return &animal{
		patas:  patas,
		nombre: nombre,
	}
}

func cambioNombre(mascota *animal) { //recibe puntero animal
	mascota.nombre = "koke"
}

func cambioPatas(bestia *animal) {
	bestia.patas = 4
}
