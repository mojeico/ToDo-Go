package main

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
	ToDo_Go "myGoMod"
	"myGoMod/handler"
	"myGoMod/repository"
	service "myGoMod/service"
)

func main() {

	fmt.Println("Test server")

	if err := initConfig(); err != nil {
		log.Fatalf("error initializing configs: %s ", err.Error())
	}

	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(ToDo_Go.Server)

	if err := srv.Run(viper.GetString("port"), handlers.InitRouters()); err != nil {
		log.Fatalf("error start server - %s", err.Error())
	}

}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
