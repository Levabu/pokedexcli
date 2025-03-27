package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/levabu/pokedexcli/commands"
	"github.com/levabu/pokedexcli/internal/pokeapi"
)

func cleanInput(text string) []string {
	lower := strings.ToLower(text)
	return strings.Fields(lower)
}

func startRepl(cfg *pokeapi.Config) {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		err := scanner.Err()
		if err != nil {
			log.Fatal(err)
		}
		text := scanner.Text()
		tokens := cleanInput(text)
		if len(tokens) == 0 {
			continue
		}
		commandName := tokens[0]
		c, commandExists := commands.GetCommands()[commandName]
		if !commandExists {
			fmt.Println("Unknown command")
			continue
		}
		
		if err := c.Callback(cfg); err != nil {
			fmt.Printf("%v\n", err)
		}
	}
}