package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func start() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex> ")
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := cleanInput(scanner.Text())
		if len(line) == 0 {
			continue
		}
		if cmd, ok := getCommands()[line[0]]; ok {
			err := cmd.callback()
			if err != nil {
				fmt.Println("Error:", err)
			}
		} else {
			fmt.Println("Unknown command")
			err := getCommands()["help"].callback()
			if err != nil {
				fmt.Println("Error:", err)
			}

		}
	}
}

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

func cleanInput(input string) []string {
	// Remove leading and trailing whitespace and lowercase the input
	lowered := strings.ToLower(input)
	words := strings.Fields(lowered)
	return words

}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exits the Pokedex",
			callback:    commandExit,
		},
	}
}
