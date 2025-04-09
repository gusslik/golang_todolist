package main

import (
	"fmt"
	"log"
	"os"
)

func add() {
	var file *os.File
	if _, err := os.Stat("tasks.csv"); os.IsNotExist(err) {
		fmt.Println("File doesn't exist")
		file, err = os.Create("tasks.csv")
		if err != nil {
			log.Fatal(err)
		}
	} else {
		file, err = os.Open("tasks.csv")
		if err != nil {
			log.Fatal(err)
		}
	}

	fmt.Println(file)
}
