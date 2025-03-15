package main

import "fmt"

// Nodo representa un nodo en una lista enlazada simple
type Nodo struct {
	valor     int
	siguiente *Nodo
}

// imprimirLista imprime todos los valores de la lista enlazada
func imprimirLista(n *Nodo) {
	for n != nil {
		fmt.Println(n.valor)
		n = n.siguiente
	}
}

func main() {
	// Crear nodos
	nodo1 := &Nodo{valor: 1}
	nodo2 := &Nodo{valor: 2}
	nodo3 := &Nodo{valor: 3}

	// Enlazar nodos
	nodo1.siguiente = nodo2
	nodo2.siguiente = nodo3

	// Imprimir lista
	imprimirLista(nodo1)
}
