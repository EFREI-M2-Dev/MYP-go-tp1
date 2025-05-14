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

	var repertoire map[string]string //key: name, value: number

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
		fmt.Println("--flag [list|add|remove|update|help]")
	}
}

func printContacts(contacts map[string]string) {
	for name, number := range contacts {
		fmt.Println(name + ": " + number + "\n")
	}
}

func addContact(name string, number string, repertoire map[string]string) bool {
	if !isNameAvailble(name, repertoire) {
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

func isNameAvailble(name string, repertoire map[string]string) bool {
	_, ok := repertoire[name]

	return ok
}
