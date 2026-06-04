package main

import (
	"fmt"

	"github.com/Lohuama/simple-go-mod/config"
)

type Task struct {
	ID int
	Description string
	Completed bool
}

func main(){
	tasks := []Task{}
	var nextID int = 1

	config.SetupDB()
}

func addTasks(tasks *[]Task, description string, nextID *int){
	newTask := Task {
		ID: *nextID,
		Description: description,
		Completed: false,
	}
	*tasks = append(*tasks, newTask)
	*nextID++
	fmt.Println("Tarefa adicionada com sucesso!")
}