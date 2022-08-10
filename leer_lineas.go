package main

import (
	"bufio"
	"fmt"
	"os"
)

// Este programa lee su propio codigo y itera sus lineas
func main() {
	archivo, error := os.Open("./leer_lineas.go") // recibimos el archivo y el error

	if error != nil { // imprimimos el error si existe
		fmt.Println(error)
	}

	scanner := bufio.NewScanner(archivo) // creamos un nuevo scanner para leer el archivo
	for scanner.Scan() {                 // iteramos las lineas
		linea := scanner.Text()
		fmt.Println(linea)
	}
}
