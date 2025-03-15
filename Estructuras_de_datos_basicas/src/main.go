package main

import "fmt"

func main() {

	// Ejemplo de array
	// Un array es una estructura de datos que almacena una colección de elementos del mismo tipo
	// Los elementos de un array se almacenan en posiciones consecutivas de memoria
	// La longitud de un array es fija y se define en el momento de su creación
	// La longitud de un array es parte de su tipo
	var array [5]int
	array[0] = 1
	array[1] = 2
	array[2] = 3
	array[3] = 4
	array[4] = 5

	fmt.Println("Array:", array)

	// Ejemplo de slice
	// Un slice es una estructura de datos que almacena una colección de elementos del mismo tipo
	// Los elementos de un slice se almacenan en posiciones consecutivas de memoria
	// La longitud de un slice es variable y se puede modificar en tiempo de ejecución
	// Un slice es una vista de un array
	// Un slice tiene tres propiedades: un puntero a un array, una longitud y una capacidad
	// La longitud de un slice es el número de elementos que contiene
	// La capacidad de un slice es el número de elementos que puede contener sin tener que redimensionarse
	// Un slice se define con la siguiente sintaxis: slice := []tipo{elemento1, elemento2, ..., elementoN}
	slice := []int{1, 2, 3, 4, 5}
	fmt.Println("Slice:", slice)

	// Agregar un elemento al slice
	slice = append(slice, 6)
	fmt.Println("Slice después de append:", slice)

	// Recorrer un slice con range
	for i, v := range slice {
		fmt.Printf("Índice: %d, Valor: %d\n", i, v)
	}

	// Ejemplo de map
	// Un map es una estructura de datos que almacena una colección de pares clave-valor
	// Las claves de un map son únicas
	// Los elementos de un map no están ordenados
	// Un map se define con la siguiente sintaxis: map := map[tipoClave]tipoValor{clave1: valor1, clave2: valor2, ..., claveN: valorN}
	mapa := map[string]int{
		"uno":    1,
		"dos":    2,
		"tres":   3,
		"cuatro": 4,
		"cinco":  5,
	}
	fmt.Println("Mapa:", mapa)

	// Recorrer un map con range
	for clave, valor := range mapa {
		fmt.Printf("Clave: %s, Valor: %d\n", clave, valor)
	}
}
