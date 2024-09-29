package main

import (
	"go-todo-cli/todo"
	"log"
)

func main() {
	todos := todo.TodoItems{}
	storage := todo.NewStorage[todo.TodoItems]("todos.json")

	if err := storage.Load(&todos); err != nil {
		log.Fatalf("Failed to load todos: %v", err)
	}

	cmdFlags := todo.NewCmdFlags()
	cmdFlags.Execute(&todos)

	if err := storage.Save(todos); err != nil {
		log.Fatalf("Failed to save todos: %v", err)
	}
}
