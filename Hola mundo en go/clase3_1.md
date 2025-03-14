# Instalar Go en Linux
## ¿Cómo instalar Go en Linux?

Instalar Go en Linux puede parecer complicado, pero es esencial para desarrollar aplicaciones eficientemente en esta plataforma. Dado que Go suele tener un mejor rendimiento en Linux, se recomienda usar este sistema operativo. Aquí tienes una guía paso a paso.

### ¿Dónde descargar Go?

1. Abre tu navegador y busca "golang download" en Google.
2. Accede a la primera URL que aparezca. Esto te llevará a la página de descargas de Go.
3. Selecciona el instalador para Linux. Se descargará un archivo comprimido.

### ¿Cómo instalar Go en tu sistema Linux?

Una vez descargado el archivo, sigue estos pasos:

1. **Abrir la Terminal**: Navega a la ubicación del archivo descargado.
2. **Descomprimir el archivo**: Usa el comando `tar -C /usr/local -xzf go1.x.x.linux-amd64.tar.gz` (reemplaza `go1.x.x` con la versión descargada).
3. **Exportar Variables de Entorno**: Añade las siguientes líneas al final de tu archivo `.bashrc` o `.zshrc`:
    ```sh
    export GOPATH=$HOME/go
    export GOBIN=$GOPATH/bin
    export GOROOT=/usr/local/go
    export PATH=$PATH:$GOBIN
    export PATH=$PATH:$GOROOT/bin
    ```
4. **Actualizar la configuración del Shell**: Ejecuta `source ~/.bashrc` o `source ~/.zshrc` para aplicar los cambios.

### ¿Cómo crear el entorno de trabajo para Go?

Organizar tu entorno de trabajo es crucial:

1. **Crear la carpeta go**:
    ```sh
    mkdir -p $HOME/go
    cd $HOME/go
    ```
2. **Crear subcarpetas bin, pkg, src**:
    ```sh
    mkdir -p $HOME/go/{bin,pkg,src}
    ```

### ¿Cómo verificar la instalación de Go?

Para asegurarte de que Go se instaló correctamente, ejecuta:
```sh
go version
```
Deberías ver la versión de Go instalada. ¡Ahora estás listo para desarrollar con Go en Linux!

Recuerda que si encuentras problemas, hay comunidades activas donde puedes buscar ayuda y aprender de otros usuarios. ¡La programación es un camino lleno de aprendizajes y crecimiento!

