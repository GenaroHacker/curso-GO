// Una interfaz es un conjunto de metodos.
package main

import "fmt"

// La interfaz es una forma de ver las estructuras.

// Forma es una interfaz que dice:
//"Todo tipo o estructura que tenga el metodo area y retorne float64 es de tipo Forma"
type Forma interface {
	area() float64
}

type Circulo struct {
	x, y, radio float64
}

type Rectangulo struct {
	ancho, alto float64
}

// Recordar usar * para acceder a la estructura directamente cuando
//tenemos un puntero apuntando a la estructura.
//                                                   ↓ ↓ ↓
func (c *Circulo) area() float64 {
	return 3.14 * c.radio * c.radio
}

//                                                   ↓ ↓ ↓
func (r *Rectangulo) area() float64 {
	return r.ancho * r.alto
}

func main() {
	cir1 := Circulo{x: 0, y: 0, radio: 5}
	rec1 := Rectangulo{ancho: 10, alto: 5}

	//                                               ↑ ↑ ↑
	// Recordar usar & cuando el metodo requiere un puntero.
	formas := []Forma{&cir1, &rec1} // <-- Arreglo

	for _, i := range formas {
		fmt.Println(i.area())
	}
}
