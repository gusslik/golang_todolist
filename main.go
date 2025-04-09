package main

import (
	"fmt"
	"log"
	"os"
)

type Command struct {
	Name        string
	Description string
}

var commands = []Command{
	{"help", "display all commands"},
	{"add <name> <description>", "add new task with specified name and description"},
}

func getArgs() []string {
	return os.Args
}

func identifyCommand(command string) {
	switch command {
	case "help":
		help(commands)
	case "add":
		add()
	default:
		println("Unknown command. Please write help to display all available commands")
	}
}

func help(commands []Command) {
	maxLength := 0
	for _, command := range commands {
		if len(command.Name) > maxLength {
			maxLength = len(command.Name)
		}
	}
	fmt.Println("Command List")
	for _, command := range commands {
		fmt.Printf("\t%-*s\t%s\n", maxLength, command.Name, command.Description)
	}
}

func add() {

	if _, err := os.Stat("tasks.csv"); os.IsNotExist(err) {
		fmt.Println("File doesn't exist")
		file, err := os.Create("tasks.csv")
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(file.Name())
	} else {
		fmt.Println("File exists")
	}
}

func main() {
	args := getArgs()
	if len(args) > 1 {
		identifyCommand(args[1])
	} else {
		println("Unknown command. Please write help to display all available commands")
	}
}
