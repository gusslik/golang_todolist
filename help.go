package main

import "fmt"

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
