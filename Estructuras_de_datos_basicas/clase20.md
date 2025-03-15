# Structs: La forma de hacer clases en Go

## ¿Qué es un struct en Go?

La programación orientada a objetos (OOP) es una práctica común en muchos lenguajes de programación. Uno de los pilares de la OOP son las clases, las cuales están compuestas por atributos y métodos que permiten la interacción con los objetos. Sin embargo, Go, con su simplicidad y eficiencia, no utiliza clases, sino una estructura alternativa llamada struct. Esta herramienta poderosa y versátil permite la creación de estructuras de datos más simples pero igualmente efectivas. En este contenido, exploraremos cómo funcionan los structs y cómo pueden ser útiles para tus proyectos en Go.

## ¿Cómo crear un struct en Go?

Para crear un struct en Go, se utiliza la palabra clave `type` seguida del nombre que deseas dar al struct, y a continuación la palabra clave `struct`. Vamos a ver esto con un ejemplo paso a paso.

```go
type Car struct {
    Brand string
    Model int
}
```

Aquí, hemos creado una estructura `Car` con dos atributos: `Brand` de tipo `string` y `Model` de tipo `int`. Estos atributos se asemejan a las propiedades que definirías en una clase dentro de un lenguaje orientado a objetos tradicional.

## ¿Cómo instanciar un struct?

Una vez definido un struct, hay varias formas de instanciarlo. Comencemos con la forma más directa:

### Instanciación directa

```go
myCar := Car{
    Brand: "Ford",
    Model: 2020,
}
```

En este ejemplo, instanciamos un objeto `Car` asignando valores directos a los atributos `Brand` y `Model`. Este método se asemeja a crear un objeto y asignar valores en lenguajes como Java o Python.

### Instanciación por declaración

Otra forma de crear una instancia de un struct es mediante la declaración inicial sin asignar valores:

```go
var anotherCar Car
anotherCar.Brand = "Ferrari"
```

Aquí se crea `anotherCar` sin valores iniciales. Posteriormente, se asigna el valor al atributo `Brand`. Si se imprime este struct, el atributo `Model` mostrará un valor por defecto de cero, debido a que Go asigna cero en campos numéricos no inicializados, tal como lo explora el tema de "zero values".

## ¿Cuál es la utilidad de los structs en Go?

Los structs no solo ofrecen una forma eficiente de agrupar datos relacionados, sino que también permiten:

- **Extensibilidad**: Puedes agregar nuevas propiedades sin cambiar la lógica básica del programa.
- **Métodos asociados**: En futuras lecciones también aprenderás cómo agregar métodos a los structs, ampliando su funcionalidad y permitiendo el encapsulamiento de comportamiento junto a los datos.

Puedes incluso personalizar tus structs para que reflejen con más precisión las características que te interesan, como agregar el color del carro o el tipo de motor.

## ¿Cómo podrías seguir desarrollando un struct?

El potencial de un struct crece a medida que avanzas en tus habilidades con Go. Un ejercicio útil es imaginar cómo mejorar tus structs con atributos adicionales:

- **Color**: Añadir un atributo que indique el color del carro.
- **EngineType**: Saber si es eléctrico, híbrido o de combustión interna.
- **Owner**: Incluir información sobre el propietario del carro.

Por ejemplo:

```go
type Car struct {
    Brand      string
    Model      int
    Color      string
    EngineType string
    Owner      string
}
```

Esta ampliación refleja cómo los structs pueden adaptarse a las necesidades cambiantes y específicas de tu aplicación. ¡No dudes en experimentar y explorar todas las capacidades que Go ofrece con los structs!
