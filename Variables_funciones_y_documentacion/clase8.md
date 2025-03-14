# Paquete fmt: algo más que imprimir en consola

## ¿Qué es el paquete fmt en Go?

El paquete `fmt` en Go es una herramienta esencial para la interacción con la consola. Aunque es conocido principalmente por su capacidad de imprimir mensajes en la consola, el paquete `fmt` ofrece funciones adicionales que son cruciales para manejar la salida formateada en Go. En esta lección, nos enfocaremos en tres funciones principales: `Println`, `Printf` y `Sprintf`.

## ¿Cómo usar Println para añadir saltos de línea automáticamente?

### Uso de Println

La función `Println` es fundamental cuando necesitamos imprimir texto en la consola con un salto de línea automático al final. Esta función es ideal para imprimir mensajes ordenados sin preocuparse por manualmente introducir nuevas líneas.

```go
package main

import (
    "fmt"
)

func main() {
    helloMessage := "Hello"
    worldMessage := "World"
    fmt.Println(helloMessage, worldMessage)
}
```

Al ejecutar este código, `Println` agrega un salto de línea automáticamente. Este comportamiento puede ser comprobado duplicando la línea de impresión en el código y observando cómo cada mensaje se imprime en líneas separadas.

## ¿Cómo formatear cadenas y enteros con Printf?

### Funcionalidad de Printf

`Printf` es una función versátil que permite formatear cadenas e introducir variables dentro del texto mediante especificadores de formato como `%s` para strings y `%d` para enteros. Además, puedes añadir un salto de línea usando `\n` en tu cadena formateada.

```go
package main

import (
    "fmt"
)

func main() {
    nombre := "Platzi"
    cursos := 500
    fmt.Printf("%s tiene más de %d cursos\n", nombre, cursos)
}
```

### Buenas prácticas en el uso de tipos de datos

Cuando no estás seguro del tipo de dato que quieres imprimir, puedes usar `%v` como un especificador genérico. No obstante, es recomendable especificar el tipo para evitar errores y mejorar la claridad del código.

## ¿Cómo generar un string formateado sin imprimir con Sprintf?

### Uso de Sprintf

La función `Sprintf` genera una cadena formateada pero no la imprime directamente en la consola. En su lugar, la guarda en una variable, lo que permite una manipulación posterior más sencilla.

```go
package main

import (
    "fmt"
)

func main() {
    nombre := "Platzi"
    cursos := 500
    mensaje := fmt.Sprintf("%s tiene más de %d cursos", nombre, cursos)
    fmt.Println(mensaje)
}
```

En este caso, la variable `mensaje` contiene el texto formateado que luego se imprime con `Println`.

## ¿Cómo determinar el tipo de dato de una variable?

Un uso importante del paquete `fmt` es conocer el tipo de variable en tiempo de ejecución. Usamos `%T` para imprimir el tipo de dato.

```go
package main

import (
    "fmt"
)

func main() {
    helloMessage := "Hello"
    cursos := 500
    fmt.Printf("El tipo de variable de helloMessage es %T\n", helloMessage)
    fmt.Printf("El tipo de variable de cursos es %T\n", cursos)
}
```

## Recomendaciones y recursos adicionales

Con el paquete `fmt`, puedes optimizar significativamente la forma en que manejas la salida en consola. Aprovecha todos los especificadores de formato disponibles para una mejor gestión de tus datos. Al escribir tus programas, considera consultar la [documentación oficial de Golang](https://pkg.go.dev/fmt) para obtener una lista completa de todos los tipos de parseo disponibles.
