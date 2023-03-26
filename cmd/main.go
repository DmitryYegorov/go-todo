package main

import (
	"log"

	todo "github.com/DmitryYegorov/go-todo/pkg"
	"github.com/DmitryYegorov/go-todo/pkg/handler"
	"github.com/DmitryYegorov/go-todo/pkg/repository"
	"github.com/DmitryYegorov/go-todo/pkg/service"
	"github.com/spf13/viper"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("Error init config: %s", err.Error())
	}

	repos := repository.NewRepository()
	services := service.NewService(repos)

	server := new(todo.Server)
	handlers := handler.NewHandler(services)

	if err := server.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		log.Fatalf("Error ocurred while running HTTP server %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigFile("config")

	return viper.ReadInConfig()
}
