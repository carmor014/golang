package main

import "fmt"

func main() {
	foo()
	bar("James")
	s1 := woo("Mooneypenny")
	fmt.Println(s1)
	x, y := saludar("Carlos", "Mora")
	fmt.Println(x)
	fmt.Println(y)
}

func foo() {
	fmt.Println("Hola desde foo")
}

func bar(s string) {
	fmt.Println("Hola,", s)
}

func woo(s string) string {
	return fmt.Sprint("Hola desde woo,", s)
}

func saludar(n string, a string) (string, bool) {
	p := fmt.Sprint(n, " ", a, " dice ")
	q := true
	return p, q
}
