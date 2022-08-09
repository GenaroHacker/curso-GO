package main

import "fmt"

func main() {
	// La función integrada make() asigna e inicializa un objeto de tipo:
	// 1. Slice: es como una matriz, pero su tamaño no es fijo.
	//Puede especificar la longitud y la capacidad de un segmento cuando lo crea.
	// 2. Map: una colección de pares clave-valor.
	// 3. Chan: un channel es una forma de comunicarse entre gorutines.

	// Crear un slice de enteros de tamaño 5 y capacidad 10
	slice := make([]int, 5, 10)

	// Un slice se compone de 3 elementos:
	// 1. Puntero: una referencia al espacio de memoria donde está almacenado.
	// 2. Longitud: la cantidad de elementos que contiene el slice.
	// 3. Capacidad: la cantidad de elementos que puede contener el slice.

	// len devuelve la longitud de un slice
	fmt.Println(len(slice))
	// cap devuelve la capacidad de un slice
	fmt.Println(cap(slice))

	// Añadir un elemento al final del slice
	slice = append(slice, 2)
	// Exceder la capacidad del slice es posible, pero no es efiiciente. Ya que
	//append crea un nuevo slice con una nueva capacidad.

	fmt.Println(slice)
}
