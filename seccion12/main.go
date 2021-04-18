package main

import "fmt"

func main() {
	xi := []int{2, 3, 4, 5, 6, 7, 8, 9}
	x := sum(xi...)
	fmt.Println("El valor total almacenado en la varible es:", x)
}

func sum(x ...int) int {
	fmt.Println(x)
	fmt.Printf("%T\n", x)

	suma := 0
	for i, v := range x {
		suma += v
		fmt.Println("El valor en el indice", i, "suma", v, "altotal, quedando", suma)
	}
	return suma
}
