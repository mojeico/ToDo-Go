package main

import (
	"fmt"
	"log"
	ToDo_Go "myGoMod"
	"myGoMod/handler"
)

func main() {

	fmt.Println("test server")

	handlers := new(handler.Handler)
	srv := new(ToDo_Go.Server)

	if err := srv.Run("8000", handlers.InitRouters()); err != nil {
		log.Fatalf("error start server - %s", err.Error())
	}
}
