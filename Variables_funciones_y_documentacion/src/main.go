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

	// Operadores aritméticos
	a := 10
	b := 5

	fmt.Println("\nOperadores Aritméticos:")
	fmt.Println("a + b =", a+b)
	fmt.Println("a - b =", a-b)
	fmt.Println("a * b =", a*b)
	fmt.Println("a / b =", a/b)
	fmt.Println("a % b =", a%b)

	// Areas
	radius := 5
	area := Pi * float64(radius) * float64(radius)
	fmt.Println("\nÁrea de un círculo con radio", radius, "es", area)
	// Área de un cuadrado
	side := 4
	areaSquare := side * side
	fmt.Println("Área de un cuadrado con lado", side, "es", areaSquare)

	// Área de un triángulo
	base := 6
	height := 3
	areaTriangle := 0.5 * float64(base) * float64(height)
	fmt.Println("Área de un triángulo con base", base, "y altura", height, "es", areaTriangle)

	// Área de un trapecio
	base1 := 8
	base2 := 5
	heightTrap := 4
	areaTrapezoid := 0.5 * float64(base1+base2) * float64(heightTrap)
	fmt.Println("Área de un trapecio con bases", base1, "y", base2, "y altura", heightTrap, "es", areaTrapezoid)
}
