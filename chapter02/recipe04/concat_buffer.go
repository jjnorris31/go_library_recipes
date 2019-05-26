package main

import (
	"bytes"
	"fmt"
)

func main() {
	strings := []string{"This ", "is ", "even ", "more ", "performant "}

	buffer := bytes.Buffer{}

	// no utilizar en todas las situaciones, sólo en las que lo ameriten
	// por ejemplo exportaciones de CSV en memoria
	for _, val := range strings {
		buffer.WriteString(val)
	}

	fmt.Println(buffer.String())
}
