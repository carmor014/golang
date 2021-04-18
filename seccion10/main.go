package main

import "fmt"

type persona struct {
	nombre   string
	apellido string
}
type agenteSecreto struct {
	persona
	licenciaParaMatar bool
}

func main() {
	as1 := agenteSecreto{
		persona: persona{
			nombre:   "Carlos",
			apellido: "Mora",
		},
		licenciaParaMatar: true,
	}
	fmt.Println(as1)
	fmt.Println(as1.persona, as1.nombre)
}
