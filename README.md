# Sistema de Gestión de Streaming en Go 🎬

Este proyecto es un sistema básico para gestionar y reproducir videos en streaming, desarrollado en Go.

## Estructura del proyecto

- `main.go`: servidor principal que configura las rutas HTTP.
- `models/video.go`: definición de la estructura `Video`.
- `handlers/videoHandler.go`: funciones que manejan las peticiones HTTP.
- `videos/`: carpeta con archivos de video para pruebas.
- `docs/`: documentos del proyecto, incluyendo la planeación.
- `README.md`: este archivo con la descripción del proyecto.

## Cómo ejecutar el proyecto

1. Instala Go: https://go.dev/dl/
2. Desde la raíz del proyecto ejecuta:

```bash
go run main.go
