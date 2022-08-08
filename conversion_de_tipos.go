package main

import (
	"fmt"
	"strconv"
)

func main() {
	edad := "22"
	// La función strconv.Atoi convierte un string a un integer
	// Esta funcion retorna valores multiples
	// Usamos el guion bajo para recibir y descartar el valor de retorno que no nos interesa
	edad_int, _ := strconv.Atoi(edad)

	fmt.Println(edad_int + 10)
}
