package main

import (
	"fmt"
	"strings"
)

const refString = "Mary_had a little_lamb"

func main() {
	words := strings.Split(refString, "_")
	for idx, words := range words {
		fmt.Printf("Word %d is: %s\n", idx, words)
	}
}
