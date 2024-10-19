package main

import (
	"fmt"
	"os"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

func getCommandsMap() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "show all available commands",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "exit the program",
			callback:    commandExit,
		},
		"map": {
			name:        "map",
			description: "get next 20 of locations",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "get previous 20 of locations",
			callback:    commandMapb,
		},
	}
}

func commandHelp() error {
	cmds := getCommandsMap()
	for _, val := range cmds {
		fmt.Println(val.name + " - " + val.description)
	}
	return nil
}

func commandExit() error {
	fmt.Println("Goodbye!")
	os.Exit(0)
	return nil
}

func commandMap() error {
	areas := poke.getLocationAreas()
	for _, area := range areas {
		fmt.Printf("Area: %v\n", area)
	}
	return nil
}

func commandMapb() error {
	return nil
}
