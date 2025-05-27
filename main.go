package main

import (
    "log"
    "net/http"

    "github.com/gorilla/mux"
    "streaming/handlers"  // paquete con los handlers que tú creas
)

func main() {
    // Crear router nuevo
    r := mux.NewRouter()

    // Definir rutas y funciones manejadoras
    r.HandleFunc("/videos", handlers.GetVideos).Methods("GET")      // Listar videos
    r.HandleFunc("/stream", handlers.StreamVideo).Methods("GET")    // Reproducir video

    // Iniciar servidor en puerto 8080
    log.Println("Servidor iniciado en http://localhost:8080")
    err := http.ListenAndServe(":8080", r)
    if err != nil {
        log.Fatal(err)
    }
}
