package main

import "fmt"

/*
Operadores relacionales
	== igual a
	< menor que
	> mayor que
	<= menor o igual que
	>= mayor o igual que
	!= diferente de
Operadores lógicos
	&& and
	|| or
	! not (ejemplo: !true = false)
*/

func main() {
	x := 10
	y := 10
	if x == y {
		fmt.Println("x es igual a y")
	} else {
		fmt.Println("x no es igual a y")
	}
}
