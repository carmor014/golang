package main

import "fmt"

type persona struct {
	nombre   string
	apellido string
}

type agenteSecreto struct {
	persona
	lpm bool
}

func (a agenteSecreto) presentar() {
	fmt.Println("soy", a.nombre, a.apellido)
}

func (p persona) presentar() {
	fmt.Println("soy", p.nombre, p.apellido)
}

//decarar una interface
type humano interface {
	//solo agrego metodos
	presentar()
}

//puede recibir persona o un agenteSecreto
func bar(h humano) {
	switch h.(type) {
	case persona:
		fmt.Println("switch de persona: ", h.(persona).nombre)
	case agenteSecreto:
		fmt.Println("switch de persona: ", h.(agenteSecreto).nombre)
	}
	fmt.Println("fui pasado a la funcion bar", h)
}

func main() {
	as1 := agenteSecreto{
		persona: persona{
			nombre:   "Carlos",
			apellido: "Mora",
		},
		lpm: true,
	}
	as2 := agenteSecreto{
		persona: persona{
			nombre:   "Maria",
			apellido: "Guz",
		},
		lpm: true,
	}
	p := persona{
		nombre:   "Carito",
		apellido: "Gomez",
	}

	fmt.Println(as1)
	as1.presentar()
	as2.presentar()

	bar(as1)
	bar(as2)
	bar(p)
}
