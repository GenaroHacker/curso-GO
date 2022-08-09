package main

import (
	"fmt"
)

func main() {
	var edad int
	fmt.Println("Coloca tu edad: ")
	// fmt.Scanln() se usa para solicitar datos al usuario, es equivalente a input() en Python
	fmt.Scanf("%d\n", &edad)
	// Los verbos se utilizan para indicar el tipo de variable y cómo se debe formatear.
	// El verbo más simple es %v, que formatea un valor en un formato predeterminado.
	// Todos los verbos en https://pkg.go.dev/fmt
	// %d es para un integer en base 10
	fmt.Printf("Mi edad es %d\n", edad)
}
