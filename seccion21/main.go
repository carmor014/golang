package main

import "fmt"

func main() {
	ca := make(chan int, 1) //pueda dejar en el canal un valor

	ca <- 42
	ca <- 43

	fmt.Println(<-ca)
}
