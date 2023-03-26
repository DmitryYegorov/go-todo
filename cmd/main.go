package main

import (
	"log"

	todo "github.com/DmitryYegorov/go-todo/pkg"
	"github.com/DmitryYegorov/go-todo/pkg/handler"
)

func main() {
	server := new(todo.Server)
	handlers := new(handler.Handler)

	if err := server.Run("8080", handlers.InitRoutes()); err != nil {
		log.Fatalf("Error ocurred while running HTTP server %s", err.Error())
	}
}
