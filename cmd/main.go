//go:build ignore

package main

import (
	"log"
	"net/http"

	"forum/internal/handlers"
	internal "forum/internal/service"
	"forum/internal/storage"
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
