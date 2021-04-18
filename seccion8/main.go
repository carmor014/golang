package main

import (
	"fmt"
)

func main() {
	et := []string{"Eduar", "Tua", "Crossfit", "Baseball", "Montañismo"}
	fmt.Println(et)
	jl := []string{"Jacinto", "Lopez", "Correr", "Nadar", "Bailar"}
	fmt.Println(jl)

	vp := [][]string{et, jl}
	fmt.Println(vp)
}
