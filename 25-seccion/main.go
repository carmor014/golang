//Package main provides math solutions
package main

type persona struct {
	nombre string
}

//Esta es la prueba de la documentacion de Go
func main() {

}

//funcion para sumar
func Sum(xi ...int) int {
	sum := 0
	for _, v := range xi {
		sum += v
	}
	return sum
}
