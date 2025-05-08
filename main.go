package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"log"
)

func main() {

	//file, _ := os.ReadFile("kleuren.json")
	// _ Negeer eventuele fouten. programma stopt gewoon als het crasht
	//betere versie

	//logbestand openen of aanmaken
	logfile, err := os.OpenFile("log.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println("kon logbestand niet openen", err)
		fmt.Println("Kon logbestand niet openen", err)
		return
	}
	log.SetOutput(logfile)

	file, err := os.ReadFile("Kleuren.json")
	if err != nil {
		log.Println("fout bij het lezen van kleuren.json", err)
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
		log.Println("Kleur herkend: ", input)
		
	} else {
		foutmelding := "Onbekende kleur ingevoerd: " + input
		os.WriteFile("output.txt", []byte(foutmelding), 0644)
		log.Println(foutmelding)
		fmt.Println(foutmelding)
	}
}
