package main

import "strings"

func cleanInput(text string) []string {
	input := strings.Fields(text)

	for i, s := range input {
		input[i] = strings.ToLower(s)
	}

	return input
}
