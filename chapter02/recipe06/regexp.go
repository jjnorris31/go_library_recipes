package main

import (
	"fmt"
	"regexp"
)

const refString = "Mary had a little woah! little"

func main() {
	// expresiones regulares
	regex := regexp.MustCompile("l[a-z]+")
	// cada check reemplazará lo existente
	out := regex.ReplaceAllString(refString, "replacement")
	fmt.Println(out)
}
