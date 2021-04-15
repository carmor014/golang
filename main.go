package main

import "fmt"

type animal struct {
	nombre string
	patas  int
}

func main() {
	perro, mensaje := inicializarAnimal("mateo", 2)
	if mensaje != "" {
		fmt.Println("paso algo")
	}
	perro.cambioNombre("lola")
	perro.cambioPatas(4)
	fmt.Println(perro)

}

func inicializarAnimal(nombre string, patas int) (animalito *animal, string2 string) { //retorna puntero
	if patas > 4 {
		return nil, "su animal tiene mas de 4 patas"
	}
	animalito = &animal{
		patas:  patas,
		nombre: nombre,
	}
	return
}

func (a *animal) cambioNombre(nombre string) { //recibe puntero animal
	a.nombre = nombre
}

func (a *animal) cambioPatas(numeroPatas int) {
	a.patas = numeroPatas
}
