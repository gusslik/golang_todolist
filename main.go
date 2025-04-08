package main

import (
	"os"
)

func getArgs() []string {
	return os.Args
}

func identifyCommand(command string) {
	switch command {
	case "help":
		help()
	default:
		println("Unknown command. Please write help to display all available commands")
	}
}

func help() {
	// Add a dictionary of commands and their descriptions which is displayed using fmt.Printf and a loop
}

func main() {
	args := getArgs()
	if len(args) > 1 {
		identifyCommand(args[1])
	} else {
		println("Unknown command. Please write help to display all available commands")
	}
}
