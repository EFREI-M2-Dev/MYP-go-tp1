package main

import (
	"flag"
	"fmt"
)

func main() {
	action := flag.String("action", "", "action to add")
	name := flag.String("name", "", "nom")
	number := flag.String("tel", "", "number")
	flag.Parse()

	repertoire := make(map[string]string) //key: name, value: number

	switch *action {
	case "list":
		printContacts(repertoire)

	case "add":
		addContact(*name, *number, repertoire)
	case "remove":
		removeContact(*name, repertoire)
	case "update":
		updateContact(*name, *number, repertoire)
	case "help":
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

	return ok
}

func printHelp() {
	fmt.Println("Usage:")
	fmt.Println("--flag [list|help]")
	fmt.Println("--flag remove --name NAME")
	fmt.Println("--flag [add|update] --name NAME --tel NUMBER")
}
