package main

import "fmt"

type punto struct {
	x int
	y int
}

// al pasar el struct como puntero accedo a sus atributos de forma directa
// es recomendable pasar el struct como puntero para evitar crear una copia
// si NO paso el struct como puntero solo puedo leer sus atributos, no modificarlos
// esto es posible porque la instancia es una direccion de memoria
func (instancia *punto) definir_y_sumar(a, b int) int {
	instancia.x = a
	instancia.y = b
	return instancia.x + instancia.y
}

func main() {
	// new crea una instancia de un struct inicalizando sus atributos a cero
	// la instancia es una variable que contiene una direccion de memoria
	mi_punto := new(punto)
	fmt.Println(mi_punto)
	//&{0 0}
	suma := mi_punto.definir_y_sumar(5, 7)
	fmt.Println(mi_punto)
	//&{5 7}
	fmt.Println(suma)
	//12

}
