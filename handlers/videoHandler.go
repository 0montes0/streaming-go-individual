package handlers

import (
    "encoding/json"
    "net/http"
    "os"
    "strconv"
    "streaming/models"
)

func GetVideos(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(models.Videos)
}

func StreamVideo(w http.ResponseWriter, r *http.Request) {
    idStr := r.URL.Query().Get("id")
    id, err := strconv.Atoi(idStr)
    if err != nil {
        http.Error(w, "ID inválido", http.StatusBadRequest)
        return
    }

    var selected string
    for _, v := range models.Videos {
        if v.ID == id {
            selected = v.FilePath
            break
        }
    }

    if selected == "" {
        http.Error(w, "Video no encontrado", http.StatusNotFound)
        return
    }

    f, err := os.Open(selected)
    if err != nil {
        http.Error(w, "No se puede abrir el archivo", http.StatusInternalServerError)
        return
    }
    defer f.Close()

    fi, err := f.Stat()
    if err != nil {
        http.Error(w, "No se puede obtener información del archivo", http.StatusInternalServerError)
        return
    }

    http.ServeContent(w, r, selected, fi.ModTime(), f)
}
