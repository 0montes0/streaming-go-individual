package models

type Video struct {
    ID       int
    Title    string
    Category string
    FilePath string
}
var Videos = []Video{
    {ID: 1, Title: "Go Tutorial", Category: "Tutorial", FilePath: "videos/go_tutorial.mp4"},
    {ID: 2, Title: "Nature", Category: "Documental", FilePath: "videos/nature.mp4"},
