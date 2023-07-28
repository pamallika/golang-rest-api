package main

import (
	todo "github.com/pamallika/golang-rest-api"
	"github.com/pamallika/golang-rest-api/pkg/handler"
	"github.com/pamallika/golang-rest-api/pkg/repository"
	"github.com/pamallika/golang-rest-api/pkg/service"
	"github.com/spf13/viper"
	"log"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("error config: %s", err.Error())
	}
	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)
	srv := new(todo.Server)
	if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		log.Fatalf("error running server: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigFile("config")
	return viper.ReadInConfig()
}
