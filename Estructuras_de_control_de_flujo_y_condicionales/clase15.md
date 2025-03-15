# Múltiple condiciones anidadas con Switch

## ¿Cómo usar el switch en Go?

Go es un lenguaje de programación que ofrece diversas herramientas para simplificar la escritura de código, y una de estas herramientas es el uso de switch. Aunque en Go podemos utilizar la sentencia if para evaluar condiciones, usar switch puede hacer que tu código sea mucho más limpio y fácil de mantener. Pero, ¿cuándo exactamente deberías usar switch en lugar de if? Vamos a explorarlo mediante ejemplos prácticos de código en Go y algunos consejos sobre cuándo usar cada uno.

### ¿Por qué elegir switch sobre if en Go?

- **Claridad y legibilidad**: Cuando tienes múltiples condiciones, if anidados pueden hacer que el código sea difícil de seguir. Switch simplifica estas situaciones al proporcionar una estructura más clara.
- **Facilidad de mantenimiento**: La legibilidad del código también ayuda a mantenerlo en el tiempo. Con switch, agregar, eliminar o modificar casos es generalmente más sencillo.

### ¿Cómo usar switch con una condición especificada?

Aquí hay un ejemplo de cómo podrías determinar si un número es par o impar usando switch en Go:

```go
package main

import (
    "fmt"
)

func main() {
    modulo := 5 % 2

    switch modulo {
    case 0:
        fmt.Println("El número es par.")
    default:
        fmt.Println("El número es impar.")
    }
}
```

**Explicación del código**: Este ejemplo utiliza una variable `modulo` para almacenar el resultado de `5 % 2`. El operador `%` devuelve el sobrante de una división entre números. Si el resultado es 0, el número es par; de lo contrario, es impar.

**Uso del default**: El caso `default` se ejecuta si ninguno de los otros casos se cumple, un buen recurso cuando las condiciones posibles son limitadas.

### ¿Cómo mejorar el switch con inicialización de variables?

También puedes encontrar el switch escrito de esta manera reducida y elegante:

```go
package main

import (
    "fmt"
)

func main() {
    switch modulo := 4 % 2; modulo {
    case 0:
        fmt.Println("El número es par.")
    default:
        fmt.Println("El número es impar.")
    }
}
```

**Variable y definición**: Definimos la variable `modulo` junto con el switch. Después del punto y coma, se evaluará `modulo`.

### ¿Cuándo usar un switch sin ninguna condición?

Un switch también puede funcionar sin una condición explícita, lo que es útil para comparar una misma variable con múltiples condiciones similares a if:

```go
package main

import (
    "fmt"
)

func main() {
    value := 200

    switch {
    case value > 100:
        fmt.Println("Es mayor a cien.")
    case value < 0:
        fmt.Println("Es menor a cero.")
    default:
        fmt.Println("No condición.")
    }
}
```

**Método de evaluación**: Aquí, evaluamos directamente dentro de cada `case`, sin necesidad de una variable en el switch.

**Mayor flexibilidad**: Ideal para múltiples condiciones sobre una misma variable.

### ¿Cuándo debo usar cada tipo de switch?

- **Con condición especificada**: Úsalo cuando sabes de antemano qué valores tendrá la variable y necesitas verificar esos valores específicos.
- **Sin condición especificada**: Es útil para decisiones más complejas, donde las condiciones se parecen más a las usadas en un if.
