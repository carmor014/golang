package main

import "fmt"

func main() {
	ca := make(chan int, 1) //pueda dejar en el canal un valor

	ca <- 42

	fmt.Println(<-ca)
}
