package main

import (
	"fmt"
	"strings"
)

const refString = "Mary had a little lamb Mary"

func main() {
	// un replacer también sirve para los mismos propósitos
	// sigue la estructura oldnew
	replacer := strings.NewReplacer("lamb", "wolf", "Mary", "Jack")
	// después utiliza la propiedad de Replace para cambiarlo en el string
	out := replacer.Replace(refString)
	fmt.Println(out)
}