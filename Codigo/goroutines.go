package main

import (
	"fmt"
	"time"
)

// test imprime los numeros del 0 al 4 en intervalos de 0.1 segundos.
func test() {
	var i int
	for i = 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Print(i, " ")
	}
}

func main() {
	go test() // la palabra reservada go ejecuta una funcion en un hilo independiente.
	go test()
	test()
	time.Sleep(100 * time.Millisecond) // da tiempo a terminar a los gorutines.
	// runtime.NumGoroutine() devuelve el numero de gorutines en ejecucion (incluida main).
}

//0 0 0 1 1 1 2 2 2 3 3 3 4 4 4
