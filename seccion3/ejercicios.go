package main

import "fmt"

type dinero int

var x dinero
var y int

func main() {
	fmt.Printf("%v\n", x)
	fmt.Printf("%T\n", x)

	x = 42
	fmt.Println(x)

	y = int(x + 58)
	fmt.Println(y)
	fmt.Printf("%T", y)
}
