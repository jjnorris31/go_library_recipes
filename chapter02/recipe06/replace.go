package main

import (
	"fmt"
	"strings"
)

const refString = "Mary had a little woah! little"
const refStringTwo = "lamb, lamb, lamb"

func main(){
	// ¿Qué representa el index?
	// Representa el número de reemplazos que deseas
	// -1 para reemplazar todas las apariciones
	out := strings.Replace(refString, "little", refStringTwo, -1)
	fmt.Println(out)

	// en este caso hay 2 reemplazos
	out = strings.Replace(refStringTwo, "lamb", "wolf", 2)
	fmt.Println(out)


}
