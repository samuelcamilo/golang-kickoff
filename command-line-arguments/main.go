// displays the different ways of working with command line arguments.
package main

import (
	"fmt"
	"os"
	"strings"
)

func getCountArgs() int {
	return len(os.Args)
}

func getAllWithLoopArgs() (s string) {
	var sep string
	for i := 0; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}

	return
}

// the range generate pair value. index value and element value.
func getAllWithRangeArgs() (s string) {
	var sep string
	for _, arg := range os.Args[0:] {
		s += sep + arg
		sep = " "
	}

	return
}

func getAllArgs() string {
	return strings.Join(os.Args[0:], " ")
}

func main() {
	fmt.Println("Arguments Count: ", getCountArgs())
	fmt.Println("1.1 - Arguments String: ", getAllWithLoopArgs())
	fmt.Println("1.2 - Arguments String: ", getAllWithRangeArgs())
	fmt.Println("1.3 - Arguments String: ", getAllArgs())
}
