package main

import "fmt"

func main() {

	// La función copy, copia elementos de un slice de origen en un slice de destino.
	// Copy devuelve el número de elementos copiados, que es el mínimo de len(origen) y len(destino).

	// En el ejemplo se crea una copia de un slice con el doble de largo del original
	//y el triple de su capacidad.
	slice := []int{1, 2, 3, 4}
	copia := make([]int, len(slice)*2, cap(slice)*3)
	copy(copia, slice) // <-- devuelve 4

	fmt.Println(slice)
	// [1 2 3 4]
	fmt.Println(copia)
	// [1 2 3 4 0 0 0 0]
	fmt.Println(cap(copia))
	// 12
}
