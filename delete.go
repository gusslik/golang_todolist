package main

import (
	"encoding/csv"
	"log"
	"os"
	"strconv"
)

func delete(file *os.File, id int) ([][]string, bool) {
	reader := csv.NewReader(file)

	records, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	var buffer [][]string
	isRecordFound := false

	// Read through the file and find the record with the specified id
	for i, record := range records {
		// Add header
		if i == 0 {
			buffer = append(buffer, record)
			continue
		}

		// Convert the id of the task to integer
		recordId, err := strconv.Atoi(record[0])
		if err != nil {
			log.Fatal(err)
		}

		// Filter the specified id
		if recordId != id {
			buffer = append(buffer, record)
		} else {
			isRecordFound = true
		}
	}

	return buffer, isRecordFound
}
