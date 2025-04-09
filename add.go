package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
)

func add(args []string) {
	if len(args) != 4 {
		fmt.Println("Wrong number of arguments. Make sure that the parameters are <name> <description>")
		return
	}

	var file *os.File
	var writer *csv.Writer
	var count = 0

	header := []string{"UUID", "Name", "Description", "Status"}

	// Creates the file if it doesn't exist, otherwise opens the file for appending
	if _, err := os.Stat("tasks.csv"); os.IsNotExist(err) {
		file, err = os.Create("tasks.csv")
		if err != nil {
			log.Fatal(err)
		}

		writer = csv.NewWriter(file)

		// Add header to the top of the CSV file
		writer.Write(header)
		count++
	} else {
		file, err = os.OpenFile("tasks.csv", os.O_APPEND|os.O_RDONLY, 0644)
		if err != nil {
			log.Fatal(err)
		}

		writer = csv.NewWriter(file)
	}

	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	count += len(records)

	// Writes the specified data to the database in the format: {uuid, name, description, isFinished}
	err = writer.Write([]string{strconv.Itoa(count), args[2], args[3], "false"})
	if err != nil {
		log.Fatal(err)
	}

	writer.Flush()
	if err := writer.Error(); err != nil {
		log.Fatal(err)
	}
}
