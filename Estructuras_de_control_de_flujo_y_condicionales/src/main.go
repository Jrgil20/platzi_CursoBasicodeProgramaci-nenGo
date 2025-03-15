package main

import "fmt"

func main() {
	// Ejemplo de ciclo for
	fmt.Println("Ciclo for:")
	// el ciclo for se compone de tres partes: inicialización, condición y post
	// la inicialización se ejecuta una sola vez al inicio del ciclo
	// la condición se evalúa antes de cada iteración del ciclo
	// el post se ejecuta al final de cada iteración del ciclo
	// el ciclo for se ejecuta mientras la condición sea verdadera
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	// Ejemplo de ciclo for while
	fmt.Println("Ciclo for while:")
	// el ciclo for puede simular un ciclo while
	// para simular un ciclo while se omite la inicialización y el post
	// el ciclo for se ejecuta mientras la condición sea verdadera
	j := 0
	for j < 5 {
		fmt.Println(j)
		j++
	}

	// Ejemplo de ciclo for forever
	fmt.Println("Ciclo for forever:")
	// el ciclo for puede simular un ciclo forever
	// para simular un ciclo forever se omite la condición
	// el ciclo for se ejecuta indefinidamente
	k := 0
	for {
		if k >= 5 {
			break
			// el ciclo for puede terminar con la instrucción break
		}
		fmt.Println(k)
		k++
	}
}
