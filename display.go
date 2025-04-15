package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

func display() {
	var file *os.File

	if _, err := os.Stat("tasks.csv"); os.IsNotExist(err) {
		fmt.Println("You don't have any tasks added yet")
		return
	} else {
		file, err = os.OpenFile("tasks.csv", os.O_RDONLY, 0644)
		if err != nil {
			log.Fatal(err)
		}

		reader := csv.NewReader(file)

		records, err := reader.ReadAll()
		if err != nil {
			log.Fatal(err)
		}

		for i, record := range records {
			if i == 0 {
				continue
			}
			fmt.Printf("ID: %s | Name: %s | Description: %s | Completed: %s\n", record[0], record[1], record[2], record[3])
		}

	}
}
