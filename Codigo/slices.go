package main

import "fmt"

func main() {
	// A diferencia de el arreglo o array, el slice no tiene un tamaño fijo.
	// Es un arreglo que puede cambiar de tamaño.
	slice := []int{1, 2}
	fmt.Println(slice)

	// La estructura de un slice contiene 3 datos:
	// - Puntero al arreglo
	// - Tamaño del arreglo al que apunta
	// - Capacidad del arreglo al que apunta

	// Podemos obtener el tamaño del slice usando la función len()
	fmt.Println(len(slice))

	// Otra forma de obtener un slice es a partir de un arreglo
	arreglo := [3]int{1, 2, 3}

	slice2 := arreglo[:]
	fmt.Println(slice2)
	// [1 2 3]

	slice3 := arreglo[1:]
	fmt.Println(slice3)
	// [2 3]

	slice4 := arreglo[:2]
	fmt.Println(slice4)
	// [1 2]

}
