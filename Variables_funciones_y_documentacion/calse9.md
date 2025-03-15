# Uso de Funciones

## ¿Cómo optimizar tu código en Go con funciones?

Las funciones son una herramienta fundamental en cualquier lenguaje de programación, y Go no es la excepción. Al utilizarlas, puedes transformar el código repetitivo en uno más limpio, conciso, y fácil de mantener. Además, las funciones no solo favorecen tu trabajo, sino que también facilitan el entendimiento y mejora por otros programadores. En esta sección, exploraremos cómo usar funciones y las funciones anónimas en Go.

### ¿Cómo definir una función en Go?

Para definir una función en Go, utilizamos la palabra clave `func`, seguida del nombre de la función y paréntesis que indican si recibe parámetros:

```go
func nombreDeLaFunción() {
    // Cuerpo de la función
}
```

Aquí algunos aspectos a tener en cuenta:

- El nombre de la función debe describir su propósito.
- Si el nombre tiene más de una palabra, la segunda y siguientes deben comenzar con mayúscula.

### Ejemplo: Imprimir mensajes

Veamos cómo definir una función simple para imprimir mensajes, evitando la repetición de código:

```go
func imprimirMensaje(mensaje string) {
    fmt.Println(mensaje)
}

func main() {
    imprimirMensaje("Hola Mundo")
    imprimirMensaje("Hola Mundo dos")
    imprimirMensaje("Hola Mundo tres")
}
```

Al llamar `imprimirMensaje`, observamos que reutilizamos la misma lógica con diferentes mensajes.

### ¿Cómo aceptar múltiples parámetros en funciones?

Puedes definir funciones que acepten múltiples parámetros de diferentes tipos:

```go
func tripleArgumento(a, b int, c string) {
    fmt.Println(a, b, c)
}
```

Es una buena práctica en Go combinar parámetros de tipos iguales en una misma declaración.

### Ejemplo: Imprimir múltiples tipos

Podemos llamar a `tripleArgumento` con distintos valores:

```go
tripleArgumento(1, 2, "Hola")
```

Este enfoque permite manejar datos variados de manera eficiente.

### ¿Cómo hacer que una función retorne valores en Go?

En Go, las funciones pueden retornar valores, lo cual es útil para cálculos o procesamiento de datos que regresarán algún resultado:

```go
func retornarDoble(valor int) int {
    return valor * 2
}
```

### Ejemplo: Retornar un cálculo

Podemos utilizar `retornarDoble` para calcular y mostrar el doble de un número:

```go
func main() {
    resultado := retornarDoble(4)
    fmt.Println("El doble es:", resultado) // Salida: "El doble es: 8"
}
```

Esta función es sencilla pero demuestra cómo retornar valores desde una función.

### ¿Cómo manejar múltiples valores de retorno?

Las funciones en Go también pueden retornar múltiples valores, algo que puede ser muy útil en diferentes escenarios:

```go
func dividir(y, z int) (int, int) {
    cociente := y / z
    resto := y % z
    return cociente, resto
}
```

### Ejemplo: Uso de múltiples valores de retorno

```go
cociente, resto := dividir(10, 3)
fmt.Println("Cociente:", cociente, "Resto:", resto)
```

Si no necesitas uno de los valores, puedes omitirlo utilizando el identificador en blanco `_`:

```go
_, resto := dividir(10, 3)
```

### ¿Cuáles son las prácticas recomendadas al usar funciones en Go?

- **Nombres claros y descriptivos**: Elige nombres que reflejen el propósito de la función.
- **Evita repeticiones**: Utiliza funciones para encapsular lógica repetitiva.
- **Parámetros y retornos claros**: Declara de forma óptima los tipos de datos para claridad y eficiencia.
- **Evita variables no usadas**: Usa el identificador en blanco `_` para ignorar valores que no necesitas.