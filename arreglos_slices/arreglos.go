package arreglos_slices

import (
	"fmt"
)

var tabla [10]int
var tabla_slices []int = []int{10, 0, 59}
var arreglos [10]int = [10]int{6, 98, 21, 674, 18, 36, 78, 9}

func MuestroArreglos() {
	for i := 0; i < len(tabla); i++ {
		tabla[i] = i + 1
	}
	for i := 0; i < len(tabla); i++ {
		println(tabla[i])
	}
}

func MuestroSlice() {
	fmt.Println(tabla_slices)
	porcion := arreglos[3:] //slice desde la posicion 3
	fmt.Println(porcion)
	porcion2 := arreglos[:5] // desde la posicion 0 hasta la 5
	fmt.Println(porcion2)
	porcion3 := arreglos[6:7] //slice desde la posicion 6 hasta las 7
	fmt.Println(porcion3)
}

func Capacidad() {
	elementos := make([]int, 5, 20)
	fmt.Printf("Largo %d , capacidad %d ", len(elementos), cap(elementos))
}
