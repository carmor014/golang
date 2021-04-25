package main

//goland:noinspection GoUnsortedImport
import (
	"fmt"
	"pruebaAnimalGo/26-seccion/gato"
)

type canino struct {
	nombre string
	edad   int
}

type felino struct {
	nombre string
	edad   int
}

func main() {
	c1 := canino{
		nombre: "Chespirito",
		edad:   perro.Edad(2),
	}
	fmt.Println(c1)

	f1 := felino{
		nombre: "sparki",
		edad:   gato.Edad(3),
	}
	fmt.Println(f1)
}
