package main

import (
	"github.com/spf13/viper"
	"log"
	"restapi"
	"restapi/pkg/handler"
	"restapi/pkg/repository"
	"restapi/pkg/service"
)

func main() {
	if err := InitConfig(); err != nil {
		log.Fatalf("error: %s", err.Error())
	}

	//handlers := new(handler.Handler)
	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(restapi.Server)
	if err := srv.Run(viper.GetString("8000"), handlers.InitRoutes()); err != nil {
		log.Fatalf("error: %s", err.Error())
	}
}

func InitConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
