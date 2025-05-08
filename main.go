package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

func main() {
	//file, _ := os.ReadFile("kleuren.json")
	// _ Negeer eventuele fouten. programma stopt gewoon als het crasht
	//betere versie
	file, err := os.ReadFile("Kleuren.json")
	if err != nill {
		fmt.Println("Fout bij het lezen van kleuren.json", err)
		return
	}
	

	var kleuren map[string]string
	json.Unmarshal(file, &kleuren)

	var input string
	fmt.Println("Wat is je lievelingskleur?")
	fmt.Scanln(&input)

	// Haal spaties en hoofdletters weg
	input = strings.ToLower(input)
	input = strings.TrimSpace(input)

	tekst, gevonden := kleuren[input]
	if gevonden {
		os.WriteFile("output.txt", []byte(tekst), 0644)
		fmt.Println(tekst)
	} else {
		foutmelding := "Onbekende kleur ingevoerd: " + input
		os.WriteFile("output.txt", []byte(foutmelding), 0644)
		fmt.Println(foutmelding)
	}
}
