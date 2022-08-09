package main

import "fmt"

func main() {
	//& y *
	//&variable = direccion de memoria de la variable
	//*puntero = valor de la direccion de memoria

	x := 0

	// “&” Dice, “obtener la dirección de la variable que sigue”
	direccion := &x
	fmt.Println(direccion)
	// 0xc000010098

	// Si tengo un puntero apuntando a una variable, usando * puedo
	//acceder a la variable directamente.
	fmt.Println(x)
	// 0
	*direccion = 1
	fmt.Println(x)
	// 1

}
