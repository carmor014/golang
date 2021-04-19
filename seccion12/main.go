package main

import "fmt"

func main() {
	n1 := factorial(4)
	fmt.Println(n1)

	n2 := cicloFac(4)
	fmt.Println(n2)
}

func factorial(n int) int {
	//caso base cuando n = 0
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

//Haciendolo con ciclos
func cicloFac(n int) int {
	total := 1
	for ; n > 0; n-- {
		total *= n
	}
	return total
}
