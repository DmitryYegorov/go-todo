package main

import (
	"log"

	_ "github.com/lib/pq"

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

	db, err := repository.NewPostgresDb(repository.Config{
		Host:     "localhost",
		Port:     "5432",
		Username: "postgres",
		Password: "postgres",
		SSLMode:  "disable",
		DBName:   "go-todo",
	})
	if err != nil {
		log.Fatalf("failed database initialization: %s", err.Error())
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos)

	server := new(todo.Server)
	handlers := handler.NewHandler(services)

	if err := server.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		log.Fatalf("Error ocurred while running HTTP server %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
