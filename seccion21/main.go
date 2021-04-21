package main

import "fmt"

func main() {
	ca := make(chan<- int, 2) //pueda dejar en el canal un valor

	ca <- 42
	ca <- 43

	//Escritura del canal por esto falla
	fmt.Println(<-ca)
	fmt.Println(<-ca)
	fmt.Println("------------")
	fmt.Printf("%T\n", ca)
}
