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

func identifyCommand(args []string) {
	switch args[1] {
	case "help":
		help(commands)
	case "add":
		add(args)
	case "complete":
		complete(args)
	default:
		println("Unknown command. Please write help to display all available commands")
	}
}

func main() {
	args := getArgs()
	if len(args) > 1 {
		identifyCommand(args)
	} else {
		println("Unknown command. Please write help to display all available commands")
	}
}
