package main

import (
	"testing"
)

func TestAddContact(t *testing.T) {
	repertoire := make(map[string]string)

	// add contact
	if !addContact("John", "123456", repertoire) {
		t.Error("L'ajout aurait dû réussir")
	}
	if num, exists := repertoire["John"]; !exists || num != "123456" {
		t.Error("Le contact n'a pas été ajouté correctement")
	}

	// add contact with existing name
	if addContact("John", "789012", repertoire) {
		t.Error("L'ajout aurait dû échouer car le nom existe déjà")
	}

	// add contact with empty name
	if addContact("", "123456", repertoire) {
		t.Error("L'ajout aurait dû échouer car le nom est vide")
	}

	// add contact with empty number
	if addContact("Alice", "", repertoire) {
		t.Error("L'ajout aurait dû échouer car le numéro est vide")
	}
}

func TestRemoveContact(t *testing.T) {
	repertoire := map[string]string{
		"John": "123456",
	}

	// remove existing contact
	removeContact("John", repertoire)
	if _, exists := repertoire["John"]; exists {
		t.Error("Le contact aurait dû être supprimé")
	}

	// remove non-existing contact (should not cause an error)
	removeContact("Inexistant", repertoire)
}

func TestUpdateContact(t *testing.T) {
	repertoire := map[string]string{
		"John": "123456",
	}

	// update existing contact
	updateContact("John", "789012", repertoire)
	if num := repertoire["John"]; num != "789012" {
		t.Error("Le numéro n'a pas été mis à jour correctement")
	}

	// update non-existing contact (creation)
	updateContact("Alice", "345678", repertoire)
	if num := repertoire["Alice"]; num != "345678" {
		t.Error("Le nouveau contact n'a pas été créé correctement")
	}
}

func TestIsNameAvailable(t *testing.T) {
	repertoire := map[string]string{
		"John": "123456",
	}

	// Test with an existing name
	if isNameAvailable("John", repertoire) {
		t.Error("Le nom John devrait être indisponible")
	}

	// Test with a new name
	if !isNameAvailable("Alice", repertoire) {
		t.Error("Le nom Alice devrait être disponible")
	}
}

func TestSaveAndLoadContacts(t *testing.T) {
	// create a test directory
	repertoireTest := map[string]string{
		"John":  "123456",
		"Alice": "789012",
	}

	// save the directory
	saveContacts(repertoireTest)

	// load the directory
	repertoireCharge := loadContacts()

	// check that the data is identical
	for name, number := range repertoireTest {
		if loadedNumber, exists := repertoireCharge[name]; !exists || loadedNumber != number {
			t.Errorf("The loaded data does not match the saved data for %s", name)
		}
	}
}
