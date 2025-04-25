Lievelingskleur Applicatie
Overzicht
Deze applicatie vraagt de gebruiker om zijn of haar lievelingskleur in te voeren.
Als de kleur bekend is in het bestand kleuren.json, wordt een bijbehorende tekst geschreven naar output.txt.
Als de kleur niet bekend is, wordt er een foutmelding naar output.txt geschreven.
JSON aanpassen
•	Open het bestand kleuren.json.
•	Voeg nieuwe kleuren toe in dezelfde structuur, bijvoorbeeld:
json
KopiërenBewerken
{
  "kleur": "oranje",
  "tekst": "Oranje straalt warmte en energie uit."
}
Applicatie uitvoeren
•	Zorg dat main.go en kleuren.json in dezelfde map staan.
•	Run de applicatie via terminal/command line:
bash
KopiërenBewerken
go run main.go
Foutafhandeling
Als een kleur niet wordt gevonden in het JSON-bestand:
•	De foutmelding wordt getoond in de console.
•	De foutmelding wordt ook opgeslagen in output.txt.
