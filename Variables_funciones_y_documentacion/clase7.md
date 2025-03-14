# Tipos de datos primitivos

## ¿Qué son los valores primitivos en Go?

Los valores primitivos, o tipos de datos primitivos, son las unidades fundamentales en la programación, y en Go no son la excepción. Estos tipos de datos ayudan a especificar la clase de datos que se manejará dentro del código, mejorando su rendimiento y eficiencia. Al definir el tipo de dato, puedes optimizar la memoria utilizada y maximizar el potencial de tu código, alineándolo con la arquitectura del sistema operativo en uso.

## ¿Cómo se manejan los números enteros en Go?

### ¿Cómo se definen los números enteros Sen Go?

En Go, los números enteros se declaran utilizando el tipo `int`. Si no se especifica un tamaño específico como `int8`, `int16`, `int32` o `int64`, el lenguaje seleccionará automáticamente el tamaño en función de la arquitectura del sistema operativo:

- 32 bits: `int32`
- 64 bits: `int64`

Cada tipo de `int` tiene un rango específico de valores que puede almacenar, determinado por su tamaño en bits:

- `int8`: Almacena números de -128 a 127.
- `int16`: Almacena desde -2^15 a 2^15-1.
- `int32`: Almacena desde -2^31 a 2^31-1.
- `int64`: Almacena desde -2^63 a 2^63-1.

### ¿Qué pasa con los números enteros positivos?

Los números enteros positivos o `uint` optimizan aún más el almacenamiento y aumentan la cantidad máxima de valores positivos que puede almacenarse:

- `uint8`: De 0 a 2^8-1.
- `uint16`: De 0 a 2^16-1.
- `uint32`: De 0 a 2^32-1.
- `uint64`: De 0 a 2^64-1.

El uso de `uint` es ideal cuando se está seguro de que los valores no serán negativos, maximizando así la eficiencia en el manejo de memoria.

## ¿Cómo se manejan los números decimales, textos, booleanos y números complejos?

### ¿Qué tipos de datos decimales existen en Go?

Go admite dos tipos de datos para números decimales:

- `float32`: Almacena números dentro del rango de ±1.8 x 10^-38 a ±3.4 x 10^38.
- `float64`: Almacena números dentro del rango de ±2.23 x 10^-308 a ±1.8 x 10^308.

Los números decimales son esenciales cuando se necesita precisión en los cálculos matemáticos y científicos.

### ¿Cómo se gestionan los textos en Go?

Para manejar textos, Go utiliza el tipo de dato `string`. Es crucial distinguir que, a diferencia de otros lenguajes de programación que permiten comillas simples o dobles, en Go, los `strings` deben delimitarse utilizando comillas dobles.

### ¿Cómo se usan los valores booleanos?

El tipo de dato booleano en Go, `bool`, admite únicamente dos valores: `true` o `false`. Estos son fundamentales para controlar estructuras condicionales y bucles dentro del código.

### ¿Qué son los números complejos en Go?

Go ofrece una capacidad avanzada al incluir soporte para números complejos, los cuales se representan mediante `complex64` y `complex128`:

- `complex64`: El componente real e imaginario son de tipo `float32`.
- `complex128`: Ambos componentes son de tipo `float64`.

En números complejos, la parte imaginaria incorpora el valor `i`. Por ejemplo, para definir un número complejo en Go, podría usarse la notación `10 + 8i`, donde 10 es la parte real y 8i es la imaginaria.

Definir correctamente los tipos de datos en Go no solo mejora el desempeño del software, sino que también fomenta buenas prácticas de programación. Por ello, aunque el código podría funcionar sin especificaciones explícitas de tipos, hacerlo concede un enfoque más profesional y sólido.
