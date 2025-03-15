# Go get: El manejador de paquetes

## ¿Qué es un manejador de paquetes en programación?

Los lenguajes de programación, como Python o Go, a menudo incluyen un manejador de paquetes, una herramienta esencial que permite a los desarrolladores compartir y gestionar librerías de código. Por ejemplo, Python utiliza pip y Go utiliza `go get`. Estos manejadores facilitan la reutilización de código y el acceso a librerías desarrolladas por la comunidad.

## Nota importante

Es importante mencionar que `go get` ha sido deprecado y ya no se utiliza para instalar paquetes. En su lugar, se debe usar `go install`.

## ¿Cómo se utiliza `go install` en Go?

Para utilizar el manejador de paquetes en Go, simplemente necesitas ejecutar `go install`. No es necesario instalar nada adicional, ya que Go ya lo incluye al momento de su instalación. Con `go install`, puedes descargar librerías como GoTour, que fue desarrollada por Google para facilitar la introducción a Go de manera online.

## ¿Cómo instalar librerías con `go install`?

Ejecuta el comando:

``` sh
go install golang.org/x/tour@latest
```

Para obtener feedback durante la descarga, añade la bandera `-v` para ver los detalles. Usa la bandera `-u` para asegurarte de que la librería se descargue nuevamente, aunque ya esté instalada.

## ¿Dónde se instala el código descargado?

Una vez descargado, el binario instalable se encuentra en la carpeta `bin` de tu `GOPATH`. El código fuente queda en la carpeta `src`, organizada de acuerdo a la estructura de directorios.

## Descubriendo la librería GoTour

GoTour ofrece una plataforma offline para aprender Go. Es como un mini-curso, permitiéndote:

- Activar o desactivar opciones de sintaxis.
- Ejecutar código y restablecerlo a su estado original.
- Cambiar el idioma, aunque se recomienda usar inglés por la compleción de los recursos.

Se puede explorar distintos módulos sobre conceptos básicos de Go, como:

- Declaración de variables
- Estructuras
- Interfaces y concurrencia

## ¿Dónde encontrar paquetes externos interesantes?

Existen varias fuentes, pero una recomendada es [awesome-go.com](https://awesome-go.com), que agrupa numerosos proyectos y librerías por categoría, como:

- Audio y música.
- CSS.
- Bases de datos.
- Machine learning, entre otros.

Adicionalmente, librerías como Echo son muy valoradas para el desarrollo web con Go. Estos recursos incluyen enlaces a sus repositorios en GitHub, proporcionando un acceso fácil al código y documentación.

## Reto práctico

Se invita a seguir explorando [awesome-go.com](https://awesome-go.com), elegir una librería de interés y crear un "Hello World" con ella. Esto no solo te ayudará a practicar lo aprendido, sino también a familiarizarte más con el ecosistema de Go. Comparte tus avances y descubrimientos en la sección de comentarios, ¡se aprecia mucho el entusiasmo por compartir conocimiento! Te esperamos en la próxima clase para descubrir el uso de Go Mods con Echo.
