# Llave valor con Maps

En Go, los Maps son estructuras de datos que permiten almacenar pares de llaves y valores. A diferencia de los Arrays, los Maps no tienen un tamaño fijo y pueden crecer dinámicamente.

## ¿Qué son los maps en Go y cómo funcionan?

Los maps en Go son estructuras de datos clave-valor que nos permiten acceder a un valor específico utilizando una llave única. Imagina un estacionamiento en el que cada puesto tiene un número asignado; para acceder a un vehículo, necesitas saber la clave asociada, que es el número de estacionamiento. En otros lenguajes como Python, estas estructuras se conocen como diccionarios. Conozcamos más sobre cómo funcionan los maps en Go.

## ¿Cómo se crea un map en Go?

Crear un map en Go es un proceso sencillo gracias al uso del keyword `make`. Comenzamos declarando nuestro map, que puede almacenar las edades de varias personas. Aquí te mostramos cómo hacerlo:

```go
m := make(map[string]int)
```

En este fragmento, `map[string]int` define un map cuya clave es de tipo `string` y el valor es de tipo `int`. A continuación, añadimos algunos valores:

```go
m["José"] = 14
m["Pepito"] = 20
```

## ¿Cómo imprimimos un map?

Para visualizar el contenido del map, simplemente lo imprimimos en consola. Al hacerlo, notamos que los elementos se separan por espacios en lugar de comas, algo característico de los maps en Go:

```go
fmt.Println(m) // Salida: map[José:14 Pepito:20]
```

## ¿Cómo recorremos un map?

Para recorrer un map en Go, utilizamos un bucle `for` combinado con `range`, que nos permite iterar sobre cada clave y valor en el map:

```go
for key, value := range m {
    fmt.Println(key, value)
}
```

Es importante destacar que el orden de los elementos en un map es aleatorio debido a su naturaleza concurrente. Si necesitas mantener el orden original, considera usar slices junto con los conocimientos previos de 'slicing'.

## ¿Cómo accedemos a un valor específico?

Si queremos obtener un valor específico, simplemente accedemos a través de la clave correspondiente, como en el siguiente ejemplo:

```go
value := m["José"]
fmt.Println(value) // Salida: 14
```

## ¿Qué sucede si accedemos a una clave inexistente?

Al acceder a una clave que no existe, Go retorna el valor cero correspondiente al tipo de dato del map. Por ejemplo, si buscamos una clave no almacenada como en el caso de "Joseph", obtendremos un cero:

```go
value := m["Joseph"]
fmt.Println(value) // Salida: 0
```

Para abordar esta situación, se puede utilizar una segunda variable llamada `ok`, que indica si la clave existe en el map:

```go
value, ok := m["Joseph"]
fmt.Println(value, ok) // Salida: 0 false
```

Si la clave existe, `ok` será `true`; en caso contrario, será `false`. Esta propiedad resulta esencial para confirmar la presencia de un elemento dentro del map.

## ¿Por qué elegir maps en Go?

Los maps son estructuras eficientes en Go, especialmente cuando trabajamos con listas relacionadas. Su implementación nativa de concurrencia permite realizar operaciones de manera más fluida y eficaz que las alternativas como slices o arrays.
