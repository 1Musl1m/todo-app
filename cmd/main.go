package main

import (
	"log"

	"github.com/1Musl1m/todo-app"
)

func main() {
	srv := new(todo.Server)
	if err := srv.Run("8000"); err != nil {
		log.Fatal("error running http server: %s", err.Error())
	} 
}