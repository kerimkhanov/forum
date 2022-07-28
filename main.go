//go:build ignore

package main

import (
	"forum/handlers"
	"forum/internal"
	"forum/internal/storage"
	"log"
	"net/http"
)

func main() {
	db := storage.CreateDb()
	router := http.NewServeMux()
	storage := storage.NewDatabase(db)
	service := internal.NewService(storage)
	handler := handlers.NewHandler(service)
	handler.Register(router)
	log.Fatal(http.ListenAndServe(":8080", router))
}
