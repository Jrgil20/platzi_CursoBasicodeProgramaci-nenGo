# Variables, constantes y zero values

## ¿Cómo declarar constantes en Go?

Las constantes son elementos fundamentales en cualquier lenguaje de programación ya que permiten definir valores que no cambian a lo largo del tiempo.

### Declaración básica con tipo explícito:

```go
const pi float64 = 3.14
```

Aquí se define una constante `pi` de tipo `float64`.

### Sin especificar el tipo (inferencia de tipo):

```go
const pi2 = 3.14
```

Ambos métodos son válidos y útiles dependiendo de la situación. La diferencia está en si el compilador infiere o necesitas especificar el tipo de la constante.

## ¿Qué son las variables y cómo se declaran en Go?

Las variables permiten almacenar datos que pueden cambiar durante la ejecución del programa. En Go, tienes varias formas de declararlas.

### Uso de la inferencia de tipo:

```go
base := 12
```

### Especificación del tipo al momento de declarar:

```go
var altura int = 14
```

### Declaración de la variable sin asignación inicial:

```go
var area int
```

### Algunos puntos clave a tener en cuenta:

- Las variables deben ser usadas en tu código. Si declaras una variable pero no la utilizas, Go te generará un error.
- El compilador de Go asigna automáticamente un valor por defecto (zero value) a las variables que no tienen un valor inicial específico:
    - `int`: 0
    - `float64`: 0.0
    - `string`: ""
    - `bool`: false

## ¿Cómo realizar operaciones utilizando constantes y variables?

Como ejercicio, podemos calcular el área de un cuadrado utilizando constantes y variables. Esto es útil para automatizar y simplificar tareas repetitivas.

### Define las dimensiones del cuadrado como constantes:

```go
const baseCuadrado = 10
```

### Calcula el área al mismo tiempo que Go infiere el tipo de `areaCuadrado`:

```go
areaCuadrado := baseCuadrado * baseCuadrado
```

### Imprime el resultado:

```go
fmt.Println("Área del cuadrado:", areaCuadrado)
```

Este ejemplo muestra lo sencillo y claro que puede ser el manejo de variables y constantes en Go. El lenguaje ofrece una estructura sólida que mezcla eficiencia y simplicidad, propiciando el desarrollo de código claro y eficiente.

Recuerda, dominar el uso de constantes y variables en Go no solo facilitará tus tareas diarias, sino que también te ayudará a prevenir errores comunes relacionados con la reasignación indebida de valores. ¡Sigue explorando y aprendiendo! En la próxima clase, aprenderás sobre operadores aritméticos. ¡Te esperamos!