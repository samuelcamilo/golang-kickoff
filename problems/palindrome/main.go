package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func formatInputString(input string) (result string) {
	reg, _ := regexp.Compile("[-,\\s]")
	result = strings.ToLower(reg.ReplaceAllString(input, ""))
	
	return
}

func reverse(input string) (string) {
	r := []rune(input)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}

	return string(r)
}

func checkPalindrome(input string) (bool) {
	return input == reverse(input)
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter your text: ")
	input, _ := reader.ReadString('\n')
	
	if checkPalindrome(formatInputString(input)) {
		fmt.Println("Is a Palindrome")
	} else {
		fmt.Println("Isn´t a Palindrome")
	}
}