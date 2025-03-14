# Instalar Go en Windows

💡 En las versiones más recientes de Go, el proceso de instalación ha cambiado. Aquí encontrarás la forma recomendada de instalar la versión más actualizada del lenguaje para desarrollar los ejercicios del curso sin problemas.

Para descargar Go, visita [la página oficial](https://go.dev/dl/) y selecciona la opción para Microsoft Windows.

Al ejecutar el instalador, sigue el proceso típico de instalación de aplicaciones en Windows.

Al finalizar la instalación, se creará automáticamente la variable de entorno `GOPATH` apuntando a `%USERPROFILE%\go` y la carpeta `pkg`. También se agregará el binario ejecutable de Go a la variable de entorno `Path`.

## Verificar la instalacion

Para verificar que la instalación fue correcta, abre una terminal en Windows PowerShell y ejecuta el siguiente comando:

```sh
go version
```

Deberías ver un resultado similar a:

``` sh
go version go1.17.6 windows/amd64
```

Si hay algún error (como que el comando no se reconozca), reinicia el sistema y verifica que las variables de entorno se hayan creado correctamente. Sigue estos pasos:

1. En el buscador de Windows, escribe: _edit the system environment variables_ y selecciona el programa que aparece.
2. Selecciona la opción _Environment Variables..._
3. Valida que `GOPATH` apunta a `%USERPROFILE%\go`
4. Valida que en `Path` esté incluido el ejecutable de Go, por ejemplo: `C:\Program Files\Go\bin`

### Ejecutar código Go

Desde la implementación de los Go Modules (a partir de la versión 1.12), puedes ejecutar código Go fuera del `GOPATH`. Sin embargo, es una buena práctica mantener la estructura anterior. En tu `GOPATH` (`%USERPROFILE%\go`), crea la carpeta `src` para tus proyectos de Go.

Dentro de esa carpeta, abre Visual Studio Code (o tu editor favorito) y crea el archivo `helloWorld.go` con el siguiente código:

```go
package main

import (
    "fmt"
)

func main() {
    fmt.Println("Hello World")
}
```

Ejecuta este código desde la consola con:

```sh
go run helloWorld.go
```

Deberías ver el siguiente mensaje en la consola:

``` sh
Hello World
```

Para compilarlo, ejecuta:

```sh
go build helloWorld.go
```

Esto generará un archivo llamado `helloWorld.exe`. Para ejecutarlo, usa:

```sh
helloWorld.exe
```

El resultado en la consola debería ser el mismo.

### Recomendaciones finales

Para el desarrollo de Go en Windows, se recomienda usar WSL (Windows Subsystem for Linux), ya que Go tiene un mejor desempeño en Linux. Sin embargo, con cada nueva versión, el rendimiento en Windows mejora.
