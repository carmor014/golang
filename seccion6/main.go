package main

import "fmt"

func main() {
	if x := 42; x == 42 {
		fmt.Println("001")
	}

	y := 30
	if y == 30 {
		fmt.Println(y)
	}
	fmt.Println(y)
}
