package main

import "fmt"

func main() {
	// Ciclo for se compone de 3 elementos: INICIO, CONDICION, INCREMENTO (todos ellos opcionales):
	// 1. Inicialización (Declaración de la variable/iterador)
	// 2. Condición (Expresión booleana que determina si se repite el ciclo)
	// 3. Incremento (Instrucción que se ejecuta después de cada iteración)
	for i := 0; i < 10; i++ {
		//continue: salta a la siguiente iteración
		if i <= 3 {
			continue
		}

		fmt.Println(i)

		// break: para salir del ciclo
		if i == 7 {
			break
		}
	}
}
