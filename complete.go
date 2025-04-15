package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
)

func complete(args []string) {
	var file *os.File

	// Check the right number of parameters
	if len(args) != 3 {
		fmt.Println("Wrong number of arguments. Make sure that the parameters are <task id>")
		return
	}

	// Check if the file with tasks exists
	if _, err := os.Stat("tasks.csv"); os.IsNotExist(err) {
		fmt.Println("You don't have any tasks added yet. Please use add <name> <description> to add a new task")
		return
	}

	// Convert the specified task id into integer
	id, err := strconv.Atoi(args[2])
	if err != nil {
		log.Fatal(err)
	}

	// Check if the id is the right number
	if id < 0 {
		fmt.Println("Wrong task id")
		return
	}

	// Open files with read and write permissions
	file, err = os.OpenFile("tasks.csv", os.O_RDWR, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	buffer, isFound := delete(file, id)
	if isFound {
		// Truncate the file by recreating it
		file, err = os.Create("tasks.csv")
		if err != nil {
			log.Fatal(err)
		}

		// Write the updated list of tasks to the file
		writer := csv.NewWriter(file)
		writer.WriteAll(buffer)

		if err := writer.Error(); err != nil {
			log.Fatal(err)
		}
		fmt.Println("Task has been completed successfully")
	} else {
		fmt.Println("Task with this id couldn't be found")
	}
}
