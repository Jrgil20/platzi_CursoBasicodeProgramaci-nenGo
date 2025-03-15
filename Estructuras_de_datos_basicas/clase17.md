# Arrays y Slices

## ¿Qué son los arrays en Go?

Go, al igual que otros lenguajes de programación, maneja estructuras de datos como los arrays y slices. Un array es una colección de elementos del mismo tipo y tiene dimensiones fijas que se definen al momento de su creación. Aquí te mostramos cómo se declaran:

```go
var array [4]int
```

Cuando se inicializa un array en Go, los elementos se llenan con el valor cero del tipo de dato especificado. En este caso, los enteros son inicializados con 0.

## ¿Cuál es la naturaleza inmutable de los arrays?

Los arrays en Go son inmutables, lo que significa que no podemos cambiar su tamaño una vez definidos. Sin embargo, podemos modificar los valores de los elementos existentes:

```go
array[0] = 1
array[1] = 2
```

Esto significa que podemos asignar nuevos valores a los índices específicos del array, pero no podemos agregar nuevos elementos.

## ¿Cómo funcionan los slices en Go?

A diferencia de los arrays, los slices en Go son dinámicos y permiten crecer en tamaño. Al igual que los arrays, los slices están compuestos de elementos del mismo tipo, pero no requieren la especificación del tamaño al ser declarados:

```go
slice := []int{0, 1, 2, 3, 4, 5, 6}
```

Cuando ejecutamos `len(slice)` obtenemos el número de elementos del slice, y `cap(slice)` nos indica la capacidad total para alojar elementos.

## ¿Qué es el slicing y cómo se usa?

El slicing es una poderosa herramienta para manipular arrays y slices. Nos permite seleccionar un subconjunto de elementos:

Imprimir el elemento en una posición específica:

```go
fmt.Println(slice[0]) // Resultado: 0
```

Imprimir elementos desde el inicio hasta una posición (excluido el índice final):

```go
fmt.Println(slice[:3]) // Resultado: [0 1 2]
```

Imprimir una parte específica del slice:

```go
fmt.Println(slice[2:4]) // Resultado: [2 3]
```

Desde una posición específica hasta el final:

```go
fmt.Println(slice[4:]) // Resultado: [4 5 6]
```

## ¿Cómo podemos añadir elementos a los slices?

La flexibilidad de los slices permite agregar elementos adicionales con el método `append`:

```go
slice = append(slice, 7)
```

Para añadir una lista completa de elementos, se aplica `append` con una sintaxis especial que descompone el slice insertado:

```go
extraSlice := []int{8, 9, 10}
slice = append(slice, extraSlice...)
```

Este uso de `...` literalmente expande los elementos del `extraSlice` y los añade individualmente al slice original.

En resumen, Go ofrece herramientas eficaces y flexibles para manejar arrays y slices, con sus propias características y usos. Los arrays son más eficientes en situaciones donde los datos no cambiarán, mientras que los slices proporcionan la flexibilidad necesaria para crecer dinámicamente.
