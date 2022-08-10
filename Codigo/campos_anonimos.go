// Un campo anónimo es una estructura que no tiene nombre.
// Esto permite que los campos se hereden del struct incrustado.

package main

import "fmt"

type Persona struct {
	Nombre string
	Edad   int
}

type AgenteSecreto struct {
	Persona           // <-- campo anónimo
	LicenciaParaMatar bool
}

func main() {
	jaime := AgenteSecreto{
		Persona: Persona{
			Nombre: "James Bond",
			Edad:   23,
		},
		LicenciaParaMatar: true,
	}
	fmt.Println(jaime.Nombre)
	fmt.Println(jaime.Edad)
	fmt.Println(jaime.LicenciaParaMatar)
}

// Sobreescribir un atributo o metodo de un struct anonimo:
// En el caso de que una estructura tenga mas de un campo anonimo en su interior,
// y estos campos compartan algun atrubuto o metodo, se puede acceder al necesario
// con el nombre del campo anonimo.
// Ejemplo:
// en vez de usar
//		jaime.Nombre
// se usa
//		jaime.Persona.Nombre
