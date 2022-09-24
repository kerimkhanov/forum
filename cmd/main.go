//go:build ignore

package main

import (
	"forum/internal/handlers"
	"forum/internal/storage"
	"log"
	"net/http"

	internal "forum/internal/service"
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
