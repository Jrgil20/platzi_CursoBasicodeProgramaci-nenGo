# Primer contacto con las Goroutines

## ¿Cómo funcionan las Goroutines en Go?

Las Goroutines son una característica fundamental de Go para gestionar la concurrencia. Permiten ejecutar funciones de forma concurrente dentro del mismo programa. Básicamente, una Goroutine es una función que se ejecuta en paralelo con otras Goroutines en el mismo espacio de memoria. Esto puede mejorar la eficiencia del programa si se utiliza correctamente.

Para activar una Goroutine, simplemente se utiliza la palabra clave `go` antes de una llamada a función. Esto indica al programa que debe ejecutarse de manera concurrente. Sin embargo, puede haber problemas si la función `main` finaliza antes que las otras Goroutines terminen de ejecutar, como no imprimir algún texto que esperamos.

## ¿Cómo podemos coordinar las Goroutines con WaitGroup?

El paquete `sync` en Go proporciona la estructura `WaitGroup`, que se utiliza para coordinar la ejecución de Goroutines. Un `WaitGroup` permite programar la espera para que una función `main` no termine hasta que todas las Goroutines hayan concluido.

Veamos cómo se implementa:

```go
package main

import (
    "fmt"
    "sync"
)

func imprimirTexto(texto string, wg *sync.WaitGroup) {
    defer wg.Done()
    fmt.Println(texto)
}

func main() {
    var wg sync.WaitGroup
    wg.Add(1)
    go imprimirTexto("Hello", &wg)
    wg.Add(1)
    go imprimirTexto("World", &wg)

    wg.Wait()
}
```

Aquí, `wg.Add(1)` se utiliza para indicar que una nueva Goroutine se ha añadido al grupo que el `WaitGroup` está esperando. `defer wg.Done()` informa que la Goroutine ha concluido, y `wg.Wait()` en el `main` asegura que el programa espera la terminación de todas las Goroutines añadidas.

## Uso de funciones anónimas con Goroutines y ¿cómo pueden recibir parámetros?

Las funciones anónimas son funciones que se definen en su lugar de uso, y son comunes cuando se desarrollan aplicaciones concurrentes con Goroutines en Go. Puedes definir y ejecutar una función anónima cuando desees realizar una tarea pequeña pero paralela.

```go
package main

import (
    "fmt"
    "time"
)

func main() {
    go func(texto string) {
        fmt.Println(texto)
    }("Adiós")

    time.Sleep(1 * time.Second)
}
```

En este ejemplo, hemos definido una función anónima que acepta un parámetro. Se ejecuta concurrentemente gracias a `go` y se le pasa el argumento "Adiós". Es importante notar que para que la ejecución concurrente de funciones anónimas se sincronice adecuadamente, se puede utilizar un `WaitGroup`, como vimos antes, o simplemente introducir un tiempo de espera temporal (no recomendable para grandes proyectos).

## ¿Por qué se prefieren los WaitGroups sobre Time.Sleep?

Usar `time.Sleep()` introduce retrasos explícitos en el código, lo cual no es una práctica eficiente. Si el programa se hace más complejo o exige mayor precisión temporal, esto podría llevar a recursos desperdiciados y problemas de rendimiento. En cambio, los `WaitGroup` proporcionan un control más preciso y eficiente sobre la sincronización entre distintas partes concurrentes del programa, lo cual es esencial en aplicaciones robustas y de gran escala.

## ¿Qué esperar de los Channels en Go?

Aunque no se exploraron en profundidad en esta clase, los Channels son otra herramienta poderosa para manejar la comunicación entre Goroutines. A diferencia de los `WaitGroup` que permiten esperar a que las rutinas se completen, los Channels permiten enviar y recibir datos entre Goroutines. Esto desbloquea la capacidad de crear estructuras de programación más complejas y eficientes que pueden ser más robustas y fáciles de mantener a largo plazo.

Es vital para todo programador de Go explorar las distintas herramientas de concurrencia y encontrar la que mejor se adapte a sus necesidades específicas.
