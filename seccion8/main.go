package main

import (
	"fmt"
)

func main() {
	x := []int{1, 2, 3, 4, 5}
	fmt.Println(x)
	fmt.Println(x[0]) // cada uno tiene una posicion empezando desde 0
	fmt.Println(x[1])
	fmt.Println(x[2])
	fmt.Println(x[3])
	fmt.Println(x[4])

	for i, v := range x { //dos variables para este caso indice  valor
		fmt.Println(i, " ", v)
	}

}
