//  El programa solicita un input al usuario mediante una
// funcion anonima que se ejecuta en una goroutine. Una vez
// que el usuario ingresa un mensaje, la goroutine lo envía
// a travas de un canal y vuelve a solicitar un nuevo mensaje.
//
//  En paralelo, la goroutine main espera recibir un mensaje
// por medio del canal. Cuando recibe un mensaje, lo imprime y
// vuelve a esperar a recibir otro mensaje.
package main

import "fmt"

func main() {
	/* Un canal es como un túnel que conecta goroutines.
	Para crear un canal se usa la palabra reservada "make"
	y se especifica qué tipo de datos se enviarán a través de él. */
	mi_canal := make(chan string)

	// Para que una goroutine pueda usar un canar debe recibirlo como parámetro.
	go func(mi_canal chan string) {
		for {
			var mensaje string
			fmt.Scan(&mensaje)  // Leemos el input del teclado.
			mi_canal <- mensaje // Enviamos el valor a través del canal.
		}
	}(mi_canal)

	for {
		msg := <-mi_canal // Recibimos el valor a través del canal.
		// Cuando una instruccion espera recibir un valor del canal
		// NO se ejecuta hasta que lo reciba, bloqueando la goroutine.

		fmt.Println("El mensaje es: " + msg)
	}
}
