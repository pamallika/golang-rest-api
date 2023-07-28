package main

import (
	todo "github.com/pamallika/golang-rest-api"
	"github.com/pamallika/golang-rest-api/pkg/handler"
	"github.com/pamallika/golang-rest-api/pkg/repository"
	"github.com/pamallika/golang-rest-api/pkg/service"
	"log"
)

func main() {
	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)
	srv := new(todo.Server)
	if err := srv.Run("0000", handlers.InitRoutes()); err != nil {
		log.Fatalf("error running server: %s", err.Error())
	}
}
