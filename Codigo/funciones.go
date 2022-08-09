package main

import "fmt"

func Saludar() {
	// defer mueve la ejecución de la declaración hasta el final de la función.
	defer fmt.Println("Mundo")
	fmt.Println("Hola")
}

//            ↓ input    ↓ output
func Calcular(x, y int) (int, int) {
	return x + y, x - y
}

// Una funcion puede ser un parametro. (funcion anonimas)
// Debemos declarar la cantidad y tipo de parametros que va a recibir la funcion-argumento.
// Asi como la cantidad y tipo de parametros que va a retornar la funcion.
func ImprimirResultados(funcion_argumento func(x, y int) (int, int), x, y int) {
	fmt.Println(funcion_argumento(x, y))
}

// Una funcion puede retornar una funcion.
func RetornarFuncion() func(x, y int) int {
	return func(x, y int) int {
		fmt.Println(x + y)
		return x + y
	}
}

func main() {
	suma, resta := Calcular(10, 20)
	fmt.Println(suma, resta)
	//30 -10

	Saludar()
	//Hola
	//Mundo

	// Una funcion puede ser una variable.
	// Puedo declarar una funcion sin parentesis para referenciar la logica sin ejecutarla.
	x := Saludar
	x()
	//Hola
	//Mundo
	ImprimirResultados(Calcular, 5, 10)
	//15 -5
	RetornarFuncion()(5, 7)
	//12
}
