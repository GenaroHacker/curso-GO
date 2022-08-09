// Un struct es un tipo de dato que representa un valor compuesto.
// Se usa para agrupar datos relacionados.
// Es como una clase, pero no puede tener funciones.
package main

import "fmt"

type punto struct {
	x int
	y int
}

func main() {
	var p punto

	fmt.Println(p)
	// {0 0}

	p.x = 10
	p.y = 20

	fmt.Println(p)
	// {10 20}

	fmt.Printf("%T\n", p) // (imprime el tipo de dato)
	// main.punto
	fmt.Printf("%T\n", p.x)
	// int

}
