package main

import "fmt"

func main() {
	// Criação e atribuição de valores para a variável...
	var firstName string = "Camilo"
	fmt.Println(firstName)

	// Criação e atribuição de endereço de memória para ptr...
	ptr := &firstName
	fmt.Println(ptr, *ptr)

	*ptr = "Samuel"
	fmt.Println(firstName, *ptr)

	// Atribuição de valores para a variável...
	firstName = "Beatriz"
	fmt.Println(ptr, *ptr)
}
