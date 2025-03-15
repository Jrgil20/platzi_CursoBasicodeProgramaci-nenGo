# Operadores Lógicos en Go

En Go, los operadores lógicos se utilizan para combinar valores booleanos y devolver un resultado booleano. Los operadores lógicos más comunes son `&&` (AND), `||` (OR) y `!` (NOT).

## Operador AND (`&&`)

El operador AND devuelve `true` solo si ambos operandos son `true`.

```go
a := true
b := false
result := a && b // result es false
```

## Operador OR (`||`)

El operador OR devuelve `true` si al menos uno de los operandos es `true`.

```go
a := true
b := false
result := a || b // result es true
```

## Operador NOT (`!`)

El operador NOT invierte el valor de un operando booleano.

```go
a := true
result := !a // result es false
```

## Ejemplo Completo

```go
package main

import "fmt"

func main() {
    a := true
    b := false

    fmt.Println("a && b:", a && b) // false
    fmt.Println("a || b:", a || b) // true
    fmt.Println("!a:", !a)         // false
    fmt.Println("!b:", !b)         // true
}
```

Este ejemplo muestra cómo utilizar los operadores lógicos en Go para evaluar expresiones booleanas.
