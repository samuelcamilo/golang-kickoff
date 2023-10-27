package helloworld

import (
	"fmt"
	"strings"
)

const messagePrefix string = "Hello, "

func Hello(s string) string {
	if strings.EqualFold(s, "") {
		return messagePrefix + "World!"
	}

	return "Hello, " + s + "!"
}

func main() {
	fmt.Println(Hello("Camilo"))
}
