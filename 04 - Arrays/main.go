package main

import "fmt"

func main() {
	// Primeira forma de criar lista estática.
	var arr1 [4]int
	arr1[0] = 1
	arr1[1] = 2
	arr1[2] = 3
	arr1[3] = 4
	fmt.Println(arr1)

	// Segunda forma de criar lista estática.
	arr2 := [4]int {5, 6, 7, 8}
	fmt.Println(arr2)

	// Slice
	slice := []int {11, 22, 33}
	fmt.Println(slice)

		// Inclusão de item dinamicamente.
		slice = append(slice, 44)
		fmt.Println(slice)

		// Usando intervalos.
		sli1 := slice[1:]
		fmt.Println("Intervalo 1: ", sli1)

		sli2 := slice[:2]
		fmt.Println("Intervalo 2: ", sli2)

		sli3 := slice[3:]
		fmt.Println("Intervalo 3: ", sli3)
}