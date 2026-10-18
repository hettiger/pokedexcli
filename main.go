package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		if ok := scanner.Scan(); !ok {
			break
		}
		args := cleanInput(scanner.Text())
		if len(args) > 0 {
			fmt.Printf("Your command was: %s", args[0])
			fmt.Println()
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("scanning for input failed: %v", err)
	}
}
