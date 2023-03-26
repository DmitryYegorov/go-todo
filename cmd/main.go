package main

import (
	"log"

	todo "github.com/DmitryYegorov/go-todo/pkg"
	"github.com/DmitryYegorov/go-todo/pkg/handler"
	"github.com/DmitryYegorov/go-todo/pkg/repository"
	"github.com/DmitryYegorov/go-todo/pkg/service"
)

func main() {

	repos := repository.NewRepository()
	services := service.NewService(repos)

	server := new(todo.Server)
	handlers := handler.NewHandler(services)

	if err := server.Run("8080", handlers.InitRoutes()); err != nil {
		log.Fatalf("Error ocurred while running HTTP server %s", err.Error())
	}
}
