package main

import (
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

func main() {
	args := getArgs()
	if len(args) > 1 {
		identifyCommand(args[1])
	} else {
		println("Unknown command. Please write help to display all available commands")
	}
}
