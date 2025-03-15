# Go modules: Ir más allá del GoPath con Echo

## ¿Qué son los problemas comunes al programar en Go y cómo GoModules los soluciona?

Durante mucho tiempo, los desarrolladores de Go enfrentaron dos grandes desafíos al programar. En primer lugar, había que dividir el código dentro de la ruta del GoPath: el home, el nombre de usuario y la carpeta Go. En segundo lugar, al llevar el código a producción, especialmente cuando se requerían librerías de terceros, muchas plataformas no permitían insertar el código completo a nivel de carpeta, sino solo el binario o código de ejecución. Esto limitaba el uso de frameworks externos para la programación web en Go.

La solución llegó con la introducción de GoModules alrededor de 2018-2019. GoModules cambió las reglas del juego al simplificar la instalación de librerías de terceros y resolver los problemas iniciales de los desarrolladores de Go.

## ¿Cómo instalar el framework Echo para desarrollo web en Go?

El framework Echo es una excelente herramienta para quienes deseen desarrollar aplicaciones web con Go. A continuación, te explico cómo puedes comenzar a trabajar con Echo:

## Acceso al framework Echo

1. Realiza una búsqueda de "Echo Golang" en Google.
2. Accede a la página del framework y revisa la documentación en GitHub o la guía de usuario disponible en GetStarted.

### Instalación de Echo

Desde la terminal, utiliza el comando `go install` para instalar el paquete necesario, asegurando siempre trabajar con la versión más estable del repositorio:

```sh
go install -u github.com/labstack/echo/v4
```

### ¿Cómo dar inicio a un módulo con GoModules y crear un Hello World?

Trabajar con GoModules para iniciar un módulo es sencillo y viene con la ventaja de gestionar las dependencias de forma eficiente.

#### Iniciar un módulo

En el terminal, dentro de la carpeta de tu proyecto, ejecuta:
```sh
go mod init github.com/tu-usuario/nombre-del-repositorio
```
Esto genera el archivo `go.mod`, que lista las dependencias del proyecto.

#### Crear un servidor básico con Echo

1. Instanciar Echo:

    ```go
    e := echo.New()
    ```

2. Configurar una ruta base y definir una función para responder "Hello World":

    ```go
    e.GET("/", func(c echo.Context) error {
        return c.String(http.StatusOK, "Hello World")
    })
    ```

3. Iniciar el servidor en el puerto 1323:

    ```go
    e.Logger.Fatal(e.Start(":1323"))
    ```

4. Ejecutar el código con `go run` y verificar que el servidor responde en `localhost:1323`.

## ¿Cómo modificar un framework o librerías en Go e instalar GoModules?

Modificar librerías en Go no es comúnmente recomendado, pero puede ser útil para fines educativos o de debugging. Aquí los pasos básicos:

### Localizar la librería

Accede al GoPath y busca la carpeta `pkg`, donde se almacenan los módulos descargados.

### Editar archivos

Localiza el archivo que deseas modificar, realiza un respaldo y edita el archivo necesario.

### Comprobar verificación del módulo

Después de modificar, utiliza `go mod verify` para asegurarte de que las modificaciones no han dañado el módulo:

```sh
go mod verify
```

Con estos pasos, ahora puedes disfrutar de la flexibilidad y eficacia que ofrecen los GoModules mientras aprendes y experimentas con el framework Echo.