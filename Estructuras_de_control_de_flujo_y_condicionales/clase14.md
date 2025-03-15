# El condicional if en Go

## ¿Qué son las sentencias condicionales en programación?

Las sentencias condicionales son fundamentales en la programación, ya que permiten ejecutar bloques de código específicos basados en condiciones dadas. Estas sentencias permiten que el programa tome decisiones, dirigidas por los valores que manejan. Una de las más comunes en la mayoría de lenguajes es la sentencia if. Vamos a explorar su funcionamiento en el lenguaje de programación Go.

## ¿Cómo funciona la sentencia if en Go?

La estructura básica de un if en Go se utiliza para verificar si una condición es verdadera. Si lo es, se ejecuta el conjunto de instrucciones asociadas a ese if. Por ejemplo, podemos verificar si una variable es igual a cierto valor:

```go
valorUno := 1
valorDos := 2

if valorUno == 1 {
    fmt.Println("Es uno")
} else {
    fmt.Println("No es uno")
}
```

En este ejemplo, si `valorUno` es igual a 1, se imprimirá "Es uno". Si no, se ejecutará el bloque else, imprimiendo "No es uno".

## ¿Cómo aplicar puertas lógicas en sentencias condicionales?

Las puertas lógicas son un poderoso componente en las sentencias condicionales, permitiendo combinar múltiples condiciones. En Go, las puertas lógicas Y (`&&`) y OR (`||`) son comúnmente usadas.

### Puerta lógica AND

Ambas condiciones deben ser verdaderas para que el bloque de código se ejecute:

```go
if valorUno == 1 && valorDos == 2 {
    fmt.Println("Es verdad")
}
```

Aquí, "Es verdad" se imprimirá solo si `valorUno` es igual a 1 y `valorDos` es igual a 2.

### Puerta lógica OR

Solo una de las condiciones debe ser verdadera para ejecutar el bloque de código:

```go
if valorUno == 0 || valorDos == 2 {
    fmt.Println("Es verdad")
}
```

En este caso, "Es verdad" se imprimirá porque `valorDos` sí es igual a 2, aunque `valorUno` no es 0.

## ¿Cómo manejar errores con sentencias if?

Las sentencias if son útiles para manejar errores en Go, especialmente cuando se involucran conversiones o parsing. Un caso común es convertir una cadena de texto a un número entero utilizando el paquete `strconv`.

```go
import "strconv"
import "log"

valor, err := strconv.Atoi("53")
if err != nil {
    log.Fatal("Error:", err)
} else {
    fmt.Println("Valor convertido:", valor)
}
```

Aquí, si la conversión falla, la variable `err` será distinta de `nil`, desencadenando un mensaje de log y terminando la ejecución del programa. Si no hay error, imprimirá el valor convertido.

## ¿Qué desafíos puedes intentar resolver?

Para practicar las sentencias condicionales en Go, aquí tienes dos desafíos:

1. Crear una función que determine si un número es par o impar.

    ```go
    func esPar(numero int) bool {
        if numero%2 == 0 {
            return true
        }
        return false
    }
    ```

    Puedes llamar a esta función pasando un número como argumento y verificar si es par o impar:

    ```go
    numero := 4
    if esPar(numero) {
        fmt.Println("El número es par")
    } else {
        fmt.Println("El número es impar")
    }
    ```

2. Desarrollar una función que reciba un usuario y un password, los compare y determine si el acceso es correcto, retornando `true` o `false`.

    ```go
    func verificarAcceso(usuario, password, usuarioCorrecto, passwordCorrecto string) bool {
        if usuario == usuarioCorrecto && password == passwordCorrecto {
            return true
        }
        return false
    }

    // Ejemplo de uso
    usuario := "admin"
    password := "1234"
    usuarioCorrecto := "admin"
    passwordCorrecto := "1234"

    if verificarAcceso(usuario, password, usuarioCorrecto, passwordCorrecto) {
        fmt.Println("Acceso concedido")
    } else {
        fmt.Println("Acceso denegado")
    }
    ```
