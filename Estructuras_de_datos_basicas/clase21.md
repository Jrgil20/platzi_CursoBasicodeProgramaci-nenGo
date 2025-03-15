# Modificadores de acceso en funciones y Struct

En Go, los modos de acceso a los campos de un struct y a las funciones son similares a los de otros lenguajes de programación. Sin embargo, Go no tiene palabras clave como `public`, `private` o `protected` para definir la visibilidad de los campos y métodos. En su lugar, Go utiliza convenciones para definir la visibilidad de los campos y métodos.

## Cómo manejar el acceso público y privado en Go

En el mundo de la programación, la gestión de acceso de las variables y funciones es fundamental para asegurar la integridad y seguridad del código. En lenguajes de programación como Go, esto se maneja de manera diferente a otros lenguajes como Java o C#. Aunque no existen palabras clave específicas para definir el acceso, con un simple cambio en la capitalización se controla la visibilidad de los elementos.

### ¿Cómo se define el acceso a variables en Go?

En Go, el acceso se controla con la capitalización del nombre.

- **Público**: Si el nombre empieza con letra mayúscula. Esto significa que es accesible desde otros paquetes.
- **Privado**: Si el nombre empieza con letra minúscula. Esto limita el acceso al paquete donde se ha declarado.

Estas reglas se aplican a las estructuras de datos, funciones y cualquier otro tipo de dato. Veamos cómo se implementa en el código.

```go
type Car struct {
    Brand string
    Year  int
}

type privateCar struct {
    brand string
    year  int
}
```

En este ejemplo, `Car` es de acceso público mientras que `privateCar` es privado.

### ¿Cómo se modifica un paquete en Go?

Para trabajar con paquetes en Go, podemos añadir archivos y mantener nombres consistentes tanto para el paquete como para la carpeta y el archivo principal. Siempre es buena práctica documentar cada elemento público para mantener claridad en el código.

#### Crear un nuevo archivo de paquete

```sh
mkdir mypackage
cd mypackage
touch mypackage.go
```

#### Definir el paquete y el tipo dentro del archivo

```go
package mypackage

// CarPublic es de acceso público
type CarPublic struct {
    Brand string
    Year  int
}

type carPrivate struct {
    brand string
    year  int
}
```

Como se observa, `CarPublic` está comentado, lo que es una práctica estandarizada para tipos públicos. Es importante notar cómo definimos las estructuras de datos dentro del paquete, utilizando mayúsculas para los públicos.

### ¿Cómo importar y usar un paquete?

La importación de paquetes en Go permite usar las funciones y tipos públicos de otros paquetes. Es frecuente ver cómo el editor facilita esta tarea mediante autocompletado. En caso contrario, se puede hacer manualmente.

#### Ejemplo de uso del paquete

```go
package main

import (
    "mypackage"
    "fmt"
)

func main() {
    myCar := mypackage.CarPublic{
        Brand: "Ferrari",
        Year:  2020,
    }
    fmt.Println(myCar)
}
```

En este código, creamos una instancia del tipo `CarPublic` y accedemos a sus campos. Observa cómo se usa la importación desde el directorio `src`.

### ¿Qué son los alias y cómo se utilizan?

Los alias son una forma de renombrar un paquete cuando se importa, útil para evitar conflictos o simplificar el nombre. Esto se hace anteponiendo el alias al paquete durante la importación.

#### Uso de alias

```go
import (
    pk "mypackage"
)

func main() {
    myCar := pk.CarPublic{
        Brand: "Tesla",
        Year:  2021,
    }
    fmt.Println(myCar)
}
```

Aquí, `pk` sirve como alias para `mypackage`, simplificando la referencia al mismo.

### ¿Cuál es la importancia de las funciones públicas y privadas?

Al igual que los tipos, las funciones también pueden ser públicas o privadas. El uso de funciones privadas ayuda a ocultar detalles de implementación, mejorando la encapsulación y modularidad.

#### Definición y uso de funciones

```go
package mypackage

import "fmt"

// PrintMessage es una función pública que imprime un mensaje
func PrintMessage(text string) {
    fmt.Println(text)
}

func printSecretMessage(secret string) {
    fmt.Println("Secret:", secret)
}
```

En el paquete principal, sólo `PrintMessage` estará accesible externamente.

```go
func main() {
    mypackage.PrintMessage("Hola, Flash!")
}
```

Con estas implementaciones, tu código estará más seguro y organizado. Recuerda, dominar Go incrementa tus capacidades como desarrollador.
