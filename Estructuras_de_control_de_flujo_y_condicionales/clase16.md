# El uso de los keywords defer, break y continue

En Go, existen palabras clave que nos permiten controlar el flujo de ejecución de un programa. En esta clase, aprenderemos a utilizar `defer`, `break` y `continue`.

## ¿Cómo funcionan las palabras clave defer, break y continue en Go?

Go es un lenguaje potente y eficiente, ideal para el desarrollo de aplicaciones de alto rendimiento. Hoy aprenderemos a utilizar tres palabras clave que te ayudarán a escribir mejor tu código en Go: `defer`, `break`, y `continue`. Estas palabras clave son esenciales para gestionar la ejecución de tus funciones y ciclos de forma eficaz.

### ¿Qué es y cómo funciona defer?

La palabra clave `defer` permite demorar la ejecución de una función hasta que la función que la contiene se termine. Esta función se ejecutará justo antes de que la función principal termine su proceso. La palabra clave puede ser especialmente útil cuando necesitas realizar acciones finales como cerrar un archivo o desconectar una base de datos, asegurándote de que esas operaciones ocurran pase lo que pase.

#### Ejemplo de uso

Aquí hay un ejemplo básico sobre cómo `defer` cambia el orden de las operaciones en una función:

```go
package main

import "fmt"

func main() {
    defer fmt.Println("Hola")
    fmt.Println("Mundo")
}
```

Al ejecutar este código, observarás que aunque `defer fmt.Println("Hola")` aparece primero, se ejecuta después de `fmt.Println("Mundo")`. Esto demuestra que `defer` pospone su ejecución hasta el final del bloque de funciones.

#### ¿Cuándo utilizar defer?

Usa `defer` cuando necesites asegurarte de que ciertas operaciones siempre se ejecutan, por ejemplo:

- Cerrando conexiones a bases de datos para liberar recursos.
- Cerrando archivos después de leer o escribir en ellos.
- Limpiando recursos temporales, como memoria o registros.

Es crucial recordar que aunque puedes usar múltiples declaraciones `defer`, es una buena práctica limitarte a una por función para evitar complicaciones.

### ¿Cómo se utilizan continue y break en bucles for?

Ambas palabras clave, `continue` y `break`, se utilizan para controlar la ejecución de un ciclo `for`.

#### Uso de continue

`continue` se usa para omitir la ejecución del bloque actual y pasar al siguiente ciclo inmediatamente:

```go
for i := 0; i < 10; i++ {
    if i == 2 {
        fmt.Println("Es dos")
        continue
    }
    fmt.Println(i)
}
```

En este ejemplo, notarás que cuando `i` es igual a 2, el ciclo imprimirá "Es dos" y luego saltará directamente a la siguiente iteración, omitiendo `fmt.Println(i)` para esa vuelta.

#### ¿Cuándo utilizar continue?

Recurre a `continue` cuando, al cumplirse una condición especificada, quieras omitir el resto del bloque y seguir con la siguiente iteración. Es útil en casos como:

- Ignorar errores controlados en bucles.
- Evitar procesamiento adicional innecesario en ciertas condiciones.

### ¿Qué hace break en un bucle for?

La clave `break` se usa para salir completamente de un ciclo cuando se cumple una condición específica:

```go
for i := 0; i < 10; i++ {
    if i == 8 {
        fmt.Println("Break aquí")
        break
    }
    fmt.Println(i)
}
```

Aquí, cuando `i` es igual a 8, imprimirá "Break aquí" y saldrá del ciclo completamente, cesando toda ejecución adicional del bucle.

#### ¿Cuándo utilizar break?

`break` resulta útil cuando quieres detener la ejecución de un ciclo por completo, si una condición crítica o de finalización se cumple, por ejemplo:

- Al encontrar un elemento específico en un conjunto de datos.
- Si un error crítico ocurre y se requiere abortar la operación actual.

### ¿Qué necesitas saber antes de implementar estas palabras clave?

Entender el flujo y propósito de la ejecución de tus funciones y bucles `for` te permitirá aumentar la eficiencia y legibilidad de tu código en Go. Las palabras clave `defer`, `break`, y `continue` son herramientas poderosas para controlar el flujo y terminar tareas como cierre de archivos automáticamente; manejar errores y condiciones específicas, lo cual es esencial a medida que avanzas en tu aprendizaje y aplicación de Go.
