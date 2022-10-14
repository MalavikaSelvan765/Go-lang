package main

import (
    "log"
    "net/http"
    "github.com/gorilla/mux"
    _ "github.com/jinzhu/gorm/dialects/mysql"
    "github.com/MalavikaSelvan765/Go-lang/Book_Store/pkg/routes"
)

func main() {
    r := mux.NewRouter()
    routes.RegisterBookStoreRoutes(r)
    http.Handle("/", r)
    log.Fatal(http.ListenAndServe("localhost:9010", r))
}