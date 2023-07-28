package main

import (
	todo "github.com/pamallika/golang-rest-api"
	"github.com/pamallika/golang-rest-api/pkg/handler"
	"log"
)

func main() {
	handlers := new(handler.Handler)
	srv := new(todo.Server)
	if err := srv.Run("0000", handlers.InitRoutes()); err != nil {
		log.Fatalf("error running server: %s", err.Error())
	}
}
