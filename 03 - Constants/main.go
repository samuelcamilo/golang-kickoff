package main

import "fmt"

// iota funciona com base no scopo...
const (
	first = iota
	second
)

const (
	third = iota
)

func main() {

	// Declaração de constante...
	const pi = 3.1415
	fmt.Println(pi)

	// Constante com "iota"
	fmt.Println(first, second, third)
}
