package main

import (
	"fmt"
	"io/ioutil"
)

// Este programa lee su propio contenido y lo imprime.
func main() {
	// ReadFile devuelve el contenido en bytes y un error si ocurre
	file_data, err := ioutil.ReadFile("./leer_archivos.go") // leemos el archivo

	if err != nil { // si hay error lo imprimimos
		fmt.Println(err)
	}

	fmt.Println(string(file_data)) // imprimimos el contenido del archivo
}
