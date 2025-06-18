package main

import (
	"log"
	todoapp "todo-app"
	"todo-app/pkg/handler"
)

func main() {
	handler := new(handler.Handler)

	srv := new(todoapp.Server)
	if err := srv.Run("8082", handler.InitRoutes()); err != nil {
		log.Fatal("Error occuerd while running http server: %s", err.Error())
	}
}
