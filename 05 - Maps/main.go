package main

import "fmt"

func main() {
	m := map[string]int{"foo": 25, "fii":2}
	fmt.Println(m)
	fmt.Println(m["foo"])

	m["foo"] = 30
	fmt.Println(m)

	delete(m, "foo")
	fmt.Println(m)
}