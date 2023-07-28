package main

import (
	"fmt"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/pamallika/golang-rest-api"
	"github.com/pamallika/golang-rest-api/pkg/handler"
	"github.com/pamallika/golang-rest-api/pkg/repository"
	"github.com/pamallika/golang-rest-api/pkg/service"
	"github.com/spf13/viper"
	"log"
	"os"
)

func main() {
	fmt.Print(viper.GetString("host"))

	if err := initConfig(); err != nil {
		log.Fatalf("error config: %s", err.Error())
	}

	if err := godotenv.Load(); err != nil {
		log.Fatalf("error env: %s", err.Error())
	}
	db, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
	})
	if err != nil {
		log.Fatalf("error connection db %s", err.Error())
	}
	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)
	srv := new(todo.Server)
	if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		log.Fatalf("error running server: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
