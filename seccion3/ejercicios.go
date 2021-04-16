package main

import "fmt"

type dinero int

var x dinero

func main() {
	fmt.Printf("%v\n", x)
	fmt.Printf("%T\n", x)
	x = 42
	fmt.Println(x)
}
