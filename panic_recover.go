package main

import "fmt"

func main() {
	// defer se ejecuta al final de la función
	defer salidaDeEmergencia()

	const Pi = 0 // <- equivocado
	if Pi != 3.14 {
		// panic() crashea la ejecución de la funcion
		// sirve para indicar fallos y detener la ejecución de la función
		// los argumentos pasados a panic seran recuperados por recover()
		panic("Pi no es 3.14!!")
	}

	fmt.Println("Ningun panic ocurrio") // <- no se ejecuta
}

func salidaDeEmergencia() {
	// recover devuelve el valor del argumento pasado a panic.
	// recover sirve para retomar el control cuando se produce un panic
	argumento_pasado_a_panic := recover()
	if argumento_pasado_a_panic != nil {
		fmt.Println("Recuperado de: ", argumento_pasado_a_panic)
	}
}
