package main

import (
	"log"

	"github.com/prongbang/wire-todos/pkg/todo"
)

func main() {
	todos, err := todo.Initial()
	if err != nil {
		log.Println("Is error")
	}
	log.Println(todos.GetTodoList())
}
