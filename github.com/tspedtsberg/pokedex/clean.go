package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startrepl() {
	reader := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		reader.Scan()

		words := cleanInput(reader.Text())
		if len(words) == 0 {
			continue
		}
		firstWord := words[0]
		fmt.Printf("Your command was: %s\n", firstWord)
	}
}


func cleanInput(text string) []string {
	s := strings.ToLower(text)
	result := strings.Split(s," ")
	return result
}