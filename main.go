package main

import (
	"fmt"
	"strings"
)

func executeCommand(cmd string) error {
	cmds := getCommandsMap()
	for key, val := range cmds {
		if strings.EqualFold(key, cmd) {
			if err := val.callback(); err != nil {
				fmt.Println(err)
			}
			return nil
		}
	}
	return fmt.Errorf("Incorrect command")
}

func getCommand() {
	fmt.Print("pokedex > ")
	var cmd string
	fmt.Scanln(&cmd)
	if err := executeCommand(cmd); err != nil {
		fmt.Println(err)
	}
}

func main() {
	fmt.Println("Welcome to pokedex")
	for {
		getCommand()
	}
}
