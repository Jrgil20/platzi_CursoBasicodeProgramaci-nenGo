# Operadores de Comparación en Go

En Go, los operadores de comparación se utilizan para comparar dos valores. El resultado de una comparación es un valor booleano (`true` o `false`). A continuación se presentan los operadores de comparación disponibles en Go:

| Operador | Descripción                  | Ejemplo       |
|----------|------------------------------|---------------|
| `==`     | Igual a                      | `a == b`      |
| `!=`     | No igual a                   | `a != b`      |
| `<`      | Menor que                    | `a < b`       |
| `<=`     | Menor o igual que            | `a <= b`      |
| `>`      | Mayor que                    | `a > b`       |
| `>=`     | Mayor o igual que            | `a >= b`      |

## Ejemplos de Uso

A continuación se muestran algunos ejemplos de cómo utilizar los operadores de comparación en Go:

```go
package main

import "fmt"

func main() {
    a := 10
    b := 20

    fmt.Println("a == b:", a == b) // false
    fmt.Println("a != b:", a != b) // true
    fmt.Println("a < b:", a < b)   // true
    fmt.Println("a <= b:", a <= b) // true
    fmt.Println("a > b:", a > b)   // false
    fmt.Println("a >= b:", a >= b) // false
}
```

En este ejemplo, se comparan los valores de `a` y `b` utilizando los diferentes operadores de comparación y se imprimen los resultados.

## Notas

- Los operadores de comparación se pueden utilizar con tipos de datos numéricos, cadenas de texto y otros tipos compatibles.
- Es importante recordar que el resultado de una comparación siempre será un valor booleano (`true` o `false`).
