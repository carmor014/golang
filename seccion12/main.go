package main

import "fmt"

func main() {
	foo()

	func() {
		fmt.Println("La funcion es anonima")
	}()

	func(x int) {
		fmt.Println("La funcion es anonima con unparametro", x)
	}(42)
}

func foo() {
	fmt.Println("foo se ejecuto")
}
