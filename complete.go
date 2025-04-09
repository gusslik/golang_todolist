package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func complete(args []string) {
	var file *os.File

	if len(args) != 3 {
		fmt.Println("Wrong number of arguments. Make sure that the parameters are <task id>")
		return
	}

	if _, err := os.Stat("tasks.csv"); os.IsNotExist(err) {
		fmt.Println("You don't have any tasks added yet. Please use add <name> <description> to add a new task")
		return
	}

	id, err := strconv.Atoi(args[2])
	if err != nil {
		log.Fatal(err)
	}

	if id < 0 {
		fmt.Println("Wrong task id")
		return
	}

	file, err = os.OpenFile("tasks.csv", os.O_RDWR, 0644)
	if err != nil {
		log.Fatal(err)
	}

	file.Name()
	fmt.Println("Id of the completed task:", id)
}
