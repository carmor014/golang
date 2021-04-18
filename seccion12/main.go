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

func (a agenteSecreto) presentarse() {
	fmt.Println("soy", a.nombre, a.apellido)
}

func main() {
	as1 := agenteSecreto{
		persona: persona{
			nombre:   "Carlos",
			apellido: "Mora",
		},
		lpm: true,
	}
	fmt.Println(as1)
	as1.presentarse()
}
