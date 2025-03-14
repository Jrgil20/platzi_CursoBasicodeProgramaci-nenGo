package main

import "fmt"

// Constantes
const Pi = 3.14
const Greeting = "Hello, World!"

func main() {
	// Variables
	var age int = 30
	var name string = "John"
	var isStudent bool = true

	// Datos mixtos
	var mixedData = []interface{}{Pi, Greeting, age, name, isStudent}

	// Imprimir valores
	fmt.Println("Constantes:")
	fmt.Println("Pi:", Pi)
	fmt.Println("Greeting:", Greeting)

	fmt.Println("\nVariables:")
	fmt.Println("Age:", age)
	fmt.Println("Name:", name)
	fmt.Println("Is Student:", isStudent)

	fmt.Println("\nDatos Mixtos:")
	for _, data := range mixedData {
		fmt.Println(data)
	}
}
