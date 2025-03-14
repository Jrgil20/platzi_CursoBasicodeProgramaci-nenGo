# ¿Cómo crear tu primer programa en Go?

Crear un "Hola Mundo" es un rito que todos los programadores deben atravesar. Este sencillo pero poderoso paso te introduce al lenguaje, sus herramientas y prácticas recomendadas. En este artículo, aprenderás a escribir tu primer programa en el lenguaje Go, utilizando herramientas como Visual Studio Code y su terminal integrada.

## ¿Cómo prepararte para escribir código en Go?

Primeramente, es esencial tener tu entorno de trabajo listo. Aquí tienes una guía básica para comenzar:

### Acceso al directorio gopad

Ingresa a la terminal y navega a tu carpeta de trabajo usando `cd`, por ejemplo:

```sh
cd gopad
```

### Directorios y archivos necesarios

1. Ve a tu carpeta `src` donde crearás tu código:

    ```sh
    cd src
    ```

2. Crea una nueva carpeta para tu proyecto:

    ```sh
    mkdir curso-golang-flaxy
    ```

3. Ingresa a tu nueva carpeta:

    ```sh
    cd curso-golang-flaxy
    ```

### Configura Visual Studio Code

1. Abre el editor desde tu terminal con el comando:

    ```sh
    code .
    ```

2. Si Visual Studio Code sugiere instalar una extensión para Go, acéptalo. Esto mejorará tu experiencia al programar.

## ¿Cómo escribir tu primer "Hola Mundo" en Go?

Dentro de Visual Studio Code, sigue estos pasos para escribir y ejecutar tu primer programa:

### Crea la estructura del programa

En tu carpeta `src`, crea un archivo nuevo llamado `main.go`:

```sh
touch src/main.go
```

### Escribe el código

Abre `main.go` y escribe el siguiente código:

```go
package main

import "fmt"

func main() {
    fmt.Println("Hola Mundo")
}
```

- `package main`: Declara el paquete principal del programa.
- `import "fmt"`: Importa el paquete estándar de Go para impresión en consola.
- `func main() { ... }`: Define la función principal donde se ejecuta el código.

## ¿Cómo ejecutar tu programa?

Ahora que tienes tu código listo, tienes dos formas de ejecutarlo:

### Compilación y ejecución tradicional con `go build`

Compila el programa para obtener un archivo ejecutable:

```sh
go build src/main.go
./main
```

- **Ventaja**: Eficiente para producción.
- **Desventaja**: Cada cambio requiere recompilar.

### Ejecución directa con `go run`

Ejecuta tu programa de inmediato sin guardar el ejecutable:

```sh
go run src/main.go
```

- **Ventaja**: Ideal para desarrollo y pruebas rápidas.
- **Desventaja**: No guarda el archivo ejecutable.

## ¿Cómo Go te obliga a seguir buenas prácticas?

Go promueve una codificación coherente y legible. La herramienta de formato te ayuda a mantener el código limpio automáticamente:

- **Formatos de bloque de código**: La extensión de Go en Visual Studio Code organizará tus llaves al guardar.
- **Espacios en blanco**: Cualquier espacio extra será eliminado al guardar.

La herramienta formatea automáticamente el código y elimina importaciones redundantes, ayudando a mantener un código limpio y eficiente.

## ¿Qué sigue después de "Hola Mundo"?

El camino a dominar Go apenas comienza. En la próxima lección, exploraremos cómo crear variables, constantes y zero values. Estos conceptos te darán una base sólida para continuar tu aprendizaje y profundizar en el lenguaje.

¡Adelante, sumérgete en el emocionante mundo de Go! Con práctica constante y curiosidad, pronto dominarás este poderoso lenguaje.
