package main

import (
	"encoding/csv"
	"log"
	"os"
	"strconv"
)

func delete(file *os.File, id int) [][]string {
	reader := csv.NewReader(file)

	records, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	var buffer [][]string

	// Read through the file and find the record with the specified id
	for i, record := range records {
		if i == 0 {
			buffer = append(buffer, record)
			continue
		}

		recordId, err := strconv.Atoi(record[0])
		if err != nil {
			log.Fatal(err)
		}

		if recordId != id {
			buffer = append(buffer, record)
		}
	}

	return buffer
}
