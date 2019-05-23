package main

import (
	"fmt"
	"regexp"
)

const refString = "Mary*had,a%little_lamb"

func main() {
	// regexp tiene la función MustCompile que a su vez puede usar Split para
	// separar un string con una expresión regular dada
	words := regexp.MustCompile("[*,%_]").Split(refString, -1)

	for idx, word := range words {
		fmt.Printf("Word %d is: %s\n", idx, word)
	}
}
