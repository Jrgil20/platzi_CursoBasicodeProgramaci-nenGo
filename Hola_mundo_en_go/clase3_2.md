# Instalar Go en Mac

El proceso de instalación de Go en Mac es muy similar al proceso de instalación en Linux. Sin embargo, en Mac, el paquete Go suele estar más actualizado respecto a la última versión estable.

## Instalación

### Método 1

Usaremos el manejador de paquetes `brew`, muy popular para la programación en macOS.

Para instalar Go, ingresa en la terminal el siguiente comando:

```sh
brew install golang
```

### Método 2

Visita la [página de descarga de Go](https://golang.org/dl/), y selecciona macOS.

Una vez descargado el paquete, ábrelo y sigue los pasos del asistente para instalar Go en tu Mac. Por defecto, se instalará en la ruta `/usr/local/go`.

Si todo ha salido bien, al abrir una terminal y ejecutar `go version`, deberías ver la versión de Go que acabas de instalar.

## Pasos Finales

Una vez tengas instalado Go, el siguiente paso es crear la carpeta `go`, preferiblemente en tu `$HOME`.

Dentro de esta carpeta, crea las siguientes tres subcarpetas: `pkg`, `src` y `bin`.

Finalmente, crea las variables de entorno. Esto dependerá del shell que estés utilizando. En el curso usaremos Bash Shell, así que abre el archivo `~/.bash_profile` y añade lo siguiente al final del archivo:

```sh
export GOPATH=$HOME/go
export GOBIN=$GOPATH/bin
export PATH=$PATH:$GOBIN
```

El paso final es instalar tu editor de código favorito. En el curso usaremos VS Code.
