package main

import "fmt"

func main() {
	// Cuando un arreglo se inicializa de esta manera obtiene un valor cero para cada elemento
	// El valor cero de los enteros es el 0
	// El valor cero de los flotantes es el 0.0
	// El valor cero de los booleanos es el false
	// El valor cero de las cadenas es el ""
	var arreglo [5]int
	// [0 0 0 0 0]
	fmt.Println(arreglo)

	arreglo_declarado := [5]int{1, 2, 3, 4, 5}
	fmt.Println(arreglo_declarado)
	// [1 2 3 4 5]

	arreglo_parcialmente_declarado := [5]int{1, 2}
	fmt.Println(arreglo_parcialmente_declarado)
	// [1 2 0 0 0]

	// Usar len para obtener el largo del arreglo
	fmt.Println(len(arreglo))

	// Para acceder a un elemento se usa la notacion [indice]
	arreglo[0] = 1
	fmt.Println(arreglo)

	// Tambien podemos declarar arreglos multidimensionales
	var matriz [3][2]int
	// [[0 0] [0 0] [0 0]]

	matriz[2][1] = 1
	fmt.Println(matriz)
	// [[0 0] [0 0] [0 1]]

}
