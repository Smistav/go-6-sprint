package main

import (
	"log"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/server"
)

func main() {
	logger := log.Default()
	server := server.NewServer(logger)
	server.Logger.Println("Сервер запущен на http://localhost:8080")
	if err := server.Server.ListenAndServe(); err != nil {
		server.Logger.Fatal(err)
	}
}
