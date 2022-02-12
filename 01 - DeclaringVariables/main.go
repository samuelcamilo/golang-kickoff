package main

import (
	"fmt"
)

func main() {

	// Primeira forma de definir variáveis...
	var i int
	i = 23
	fmt.Println(i)

	// Segunda forma de definir variáveis...
	var f float32 = 3.14
	fmt.Println(f)

	// Terceira forma de definir variáveis...
	firstName := "Camilo"
	fmt.Println(firstName)

	// Definição de números complexos...
	complexNumber := complex(1, 4)
	fmt.Println(complexNumber)
}
