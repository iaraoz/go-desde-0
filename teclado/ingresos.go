package teclado

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var numero1 int
var numero2 int
var leyenda string
var err error

func IngresoNumeros() {
	fmt.Println("Ingreso numero 1:")
	keyboard := bufio.NewScanner(os.Stdin)
	if keyboard.Scan() {
		numero1, err = strconv.Atoi(keyboard.Text())
		if err != nil {
			panic("El dato ingresado es incorrecto" + err.Error())
		}

	}

	fmt.Println("Ingreso el String :")
	keyboard = bufio.NewScanner(os.Stdin)
	if keyboard.Scan() {
		numero2, err = strconv.Atoi(keyboard.Text())
		if err != nil {
			panic("El dato ingresado es incorrecto" + err.Error())
		}

	}

	fmt.Println("Ingreso leyenda :")
	keyboard = bufio.NewScanner(os.Stdin)
	if keyboard.Scan() {
		leyenda = keyboard.Text()

	}
	fmt.Println(leyenda, numero1*numero2)

}
