package variables

import (
	"fmt"
	"strconv"
	"time"
)

var Nombre string
var Estado bool
var Sueldo float32
var Fecha time.Time

func RestoVariable() {
	Nombre = "Pedro"
	Estado = true
	Sueldo = 2400.66
	Fecha = time.Now()
	fmt.Println(Nombre)
	fmt.Println(Estado)
	fmt.Println(Sueldo)
	fmt.Println(Fecha)
}

func ConveritoaTexto(numero int) (bool, string) {
	var texto string
	texto = strconv.Itoa(numero)
	return true, texto

}
