package main

import (
	"fmt"
	"strings"
)

const refString = "Mary*had,a%little_lamb"

func main() {

	// la función splitFunc es llamada para
	// cada runa en un string. Si la runa
	// equivale a cualquier cáracter en "*%,_"
	// refString se separará

	splitFunc := func(r rune) bool {
		return strings.ContainsRune("*%,_", r)
	}

	// el segundo argumento es una función que consume una runa
	// de un determinado string y regresa true si el string
	// debe separarse en este punto
	words := strings.FieldsFunc(refString, splitFunc)

	for idx, word := range words {
		fmt.Printf("Word %d is: %s\n", idx, word)
	}
}
