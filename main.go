package main

import (
	"fmt"

	"github.com/iaraoz/go-desde-0/variables"
)

func main() {
	estado, texto := variables.ConveritoaTexto(1986)
	fmt.Println(estado)
	fmt.Println(texto)
}
