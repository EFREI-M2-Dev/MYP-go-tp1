package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
)

const REPERTOIREJSON = "repertoire.json"

func loadContacts() map[string]string {
	repertoire := make(map[string]string) //key: name, value: number

	data, err := os.ReadFile(REPERTOIREJSON)
	if err != nil {
		if os.IsNotExist(err) {
			return repertoire // return empty if file doesn't exist
		}
		fmt.Println("Error reading file:", err)
		os.Exit(1)
	}

	err = json.Unmarshal(data, &repertoire)
	if err != nil {
		fmt.Println("Error parsing JSON:", err)
		os.Exit(1)
	}

	return repertoire
}

func saveContacts(repertoire map[string]string) {
	data, err := json.MarshalIndent(repertoire, "", "  ")
	if err != nil {
		fmt.Println("Error encoding JSON:", err)
		os.Exit(1)
	}

	err = os.WriteFile(REPERTOIREJSON, data, 0644)
	if err != nil {
		fmt.Println("Error writing file:", err)
		os.Exit(1)
	}
}

func main() {
	action := flag.String("action", "", "action to add")
	name := flag.String("name", "", "nom")
	number := flag.String("tel", "", "number")
	flag.Parse()

	repertoire := loadContacts()

	switch *action {
	case "list":
		printContacts(repertoire)

	case "add":
		addContact(*name, *number, repertoire)
		saveContacts(repertoire)

	case "remove":
		removeContact(*name, repertoire)
		saveContacts(repertoire)

	case "update":
		updateContact(*name, *number, repertoire)
		saveContacts(repertoire)

	case "help":
		printHelp()
	default:
		printHelp()
	}
}

func printContacts(repertoire map[string]string) {
	for name, number := range repertoire {
		fmt.Printf("%s: %s\n", name, number)
	}
}

func addContact(name string, number string, repertoire map[string]string) bool {
	if name == "" || number == "" {
		fmt.Println("--name and --tel is required")
		return false
	}

	if !isNameAvailable(name, repertoire) {
		fmt.Println("Name already in use")
		return false
	}

	repertoire[name] = number
	return true
}

func removeContact(name string, repertoire map[string]string) {
	delete(repertoire, name)
}

func updateContact(name string, number string, repertoire map[string]string) {
	repertoire[name] = number
}

func isNameAvailable(name string, repertoire map[string]string) bool {
	_, ok := repertoire[name]

	return !ok
}

func printHelp() {
	fmt.Println("Usage:")
	fmt.Println("--flag [list|help]")
	fmt.Println("--flag remove --name NAME")
	fmt.Println("--flag [add|update] --name NAME --tel NUMBER")
}
